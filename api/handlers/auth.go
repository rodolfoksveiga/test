package handlers

import (
	"html"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rodolfoksveiga/k8s-go/config"
	"github.com/rodolfoksveiga/k8s-go/models"
	"github.com/rodolfoksveiga/k8s-go/utils"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserInput struct {
	Email    string
	Password string
}

func RegisterUser(context *gin.Context) {
	if context.Request.Method != http.MethodPost {
		context.Status(http.StatusMethodNotAllowed)
		return
	}

	var input models.User
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Error to encrypt password",
		})
		return
	}

	var user models.User
	user.Email = html.EscapeString(strings.TrimSpace(input.Email))
	user.Password = string(hashedPassword)

	err = config.DB.Create(&user).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to save user",
		})
		return
	}

	context.Status(http.StatusOK)
}

func LoginUser(context *gin.Context) {
	if context.Request.Method != http.MethodPost {
		context.Status(http.StatusMethodNotAllowed)
		return
	}

	var input models.User
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	err = config.DB.Model(models.User{}).Where("email = ?", input.Email).Take(&user).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong email or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong email or password",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong email or password",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}

func GetCurrentUser(context *gin.Context) {
	if context.Request.Method != http.MethodGet {
		context.Status(http.StatusMethodNotAllowed)
		return
	}

	userID, err := utils.ExtractTokenID(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid token",
		})
		return
	}

	var user models.User

	err = config.DB.Model(models.User{}).Where("id = ?", userID).Take(&user).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	user.Password = ""

	context.JSON(http.StatusOK, &user)
}
