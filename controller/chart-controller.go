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

func AddChart(w http.ResponseWriter, r *http.Request) {
	var request rq.ChartRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newChartData := m.InsertChartMongoModels{
		UserId:    request.UserId,
		ProductId: request.ProductId,
	}
	chartData, err := bson.Marshal(newChartData)
	if err != nil {
		panic(err)
	}
	data, err := service.InsertOneChart(chartData)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
	}
	fmt.Println(data)
	response := rs.GetSuccessResponseData()
	response.Data["message"] = "Successffully add product to chart."
	json.NewEncoder(w).Encode(response)
}

func GetChartDetailByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	fmt.Println("userID", userId)
	chart := service.GetChartListByUserId(userId)
	if reflect.ValueOf(chart).IsZero() {
		json.NewEncoder(w).Encode(rs.GetFailedResponse("Chart empty"))
		return
	}
	response := rs.GetSuccessResponse(&fiber.Map{"product": chart})
	json.NewEncoder(w).Encode(response)
}