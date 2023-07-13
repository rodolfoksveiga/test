package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodolfoksveiga/k8s-go/config"
	"github.com/rodolfoksveiga/k8s-go/models"
)

func GetAlbums(context *gin.Context) {
	if context.Request.Method != http.MethodGet {
		context.Status(http.StatusMethodNotAllowed)
		return
	}

	albums := []models.Album{}
	config.DB.Model(&models.Album{}).Find(&albums)
	context.JSON(http.StatusOK, &albums)
}

func GetAlbumById(context *gin.Context) {
	if context.Request.Method != http.MethodGet {
		context.Status(http.StatusMethodNotAllowed)
		return
	}

	var album models.Album
	err := config.DB.Model(&models.Album{}).Where("id = ?", context.Param("id")).Take(&album).Error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "AAAAAAAAAA",
		})
		return
	}

	context.JSON(http.StatusOK, &album)
}

func CreateAlbum(context *gin.Context) {
	if context.Request.Method != http.MethodPost {
		context.Status(http.StatusMethodNotAllowed)
		return
	}

	var album models.Album
	err := context.ShouldBindJSON(&album)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Create(&album)
	context.JSON(http.StatusOK, &album)
}

func DeleteAlbum(context *gin.Context) {
	if context.Request.Method != http.MethodDelete {
		context.Status(http.StatusMethodNotAllowed)
		return
	}

	var album models.Album
	err := config.DB.Model(&models.Album{}).Where("id = ?", context.Param("id")).Take(&album).Error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Album not found",
		})
		return
	}

	config.DB.Delete(&album)
	context.JSON(http.StatusOK, &album)
}

func UpdateAlbum(context *gin.Context) {
	if context.Request.Method != http.MethodPut {
		context.Status(http.StatusMethodNotAllowed)
		return
	}

	var album models.Album
	err := config.DB.Model(&models.Album{}).Where("id = ?", context.Param("id")).Take(&album).Error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Album not found",
		})
		return
	}

	err = context.ShouldBindJSON(&album)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Save(&album)
	context.JSON(http.StatusOK, &album)
}
