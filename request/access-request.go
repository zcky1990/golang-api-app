package request

type AccessRequest struct {
	AccessLevel int64    `bson:"access_level" json:"access_level"`
	ListUrl     []string `bson:"list_url" json:"list_url"`
}
