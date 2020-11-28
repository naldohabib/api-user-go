package handler

import (
	"TestKriya/common"
	"TestKriya/model"
	"TestKriya/users"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

//UserHandler struct use for get funcntion in business logic
type UserHandler struct {
	userService users.UserService
}

func (h UserHandler) update(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id := pathVar["id"]

	_, err := h.userService.FindUserByID(id)
	if err != nil {
		fmt.Printf("[UserHandler.update] Error when check id to usecase with error: %v\n", err)
		common.HandleError(resp, http.StatusBadRequest, "ID DOES NOT EXIST")
		return
	}

	var user = model.Users{}

	err = json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		common.HandleError(resp, http.StatusInternalServerError, "Oopss, something error")
		fmt.Printf("[UserHandler.getData] Error when decode data with error: %v\n", err)
		return
	}

	//fmt.Println(&user)
	//fmt.Println(user)

	dataUpdate, err := h.userService.UpdateUser(id, &user)
	if err != nil {
		common.HandleError(resp, http.StatusInternalServerError, err.Error())
		fmt.Printf("[UserHandler.update] Error when send data to usecase with error : %v", err)
		return
	}

	common.HandleSuccess(resp, http.StatusOK, dataUpdate)

}

func (h UserHandler) findUserByID(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id := pathVar["id"]

	data, err := h.userService.FindUserByID(id)
	if err != nil {
		common.HandleError(resp, http.StatusInternalServerError, "User ID not Found!")
		fmt.Printf("[UserHandler.getByID]Error when request data with error : %v \n", err)
		return
	}

	common.HandleSuccess(resp, http.StatusOK, data)
}

func (h UserHandler) delete(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	key := pathVar["id"]

	err := h.userService.DeleteUser(key)
	if err != nil {
		common.HandleError(resp, http.StatusNoContent, "Oppss, something error")
		fmt.Printf("[UserHandler.delete]Error when request data to usecase with error : %v\n", err)
		return
	}

	common.HandleSuccess(resp, http.StatusOK, nil)
}

func (h UserHandler) findAll(resp http.ResponseWriter, req *http.Request) {
	listUser, err := h.userService.FindAllUser()
	if err != nil {
		common.HandleError(resp, http.StatusBadRequest, err.Error())
		logrus.Error(err)
		return
	}
	common.HandleSuccess(resp, http.StatusOK, listUser)
	return
}

func (h UserHandler) signup(resp http.ResponseWriter, req *http.Request) {
	user := new(model.Users)
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		common.HandleError(resp, http.StatusInternalServerError, "Oppss, something error")
		fmt.Printf("[UserHandler.signup] Error when decoder data from body with error : %v\n", err)
		return
	}

	uid, _ := uuid.NewRandom()
	user.ID = uid.String()
	//fmt.Println(user)

	response, err := h.userService.SignUp(user)
	if err != nil {
		common.HandleError(resp, http.StatusInternalServerError, err.Error())
		fmt.Printf("[UserHandler.InsertData] Error when request data to usecase with error: %v\n", err)
		return
	}

	common.HandleSuccess(resp, http.StatusOK, response)
}

// CreateUserHandler use for handling request
func CreateUserHandler(r *mux.Router, userService users.UserService) {
	userHandler := UserHandler{userService}

	r.HandleFunc("/api/v1/signup", userHandler.signup).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/findAll", userHandler.findAll).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/delete/{id}", userHandler.delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/findUserByID/{id}", userHandler.findUserByID).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/update/{id}", userHandler.update).Methods(http.MethodPut)

}
