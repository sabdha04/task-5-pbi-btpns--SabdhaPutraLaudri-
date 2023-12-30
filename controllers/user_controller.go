package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"FinalProject/database"
	"FinalProject/models"
	"FinalProject/helpers"
	"FinalProject/middlewares"
	"golang.org/x/crypto/bcrypt"
)

func verifyPassword(savedPassword, enteredPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(savedPassword), []byte(enteredPassword))
	return err == nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func RegisterUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser.Password = hashedPassword

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newUser})
}


func LoginUser(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	
	if !verifyPassword(user.Password, loginData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	
	
	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	
	
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("userId")
	var updatedUser models.User

	middlewares.JWTAuthMiddleware()(c)

	if val, exists := c.Get("error"); exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": val})
		return
	}

	if err := database.DB.First(&updatedUser, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updatedUser.Password != "" {
		hashedPassword, err := HashPassword(updatedUser.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		updatedUser.Password = hashedPassword
	}

	database.DB.Save(&updatedUser)

	c.JSON(http.StatusOK, gin.H{"data": updatedUser})
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("userId")
	var user models.User

	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
