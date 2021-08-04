package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marqueswsm/web-api-with-golang/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/", controllers.ShowUser)
		}
		status := main.Group("status")
		{
			status.GET("/", controllers.SendStatus)
		}
	}

	return router
}
