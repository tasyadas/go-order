package routers

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/order", controllers.GetData)
	router.POST("/order", controllers.CreateOrder)
	router.PUT("/order", controllers.UpdateData)
	router.DELETE("/order", controllers.DeleteData)

	return router
}
