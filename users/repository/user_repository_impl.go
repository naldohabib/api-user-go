package repository

import (
	"TestKriya/model"
	"TestKriya/users"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// UserRepoImpl is use sharing connection
type UserRepoImpl struct {
	db *gorm.DB
}

// UpdateUser used to update 1 user data
func (u UserRepoImpl) UpdateUser(id string, data *model.Users) (*model.Users, error) {
	err := u.db.Model(&data).Where("id = ?", id).Update(data).Error
	if err != nil {
		return nil, fmt.Errorf("UserRepoImpl.Update Error when query update data with error: %w", err)
	}
	return data, nil
}

// FindUserByID used to find one  user data
func (u UserRepoImpl) FindUserByID(id string) (*model.User, error) {
	dataUser := new(model.User)

	if err := u.db.Table("users").Where("id = ?", id).First(&dataUser).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("ERROR: Error no data user with id you entered")
	}

	return dataUser, nil
}

// DeleteUser used to delete user data
func (u UserRepoImpl) DeleteUser(id string) error {
	data := model.User{}
	err := u.db.Table("users").Where("id = ?", id).Delete(&data).Error
	if err != nil {
		logrus.Error(err)
		return errors.New("ERROR: Error when delete data user")
	}

	return nil
}

// FindAllRole use when you want to show all user data
func (u UserRepoImpl) FindAllRole() (*[]model.Rolee, error) {
	var role []model.Role

	if err := u.db.Table("roles").Find(&role).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("ERROR: Error when get all data users")
	}

	var result []model.Rolee
	for i := 0; i < len(role); i++ {
		roleWrap := model.Rolee{
			ID:        role[i].ID,
			CreatedAt: role[i].CreatedAt,
			UpdatedAt: role[i].UpdatedAt,
			DeletedAt: role[i].DeletedAt,
		}
		err := json.Unmarshal([]byte(role[i].Data), &roleWrap.Data)
		if err != nil {
			fmt.Println("Error unmarhal")
		}
		result = append(result, roleWrap)
	}

	return &result, nil
}

// FindAllUser use when you want to show all user data
func (u UserRepoImpl) FindAllUser() (*[]model.Users, error) {
	var user []model.User

	if err := u.db.Table("users").Find(&user).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("ERROR: Error when get all data users")
	}
	//fmt.Println(user)

	var result []model.Users

	for i := 0; i < len(user); i++ {
		var userWrap model.UserWrap
		err := json.Unmarshal([]byte(user[i].Data), &userWrap)
		if err != nil {
			fmt.Println("Eroor")
		}

		var wrap = model.Users{
			ID:        user[i].ID,
			RoleId:    user[i].RoleId,
			CreatedAt: user[i].CreatedAt,
			UpdatedAt: user[i].UpdatedAt,
			DeletedAt: user[i].DeletedAt,
			Data:      userWrap,
		}
		result = append(result, wrap)
	}

	//fmt.Println(userWrap)

	return &result, nil
}

// SignUp use for create a new account user
func (u UserRepoImpl) SignUp(user *model.Users) (*model.Users, error) {
	output2, _ := json.Marshal(user.Data)

	var userdata = model.User{
		ID:        user.ID,
		Data:      string(output2),
		RoleId:    user.RoleId,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	err := u.db.Table("users").Save(&userdata).Error
	//fmt.Println(string(output2))
	if err != nil {
		fmt.Errorf("[UserRepoImpl.SignUp] Error when trying to create use error is : %v\n", err)
		return nil, err
	}
	return user, nil
}

func CreateRepoImpl(db *gorm.DB) users.UserRepo {
	return &UserRepoImpl{db}
}
