package controllers

import (
	"fmt"
	"net/http"
	// "os"
	"go-first/database"
	"go-first/models"
	"go-first/utils"

	"github.com/gin-gonic/gin"
)

func GetRooms(c *gin.Context) {
	var rooms []models.Room

	result := database.DB.Find(&rooms)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, rooms)
}

func CreateRoom(c *gin.Context) {
	name := c.PostForm("name")
	no := c.PostForm("no")

	fileHeader, err := c.FormFile("im")
	fmt.Printf("%+v\n", err)
	// os.Exit(1)

	var imgURL string

	if err == nil {
		url, uploadErr := utils.UploadToCloudinary(fileHeader)
		if uploadErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengunggah gambar ke Cloudinary: " + uploadErr.Error()})
			return
		}
		imgURL = url
	}

	room := models.Room{
		Name: name,
		No:   no,
		Img:  imgURL,
	}

	if err := database.DB.Create(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, room)
}

func GetRoomByID(c *gin.Context) {
	id := c.Param("id")
	var room models.Room

	if err := database.DB.First(&room, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room tidak ditemukan: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, room)
}

func UpdateRoom(c *gin.Context) {
	id := c.Param("id")
	var room models.Room

	if err := database.DB.First(&room, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room tidak ditemukan: " + err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, room)
}

func DeleteRoom(c *gin.Context) {
	id := c.Param("id")
	var room models.Room

	if err := database.DB.First(&room, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room tidak ditemukan: " + err.Error()})
		return
	}

	if err := database.DB.Delete(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room berhasil dihapus"})
}
