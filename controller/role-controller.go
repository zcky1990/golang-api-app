package controller

import (
	"encoding/json"
	"net/http"

	"reflect"

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
	role := service.FindRoleByName(request.RoleName)
	if !reflect.ValueOf(role).IsZero() {
		json.NewEncoder(w).Encode(rs.GetFailedResponse("Role exists"))
		return
	}
	access := service.FindAccessById(request.AccessLevelID)
	if reflect.ValueOf(access).IsZero() {
		json.NewEncoder(w).Encode(rs.GetFailedResponse("Access Id not exists"))
		return
	}
	newRoleData := m.InsertRoleMongoModels{
		RoleName:      request.RoleName,
		Description:   request.Description,
		AccessLevelId: access.Id,
	}
	roleData, err := bson.Marshal(newRoleData)
	if err != nil {
		panic(err)
	}
	data, err := service.InsertOneRole(roleData)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
	}
	response := rs.GetSuccessResponse(&fiber.Map{"data": data})
	json.NewEncoder(w).Encode(response)
}
