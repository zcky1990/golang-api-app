package request

type AccessRequest struct {
	AccessLevel int64    `bson:"access_level" json:"access_level"`
	AccessType  string   `bson:"access_type" json:"access_type"`
	ListUrl     []string `bson:"list_url" json:"list_url"`
}

type AccessRequestWithCompanyId struct {
	AccessLevel int64    `bson:"access_level" json:"access_level"`
	AccessType  string   `bson:"access_type" json:"access_type"`
	ListUrl     []string `bson:"list_url" json:"list_url"`
	CompanyId   string   `bson:"company_id" json:"company_id"`
}
