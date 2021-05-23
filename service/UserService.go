package service

import (
	"go.mongodb.org/mongo-driver/mongo"

	received "authentication/model/received"
	response "authentication/model/response"
	repository "authentication/repository"
)

//Methods that must have UserService
type IUserService interface {
	AuthenticateUser(client *mongo.Client, userInfo received.User) (response.User, error)
}

func AuthenticateUser(client *mongo.Client, userInfo received.User) (response.User, error) {
	userRepository := repository.NewUserRepository(client)
	users, err := userRepository.ValidUserLogin(userInfo)
	if err != nil {
		return response.User{}, err
	}
	return users[0], nil
}
