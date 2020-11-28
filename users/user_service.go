package users

import (
	"TestKriya/model"
)

//UserService is a interface for function business logic
type UserService interface {
	SignUp(user *model.Users) (*model.Users, error)
	FindAllUser() (*[]model.UserList, error)
	DeleteUser(id string) error
	FindUserByID(id string) (*model.User, error)
	UpdateUser(id string, data *model.Users) (*model.Users, error)
}
