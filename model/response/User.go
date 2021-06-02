package response

//Struct that represents a User Object for json response of API
type User struct {
	Username string `json:"username" bson:"username"`
	ID       int32  `json:"_id,omitempty" bson:"_id,omitempty"`
	Photo    string `json:"photo,omitempty" bson:"photo,omitempty"`
	Active   bool   `json:"active,omitempty" bson:"active,omitempty"`
}
