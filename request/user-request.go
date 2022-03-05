package request

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

type UserMembersAddRequest struct {
	UserName  string `bson:"username" json:"username"`
	Email     string `bson:"email" json:"email"`
	FirstName string `bson:"firstname" json:"firstname"`
	LastName  string `bson:"lastname" json:"lastname"`
	Password  string `bson:"password" json:"password"`
	Birthday  string `bson:"birthday" json:"birthday"`
	RoleId    string `bson:"role_id" json:"role_id"`
	CompanyId string `bson:"company_id" json:"company_id"`
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
