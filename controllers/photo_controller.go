package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"FinalProject/database"
	"FinalProject/models"
	"FinalProject/helpers"
	"github.com/go-playground/validator/v10"
)

func CreatePhoto(c *gin.Context) {
	var newPhoto models.Photo
	if err := c.ShouldBindJSON(&newPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(newPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&newPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newPhoto})
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	if err := database.DB.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func UpdatePhoto(c *gin.Context) {
	photoID := c.Param("photoId")
	var updatedPhoto models.Photo

	claims, err := helpers.VerifyToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
		return
	}
	userIDFromToken := claims.UserID

	if err := database.DB.First(&updatedPhoto, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if updatedPhoto.UserID != userIDFromToken {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: You don't have permission to update this photo"})
		return
	}

	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&updatedPhoto)

	c.JSON(http.StatusOK, gin.H{"data": updatedPhoto})
}

func DeletePhoto(c *gin.Context) {
	photoID := c.Param("photoId")
	var photo models.Photo

	claims, err := helpers.VerifyToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
		return
	}
	userIDFromToken := claims.UserID

	if err := database.DB.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if photo.UserID != userIDFromToken {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: You don't have permission to delete this photo"})
		return
	}

	database.DB.Delete(&photo)

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
