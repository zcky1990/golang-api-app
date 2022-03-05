package request

type RoleRequest struct {
	RoleName      string `bson:"role_name" json:"role_name"`
	Description   string `bson:"description" json:"description"`
	AccessLevelID string `bson:"access_level_id" json:"access_level_id"`
}
