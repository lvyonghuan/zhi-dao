package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NormErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}

func ResponseOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "success",
	})
}

func ResponseTokenOK(c *gin.Context, token, refreshToken string) {
	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"token":         token,
		"refresh_token": refreshToken,
	})
}
