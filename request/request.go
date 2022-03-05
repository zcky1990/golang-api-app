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
	RoleName      string `bson:"role_name" json:"role_name"`
	Description   string `bson:"description" json:"description"`
	AccessLevelID string `bson:"access_level_id" json:"access_level_id"`
}

type CompanyRequest struct {
	CompanyName    string `bson:"company_name" json:"company_name"`
	CompanyAddress string `bson:"company_address" json:"company_address"`
	CompanyEmail   string `bson:"company_email" json:"company_email"`
	CompanyPhone   string `bson:"company_phone" json:"company_phone"`
}

type UserAddRequest struct {
	UserName  string `bson:"username" json:"username"`
	Email     string `bson:"email" json:"email"`
	FirstName string `bson:"firstname" json:"firstname"`
	LastName  string `bson:"lastname" json:"lastname"`
	Password  string `bson:"password" json:"password"`
	Birthday  string `bson:"birthday" json:"birthday"`
	RoleId    string `bson:"role_id" json:"role_id"`
	Company   CompanyRequest
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

type AccessRequest struct {
	AccessLevel int64    `bson:"access_level" json:"access_level"`
	ListUrl     []string `bson:"list_url" json:"list_url"`
}
