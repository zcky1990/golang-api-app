package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"

	m "webappsapi/main/models"
	rq "webappsapi/main/request"
	rs "webappsapi/main/response"
	service "webappsapi/main/service"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var request rq.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%+v", request)
	newProductData := m.InsertProductMongoModels{
		Tilte:      request.Tilte,
		PictureUrl: request.PictureUrl,
		Price:      request.Price,
	}
	productData, err := bson.Marshal(newProductData)
	if err != nil {
		panic(err)
	}
	data, err := service.InsertOneProduct(productData)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
	}
	response := rs.GetSuccessResponse(&fiber.Map{"data": data})
	json.NewEncoder(w).Encode(response)
}

func GetProductDetailById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	product := service.FindProductById(id)
	if reflect.ValueOf(product).IsZero() {
		json.NewEncoder(w).Encode(rs.GetFailedResponse("Product Doesn't exists"))
		return
	}
	response := rs.GetSuccessResponse(&fiber.Map{"product": product})
	json.NewEncoder(w).Encode(response)
}

func GetProductList(w http.ResponseWriter, r *http.Request) {
	data := service.GetProductList()
	response := rs.GetSuccessResponse(&fiber.Map{"product": data})
	json.NewEncoder(w).Encode(response)
}
