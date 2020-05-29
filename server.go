package main

import (
	"github.com/ammorteza/clean_architecture/controller"
	router "github.com/ammorteza/clean_architecture/http"
	"github.com/ammorteza/clean_architecture/http/gin"
	"github.com/ammorteza/clean_architecture/repository"
	"github.com/ammorteza/clean_architecture/repository/gorm"
	"github.com/ammorteza/clean_architecture/service"
)

var (
	repo repository.DbRepository = gorm.New()
	appService service.AppService = service.New(repo)
	httpRouter router.Router = gin.New()
	appController controller.AppController = controller.New(appService)
)

func main()  {
	httpRouter.GET("/posts", appController.GetPosts)
	httpRouter.POST("/post/add", appController.AddPost)
	httpRouter.GET("/users", appController.GetUsers)
	httpRouter.POST("/user/add", appController.AddUser)
	httpRouter.SERVE("8080")
}
