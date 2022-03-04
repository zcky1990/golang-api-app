package controller

import (
	"encoding/json"
	"net/http"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"

	"webappsapi/main/jwtconfig"
	m "webappsapi/main/models"
	rq "webappsapi/main/request"
	rs "webappsapi/main/response"
	service "webappsapi/main/service"
)

type UserLoginRequest = rq.UserLoginRequest
type UserSignUpRequest = rq.UserSignUpRequest
type UserAddRequest = rq.UserAddRequest
type RegisterCompanyRequest = rq.RegisterCompanyRequest

type User = m.User
type Company = m.Company
type Role = m.Role

// http function
func UserList(w http.ResponseWriter, r *http.Request) {
	response := service.GetAllUserList()
	json.NewEncoder(w).Encode(response)
}

func Search(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	query := values["query"][0]
	search_type := values["search_type"][0]

	response := service.SearchUser(search_type, query)
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var request UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email := request.Email
	password := request.Password
	user, err := service.GetUserByEmailAndPassword(email, password)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	token, err := jwtconfig.CreateToken(user.Id, user.Email)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
	}
	user.Authtoken = token
	response := rs.GetSuccessResponse(&fiber.Map{"data": user})
	json.NewEncoder(w).Encode(response)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var request UserAddRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := service.GetUserByEmail(request.Email)
	if !reflect.ValueOf(user).IsZero() {
		json.NewEncoder(w).Encode(rs.GetFailedResponse("User exists"))
		return
	}
	userData, err := bson.Marshal(request)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := service.AddUser(userData)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := rs.GetSuccessResponse(&fiber.Map{"data": data})
	json.NewEncoder(w).Encode(response)

}

type UserMongo struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Birthday  string `json:"birthday"`
	Role      Role
	Company   Company
}

func RegisterCompany(w http.ResponseWriter, r *http.Request) {
	var request RegisterCompanyRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	roleOwner := service.FindRoleByName(request.User.AccessType)
	companyModels := service.FindCompanyByEmail(request.Company.CompanyEmail)
	if (companyModels == m.Company{}) {
		companyData, err := bson.Marshal(request.Company)
		if err != nil {
			response := rs.GetFailedResponse(err.Error())
			json.NewEncoder(w).Encode(response)
			return
		}
		company, err := service.InsertOneCompany(companyData)
		if err != nil {
			response := rs.GetFailedResponse(err.Error())
			json.NewEncoder(w).Encode(response)
			return
		}
		id := company.InsertedID.(primitive.ObjectID)
		companyModels = Company{
			Id:             id,
			CompanyName:    request.Company.CompanyName,
			CompanyAddress: request.Company.CompanyAddress,
			CompanyEmail:   request.Company.CompanyEmail,
			CompanyPhone:   request.Company.CompanyPhone,
		}
	}

	user := UserMongo{
		Username:  request.User.UserName,
		Email:     request.User.Email,
		Password:  request.User.Password,
		Firstname: request.User.FirstName,
		Lastname:  request.User.LastName,
		Birthday:  request.User.Birthday,
		Role:      roleOwner,
		Company:   companyModels,
	}

	responseUser := service.GetUserByEmail(user.Email)
	if !reflect.ValueOf(responseUser).IsZero() {
		json.NewEncoder(w).Encode(rs.GetFailedResponse("User exists"))
		return
	}

	userData, err := bson.Marshal(user)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	data, err := service.AddUser(userData)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := rs.GetSuccessResponse(&fiber.Map{"data": data})
	json.NewEncoder(w).Encode(response)
}
