package controller

import (
	"encoding/json"
	"net/http"

	"reflect"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	rq "webappsapi/main/request"
	rs "webappsapi/main/response"
	service "webappsapi/main/service"
)

func AddAccess(w http.ResponseWriter, r *http.Request) {
	var request rq.AccessRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	access := service.FindAccessByAccessLevel(request.AccessLevel)
	if !reflect.ValueOf(access).IsZero() {
		json.NewEncoder(w).Encode(rs.GetFailedResponse("Access exists"))
		return
	}

	//convert struct to bson
	accessData, err := bson.Marshal(request)
	if err != nil {
		panic(err)
	}
	data, err := service.InsertOneAccess(accessData)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
	}
	response := rs.GetSuccessResponse(&fiber.Map{"data": data})
	json.NewEncoder(w).Encode(response)
}
