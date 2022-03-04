package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	s "strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/joho/godotenv"

	rs "webappsapi/main/response"
)

var ctx context.Context
var cld *cloudinary.Cloudinary

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	cld, _ = cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("CLOUD_API_KEY"), os.Getenv("CLOUD_API_SECRET"))
	ctx = context.Background()
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	imageType := r.Form.Get("type")

	if err != nil || imageType == "" {
		response := rs.FailedResponse{}
		if imageType == "" {
			response = rs.GetFailedResponse("key type not found")
		} else {
			response = rs.GetFailedResponse(err.Error())
		}
		json.NewEncoder(w).Encode(response)
	} else {
		defer file.Close()
		resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: s.Join([]string{imageType, handler.Filename}, "/")})
		if err != nil {
			log.Printf("Error while Upload File, Reason: %v\n", err)
			json.NewEncoder(w).Encode(rs.GetFailedResponse(err.Error()))
		} else {
			response := rs.GetSuccessResponseData()
			response.Data["public_url"] = resp.SecureURL
			json.NewEncoder(w).Encode(response)
		}
	}

}
