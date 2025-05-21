package routes

import (
	"fmt"
	"net/http"

	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/torcuata22/rest_api/models"
	"github.com/torcuata22/rest_api/utils"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse user"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user) //populate struct with data from request
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse user"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Could not authenticate user credentials"})
		return
	}

	//generate JWT
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		fmt.Println("Error generating token:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User logged in!", "token": token})
}
