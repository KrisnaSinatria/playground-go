package controllers

import (
	"fmt"
	"go-first/database"
	"go-first/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPackages(c *gin.Context) {
	var packages []models.Package

	result := database.DB.Preload("CategoryPackage").Find(&packages)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, packages)
}

func CreatePackage(c *gin.Context) {
	var input models.Package

	err := c.ShouldBindJSON(&input)
	fmt.Printf("ERROR BIND: %+v\n", err)
	fmt.Printf("INPUT SETELAH BIND: %+v\n", input)
	

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("CategoryPackage").First(&input, input.ID)

	c.JSON(http.StatusCreated, input)
}

func GetPackageByID(c *gin.Context) {
	id := c.Param("id")
	var pkg models.Package

	if err := database.DB.Preload("CategoryPackage").First(&pkg, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package tidak ditemukan: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, pkg)
}

func UpdatePackage(c *gin.Context) {
	id := c.Param("id")
	var pkg models.Package

	if err := database.DB.First(&pkg, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package tidak ditemukan: " + err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&pkg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&pkg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("CategoryPackage").First(&pkg, pkg.ID)

	c.JSON(http.StatusOK, pkg)
}

func DeletePackage(c *gin.Context) {
	id := c.Param("id")
	var pkg models.Package

	if err := database.DB.First(&pkg, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package tidak ditemukan: " + err.Error()})
		return
	}

	if err := database.DB.Delete(&pkg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package berhasil dihapus"})
}
