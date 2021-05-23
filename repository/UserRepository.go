package repository

import (
	received "authentication/model/received"
	response "authentication/model/response"
	"context"

	errors "authentication/error"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Methods that must have UserRepository
type IUserRepository interface {
	ValidUserLogin(received.User) ([]response.User, error)
}

type userRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) userRepository {
	return userRepository{client: client}
}

func (r userRepository) ValidUserLogin(user received.User) ([]response.User, error) {
	collection := r.client.Database("app").Collection("users")
	found, err := collection.Find(context.Background(), bson.M{"username": user.Username, "password": user.Password})
	if err != nil {
		panic(err)
	}
	defer found.Close(context.Background())
	var users []response.User
	if err = found.All(context.Background(), &users); err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return []response.User{}, &errors.GeneralError{Message: "Credentials not valid"}
	}
	return users, nil
}
