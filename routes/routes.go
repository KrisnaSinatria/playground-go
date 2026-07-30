package routes

import (
	"go-first/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/categories", controllers.GetCategories)
	router.POST("/categories", controllers.CreateCategory)
	router.GET("/categories/:id", controllers.GetCategoryByID)
	router.PUT("/categories/:id", controllers.UpdateCategory)
	router.DELETE("/categories/:id", controllers.DeleteCategory)
}
