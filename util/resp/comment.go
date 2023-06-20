package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhi-dao/model"
)

func ResponseCommentID(c *gin.Context, id int) {
	c.JSON(http.StatusOK, gin.H{
		"status":     200,
		"comment_id": id,
	})
}

func ResponseCommentList(c *gin.Context, commentList model.CommentList) {
	c.JSON(http.StatusOK, gin.H{
		"status":       200,
		"comment_list": commentList,
	})
}
