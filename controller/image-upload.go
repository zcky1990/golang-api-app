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
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	imageType := r.Form.Get("type")

	if err != nil || imageType == "" {
		response := FailedResponse{}
		if imageType == "" {
			response = GetFailedResponse("key type not found")
		} else {
			response = GetFailedResponse(err.Error())
		}
		json.NewEncoder(w).Encode(response)
	} else {
		defer file.Close()
		// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		// fmt.Printf("File Size: %+v\n", handler.Size)
		// fmt.Printf("MIME Header: %+v\n", handler.Header)

		resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: s.Join([]string{imageType, handler.Filename}, "/")})
		if err != nil {
			log.Printf("Error while Upload File, Reason: %v\n", err)
			json.NewEncoder(w).Encode(GetFailedResponse(err.Error()))
		} else {
			response := GetSuccessResponse()
			response.Data["public_url"] = resp.SecureURL
			json.NewEncoder(w).Encode(response)
		}
	}

}
