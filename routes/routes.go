package routes

import (
	"go-first/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/category-packages", controllers.GetCategoryPackages)
	router.POST("/category-packages", controllers.CreateCategoryPackage)
	router.GET("/category-packages/:id", controllers.GetCategoryPackageByID)
	router.PUT("/category-packages/:id", controllers.UpdateCategoryPackage)
	router.DELETE("/category-packages/:id", controllers.DeleteCategoryPackage)
}
