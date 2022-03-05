package request

type RoleRequest struct {
	RoleName      string `bson:"role_name" json:"role_name"`
	Description   string `bson:"description" json:"description"`
	AccessLevelID string `bson:"access_level_id" json:"access_level_id"`
}

type RoleRequestWithCompanyId struct {
	RoleName      string `bson:"role_name" json:"role_name"`
	Description   string `bson:"description" json:"description"`
	AccessLevelID string `bson:"access_level_id" json:"access_level_id"`
	CompanyId     string `bson:"company_id" json:"company_id"`
}
