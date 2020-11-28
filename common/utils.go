package common

import (
	"TestKriya/model"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Encrypt use for encypt
func Encrypt(user *model.Users) (*model.Users, error) {
	hashedPassword, err := Hash(user.Data.Password)
	if err != nil {
		return nil, err
	}
	user.Data.Password = string(hashedPassword)
	return user, nil
}

// Hash use for hashing password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// HandleSuccess ...
func HandleSuccess(resp http.ResponseWriter, status int, data interface{}) {
	responses := model.ResponseWrapper{
		Success: true,
		Message: "Success",
		Data:    data,
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)

	err := json.NewEncoder(resp).Encode(responses)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Oopss, something error"))
		fmt.Printf("[HandleSuccess] error when encode data with error : %v \n", err)
	}
}

//HandleError ...
func HandleError(resp http.ResponseWriter, status int, msg string) {
	responses := model.ResponseWrapper{
		Success: false,
		Message: msg,
		Data:    nil,
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)

	err := json.NewEncoder(resp).Encode(responses)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("ooppss, something error"))
		fmt.Printf("[HandleError] error when encode data with error : %v \n", err)
	}
}
