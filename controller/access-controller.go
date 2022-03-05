package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"reflect"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	m "webappsapi/main/models"
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

// Members
func AddMemberAccess(w http.ResponseWriter, r *http.Request) {
	var request rq.AccessRequestWithCompanyId
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	access := service.FindAccessByAccessLevelAndCompanyId(request.AccessLevel, request.CompanyId)
	if !reflect.ValueOf(access).IsZero() {
		json.NewEncoder(w).Encode(rs.GetFailedResponse("Access exists"))
		return
	}
	companyId, err := primitive.ObjectIDFromHex(request.CompanyId)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	accessParams := m.InsertAccessMongoModelsWithCompanyId{
		AccessLevel: request.AccessLevel,
		AccessType:  request.AccessType,
		ListUrl:     request.ListUrl,
		CompanyId:   companyId,
	}
	accessData, err := bson.Marshal(accessParams)
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

func GetListMemberAccess(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Println(values)
	companyId := values["company_id"][0]
	data := service.GetAccessListBaseOnCompanyId(companyId)
	response := rs.GetSuccessResponse(&fiber.Map{"data": data})
	json.NewEncoder(w).Encode(response)
}
