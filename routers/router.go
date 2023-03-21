package routers

import (
	"github.com/gin-gonic/gin"
	"go_mvc/controllers"
)

//路由设置
func RegisterRouter(router *gin.Engine) {
	routerUser(router)
	routerStructureEvent(router)
}

//用户路由
func routerUser(engine *gin.Engine) {
	var group = engine.Group("/api/user")
	{
		group.GET("/getAll", controllers.GetUsers)
	}
}

func routerStructureEvent(engine *gin.Engine) {
	var group = engine.Group("/api/events")
	{
		group.GET("/getEvents", controllers.GetEvents())
	}
}
