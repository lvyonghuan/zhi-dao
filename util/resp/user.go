package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseTokenOK(c *gin.Context, token, refreshToken string) {
	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"token":         token,
		"refresh_token": refreshToken,
	})
}
