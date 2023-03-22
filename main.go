package main

// 导入gin包
import (
	"github.com/gin-gonic/gin"
	"go_mvc/routers"
)

func main() {
	router := gin.Default()
	routers.RegisterRouter(router)
	//router.GET("/user/:name", controllers.GetUser)

	router.Run(":9898")
}
