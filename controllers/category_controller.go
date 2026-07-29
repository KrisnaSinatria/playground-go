package controllers

import (
    "go-first/database"
    "go-first/models"
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
    var categories []models.Category
    database.DB.Find(&categories)
    c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
    var input models.Category
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&input)
    c.JSON(http.StatusCreated, input)
}