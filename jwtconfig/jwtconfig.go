package jwtconfig

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"

	m "webappsapi/main/models"
	rs "webappsapi/main/response"
	"webappsapi/main/service"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func CreateToken(user m.User) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.Id
	atClaims["email"] = user.Email
	atClaims["company_id"] = user.CompanyId
	atClaims["role_id"] = user.RoleId
	atClaims["exp"] = time.Now().Add(time.Minute * 200).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Uid"] == nil {
			response := rs.GetFailedResponse("No Uid Found")
			json.NewEncoder(w).Encode(response)
			return
		}

		if r.Header["Authorization"] == nil {
			response := rs.GetFailedResponse("No Token Found")
			json.NewEncoder(w).Encode(response)
			return
		}

		token, err := VerifyToken(r)

		if err != nil {
			response := rs.GetFailedResponse("Your Token has been expired.")
			json.NewEncoder(w).Encode(response)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			pathUrl := r.URL.Path
			claims := token.Claims.(jwt.MapClaims)
			data := claims["role_id"].(string)
			isValid := isHasValidUrlAccess(data, pathUrl)
			fmt.Println("user_id : ", claims["user_id"])
			fmt.Println("pathUrl : ", pathUrl)
			fmt.Println("role_id : ", data)
			if isValid {
				handler.ServeHTTP(w, r)
				return
			} else {
				response := rs.GetFailedResponse("Not Authorized to access this URL")
				json.NewEncoder(w).Encode(response)
				return
			}
		}
		response := rs.GetFailedResponse("Not Authorized token")
		json.NewEncoder(w).Encode(response)
	}
}

func isHasValidUrlAccess(role_id string, path string) bool {
	role := service.FindRoleById(role_id)
	accessLevelId := role.AccessLevelId.Hex()
	access := service.FindAccessByIdAndUrl(accessLevelId, path)
	return !reflect.ValueOf(access).IsZero()
}
