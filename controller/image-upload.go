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

type success struct {
	Message   string `json:"message"`
	PublicUrl string `json:"public_url"`
}

type failed struct {
	ErrorMessage string `json:"error_message"`
}

type ImageUploadSuccessResponse struct {
	Status string  `json:"status"`
	Code   uint8   `json:"code"`
	Data   success `json:"data"`
}

type ImageUploadFailedResponse struct {
	Status string `json:"status"`
	Code   uint8  `json:"code"`
	Data   failed `json:"data"`
}

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
		failed := failed{}
		if imageType == "" {
			failed.ErrorMessage = "key type not found"
		} else {
			failed.ErrorMessage = err.Error()
		}
		response := ImageUploadFailedResponse{"failed", 200, failed}
		json.NewEncoder(w).Encode(response)
	} else {
		defer file.Close()
		// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		// fmt.Printf("File Size: %+v\n", handler.Size)
		// fmt.Printf("MIME Header: %+v\n", handler.Header)

		resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: s.Join([]string{imageType, handler.Filename}, "/")})
		if err != nil {
			log.Printf("Error while Upload File, Reason: %v\n", err)
			failed := failed{err.Error()}
			response := ImageUploadFailedResponse{"failed", 200, failed}
			json.NewEncoder(w).Encode(response)
		} else {
			success := success{"success upload file", resp.SecureURL}
			response := ImageUploadSuccessResponse{"success", 200, success}

			json.NewEncoder(w).Encode(response)
		}
	}

}
