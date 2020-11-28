package users

import (
	"TestKriya/model"
)

// UserRepo is a interface for function repository
type UserRepo interface {
	SignUp(user *model.Users) (*model.Users, error)
	FindAllUser() (*[]model.Users, error)
	FindAllRole() (*[]model.Rolee, error)
	DeleteUser(id string) error
	FindUserByID(id string) (*model.User, error)
	UpdateUser(id string, data *model.Users) (*model.Users, error)
}
