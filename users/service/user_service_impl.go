package service

import (
	"TestKriya/common"
	"TestKriya/model"
	"TestKriya/users"
	"github.com/pkg/errors"
)

// UserServiceImpl use for get a repo connection
type UserServiceImpl struct {
	userRepo users.UserRepo
}

func (u UserServiceImpl) UpdateUser(id string, data *model.Users) (*model.Users, error) {
	_, err := u.userRepo.FindUserByID(id)
	if err != nil {
		return nil, errors.New("userID does not exist")
	}

	user, err := u.UpdateUser(id, data)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (u UserServiceImpl) FindUserByID(id string) (*model.User, error) {
	return u.userRepo.FindUserByID(id)
}

func (u UserServiceImpl) DeleteUser(id string) error {
	err := u.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (u UserServiceImpl) FindAllUser() (*[]model.UserList, error) {
	role, err := u.userRepo.FindAllRole()
	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.FindAllUser()
	if err != nil {
		return nil, err
	}

	var data []model.UserList

	for i := 0; i < len(*user); i++ {
		for k := 0; k < len(*role); k++ {
			if (*user)[i].RoleId == (*role)[k].ID {
				var dataResult = model.UserList{
					UserID:   (*user)[i].ID,
					Username: (*user)[i].Data.Username,
					Email:    (*user)[i].Data.Email,
					RoleName: (*role)[k].Data.RoleName,
				}
				data = append(data, dataResult)
			}
		}
	}

	return &data, nil
}

func (u UserServiceImpl) SignUp(user *model.Users) (*model.Users, error) {
	data, err := common.Encrypt(user)
	if err != nil {
		return nil, err
	}

	user, err = u.userRepo.SignUp(data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUserService use for get connection from repository
func CreateUserServiceImpl(userRepo users.UserRepo) users.UserService {
	return &UserServiceImpl{userRepo}
}
