package gin

import (
	"github.com/ammorteza/clean_architecture/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ginRouter struct {}

func New() router.Router {
	return &ginRouter{}
}

var  (
	ginDispather = gin.Default()
)

func (*ginRouter)GET(uri string, f func(w http.ResponseWriter, r *http.Request)){
	ginDispather.GET(uri, gin.WrapF(f))
}

func (*ginRouter)POST(uri string, f func(w http.ResponseWriter, r *http.Request)){
	ginDispather.POST(uri, gin.WrapF(f))
}

func (*ginRouter)SERVE(port string){
	ginDispather.Run(":8080")
}