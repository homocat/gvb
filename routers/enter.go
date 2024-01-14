package routers

import (
	"main/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	r.GET("/hello", func (c *gin.Context)  {
		c.String(200, "hello world")
	})
	return r
}