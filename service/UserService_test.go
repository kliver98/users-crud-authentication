package service

import (
	"testing"
	
	response "authentication/model/response"
	received "authentication/model/received"
)

type RepositoryMock struct {

}

func NewUserRepositoryMock() RepositoryMock {
	return RepositoryMock{}
}

func (repo *RepositoryMock) ValidUserLogin(user received.User) []response.User {
	userResponse := response.User{user.Username, 989654, user.Password, true}
	var users []response.User = []response.User{userResponse}
	return users
}
func TestAuthenticateUser(t *testing.T) {
	//uri := os.Getenv("MONGODB_URL")
	//client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	//client := *mongo.Client{}
	//userRepository := repository.NewUserRepository(client)
	userRepository := NewUserRepositoryMock()
	users, err := userRepository.ValidUserLogin(received.User{"saloar", "989654"} )
	
}
