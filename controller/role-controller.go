package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	m "webappsapi/main/models"
	rq "webappsapi/main/request"
	rs "webappsapi/main/response"
	service "webappsapi/main/service"
)

func AddRole(w http.ResponseWriter, r *http.Request) {
	var request rq.RoleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	role := service.FindRoleByName(request.Rolename)
	if (role != m.Role{}) {
		json.NewEncoder(w).Encode(rs.GetFailedResponse("Role exists"))
		return
	}

	//convert struct to bson
	role_data, err := bson.Marshal(request)
	if err != nil {
		panic(err)
	}
	data, err := service.InsertOneRole(role_data)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
	}
	response := rs.GetSuccessResponse(&fiber.Map{"data": data})
	json.NewEncoder(w).Encode(response)
}
