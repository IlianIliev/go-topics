package users

import "fmt"

type UserService interface {
	CreateUser(input CreateUserInput) error
}

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

type mySQLUserService struct{}

func (service *mySQLUserService) CreateUser(input CreateUserInput) error {
	fmt.Println("This is implementation of Create User")
	return nil
}

func NewUserService() UserService {
	return &mySQLUserService{}
}
