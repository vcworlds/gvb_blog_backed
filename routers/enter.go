package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(200, "nihao")
	})
	return router
}
