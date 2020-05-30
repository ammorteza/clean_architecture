package controller

import (
	"encoding/json"
	"github.com/ammorteza/clean_architecture/entity"
	"net/http"
)

type postController interface {
	AddPost(w http.ResponseWriter, r *http.Request)
	GetPosts(w http.ResponseWriter, r *http.Request)
}

func (c *controller)AddPost(w http.ResponseWriter, r *http.Request){
	var input entity.Post

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := c.service.IsValidPost(&input); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	//////////////// begin transaction ///////////////////////
	tx, err := c.service.BeginTx()
	if err != nil{
		panic(err.Error())
	}
	txService := c.service.WithTx(tx)
	defer func() {
		if r := recover(); r != nil{
			txService.RollbackTx()
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(r.(string)))
			return
		}
	}()

	var userInfo = entity.User{ID: input.UserId}
	if err := txService.FirstUser(&userInfo); err != nil{
		panic(err.Error())
	}

	userInfo.CommentCount += 1
	if err := txService.UpdateUser(&userInfo); err != nil{
		panic(err.Error())
	}

	if err := txService.CreatePost(&input); err != nil{
		panic(err.Error())
	}

	if err := txService.CommitTx(); err != nil{
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
}

func (c *controller)GetPosts(w http.ResponseWriter, r *http.Request){
	res, err := c.service.FetchPosts()
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
