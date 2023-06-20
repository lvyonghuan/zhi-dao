package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"zhi-dao/service"
	"zhi-dao/util/resp"
)

func creatComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	answerIDStr := c.Param("answer_id")
	text := c.PostForm("text")
	answerID, err := strconv.Atoi(answerIDStr)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, "answer id非法")
		return
	}
	id, err := service.CreateComment(token, text, answerID)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseCommentID(c, id)
}
