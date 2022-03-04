package request

import (
	m "webappsapi/main/models"
)

type Role = m.Role
type Company = m.Company

type RegisterCompanyRequest struct {
	Company CompanyRequest `json:"company"`
	User    UserAddRequest `json:"user"`
}

type RoleRequest struct {
	Rolename    string `json:"rolename"`
	Description string `json:"description"`
}

type CompanyRequest struct {
	Companyname    string `json:"company_name"`
	Companyaddress string `json:"company_address"`
	Companyemail   string `json:"company_email"`
	Companyphone   string `json:"company_phone"`
}

type UserAddRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Role      Role
	Company   CompanyRequest
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUpRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
