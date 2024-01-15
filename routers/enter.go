package routers

import (
	"main/global"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	// 分组
	apiRouterGroup := router.Group("api")
	
	// 系统配置 api
	routerGroupApp := RouterGroup{ apiRouterGroup }
	routerGroupApp.SettingsRouter()

	return router
}