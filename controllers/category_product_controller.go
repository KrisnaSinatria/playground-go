package controllers

import (
	"go-first/database"
	"go-first/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategoryPackages(c *gin.Context) {
	
	var category_packages []models.CategoryPackage
	result := database.DB.Find(&category_packages)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, category_packages)
}

func CreateCategoryPackage(c *gin.Context) {
	var input models.CategoryPackage
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

func GetCategoryPackageByID(c *gin.Context) {
	id := c.Param("id")
	var category_package models.CategoryPackage

	if err := database.DB.First(&category_package, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category_package)
}

func UpdateCategoryPackage(c *gin.Context) {
	id := c.Param("id")
	var category_package models.CategoryPackage

	if err := database.DB.First(&category_package, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&category_package); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&category_package)
	c.JSON(http.StatusOK, category_package)

}

func DeleteCategoryPackage(c *gin.Context) {
	id := c.Param("id")

	var category_package models.CategoryPackage
	if err := database.DB.First(&category_package, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	database.DB.Delete(&category_package)
	c.JSON(http.StatusOK, gin.H{"message": "Category Package berhasil dihapus"})
}
