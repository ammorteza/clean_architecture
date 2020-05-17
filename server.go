package main

import (
	"github.com/ammorteza/clean_architecture/controller"
	router "github.com/ammorteza/clean_architecture/http"
	"github.com/ammorteza/clean_architecture/repository"
	"github.com/ammorteza/clean_architecture/service"
)

var (
	repo repository.PostRepository = repository.NewMysqlRepository()
	postService service.PostService = service.NewPostService(repo)
	httpRouter router.Router = router.NewGinRouter()
	postController controller.PostController = controller.NewPostController(postService)
)

func main()  {
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/post/add", postController.AddPost)
	httpRouter.SERVE("8080")
}
