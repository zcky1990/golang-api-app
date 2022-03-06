package main

import (
	"fmt"
	"webappsapi/main/controller"
	"webappsapi/main/jwtconfig"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Run Server on : localhost:10000")
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/api/users/all", jwtconfig.IsAuthorized(controller.UserList)).Methods("GET")
	myRouter.HandleFunc("/api/users/search", jwtconfig.IsAuthorized(controller.Search)).Methods("GET")
	myRouter.HandleFunc("/api/users/login", controller.Login).Methods("POST")
	myRouter.HandleFunc("/api/users/company_register", controller.RegisterCompany).Methods("POST")
	myRouter.HandleFunc("/api/image/upload", jwtconfig.IsAuthorized(controller.UploadFile)).Methods("POST")
	myRouter.HandleFunc("/api/role/add_role", controller.AddRole).Methods("POST")
	myRouter.HandleFunc("/api/access/add_access", controller.AddAccess).Methods("POST")

	myRouter.HandleFunc("/api/members/add_members", jwtconfig.IsAuthorized(controller.AddUserMember)).Methods("POST")
	myRouter.HandleFunc("/api/members/add_members_access", jwtconfig.IsAuthorized(controller.AddMemberAccess)).Methods("POST")
	myRouter.HandleFunc("/api/members/get_list_members_access", jwtconfig.IsAuthorized(controller.GetListMemberAccess)).Methods("GET")
	myRouter.HandleFunc("/api/members/add_members_role", jwtconfig.IsAuthorized(controller.AddMembersRole)).Methods("POST")
	myRouter.HandleFunc("/api/members/get_list_members_role", jwtconfig.IsAuthorized(controller.GetListMemberRole)).Methods("GET")

	myRouter.HandleFunc("/api/v1/add_product", controller.AddProduct).Methods("POST")
	myRouter.HandleFunc("/api/v1/detail/{id}", jwtconfig.IsAuthorized(controller.GetProductDetailById)).Methods("GET")
	myRouter.HandleFunc("/api/v1/lists", jwtconfig.IsAuthorized(controller.GetProductList)).Methods("GET")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
