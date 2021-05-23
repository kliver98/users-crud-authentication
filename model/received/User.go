package model

//Struct that represents a User Object for json response of API
type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
