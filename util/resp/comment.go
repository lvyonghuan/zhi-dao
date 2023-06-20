package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseCommentID(c *gin.Context, id int) {
	c.JSON(http.StatusOK, gin.H{
		"status":     200,
		"comment_id": id,
	})
}
