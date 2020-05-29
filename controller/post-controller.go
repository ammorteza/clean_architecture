package controller

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(input)

	if err := c.service.IsValidPost(&input); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err := c.service.CreatePost(&input)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	//js, err := json.Marshal(res)
	//if err != nil{
	//	w.WriteHeader(http.StatusInternalServerError)
	//	w.Write([]byte(err.Error()))
	//}
	//w.Write(js)
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
