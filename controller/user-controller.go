package controller

import (
	"encoding/json"
	"fmt"
	"github.com/ammorteza/clean_architecture/entity"
	"net/http"
)

type userController interface {
	AddUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
}

func (c controller)AddUser(w http.ResponseWriter, r *http.Request){
	var input entity.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(input)

	if err := c.service.IsValidUser(&input); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err := c.service.RegisterUser(&input)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *controller)GetUsers(w http.ResponseWriter, r *http.Request){
	res, err := c.service.FetchUsers()
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	js, err := json.Marshal(res)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(js)
}