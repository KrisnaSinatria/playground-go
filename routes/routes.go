package routes

import (
	"go-first/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/rooms", controllers.GetRooms)
	router.POST("/rooms", controllers.CreateRoom)
	router.GET("/rooms/:id", controllers.GetRoomByID)
	router.PUT("/rooms/:id", controllers.UpdateRoom)
	router.DELETE("/rooms/:id", controllers.DeleteRoom)
}
