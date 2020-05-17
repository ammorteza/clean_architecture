package controller

import (
	"encoding/json"
	"fmt"
	"github.com/ammorteza/clean_architecture/entity"
	service2 "github.com/ammorteza/clean_architecture/service"
	"net/http"
)

type controller struct {}

type PostController interface {
	AddPost(w http.ResponseWriter, r *http.Request)
	GetPosts(w http.ResponseWriter, r *http.Request)
}

var (
	postService service2.PostService
)

func NewPostController(service service2.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller)AddPost(w http.ResponseWriter, r *http.Request){
	var input entity.Post

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(input)

	if err := postService.Validate(&input); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := postService.Create(&input)
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

func (*controller)GetPosts(w http.ResponseWriter, r *http.Request){
	res, err := postService.FetchAll()
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
