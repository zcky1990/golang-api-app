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
	myRouter.HandleFunc("/api/users/add_user_members", jwtconfig.IsAuthorized(controller.AddUser)).Methods("POST")
	myRouter.HandleFunc("/api/image/upload", controller.UploadFile).Methods("POST")
	myRouter.HandleFunc("/api/role/add_role", controller.AddRole).Methods("POST")
	myRouter.HandleFunc("/api/users/company_register", controller.RegisterCompany).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
