package routes

import (
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/torcuata22/rest_api/models"
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
	context.JSON(http.StatusOK, gin.H{"message": "User logged in!"})
}
