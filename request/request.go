package request

import (
	m "webappsapi/main/models"
)

type Role = m.Role
type Company = m.Company
type AccessLevel = m.AccessLevel

type RegisterCompanyRequest struct {
	Company CompanyRequest `json:"company"`
	User    UserAddRequest `json:"user"`
}

type RoleRequest struct {
	RoleName    string        `json:"role_name"`
	Description string        `json:"description"`
	Access      []AccessLevel `json:"access"`
}

type CompanyRequest struct {
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	CompanyEmail   string `json:"company_email"`
	CompanyPhone   string `json:"company_phone"`
}

type UserAddRequest struct {
	UserName   string `json:"username"`
	Email      string `json:"email"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Password   string `json:"password"`
	Birthday   string `json:"birthday"`
	AccessType string `json:"access_type"`
	Role       Role
	Company    CompanyRequest
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUpRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
