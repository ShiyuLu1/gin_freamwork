package routers
import (
	"go_mvc/controllers"
	"github.com/gin-gonic/gin"
)
//路由设置
func RegisterRouter(router *gin.Engine) {
	routerUser(router)
}
//用户路由
func routerUser(engine *gin.Engine) {
	var group = engine.Group("/api/user")
	{
		group.GET("/getAll", controllers.GetUsers)
	}
}
