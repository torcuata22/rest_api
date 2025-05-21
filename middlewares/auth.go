package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/torcuata22/rest_api/utils"
)

func Authenticate(context *gin.Context) {
	fmt.Println("Authenticate middleware called")
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	// Check and strip Bearer prefix
	const prefix = "Bearer "
	if len(token) <= len(prefix) || token[:len(prefix)] != prefix {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
		return
	}
	token = token[len(prefix):]

	// Verify token
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "error": err.Error()})
		return
	}
	context.Set("userId", userId) //attach data to the context, uses a key (identifier) and the value
	context.Next()
}
