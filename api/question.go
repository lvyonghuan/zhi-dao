package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"zhi-dao/service"
	"zhi-dao/util/resp"
)

func createQuestion(c *gin.Context) {
	token := c.GetHeader("Authorization")
	title := c.PostForm("title")
	introduce := c.PostForm("introduce")
	topic := c.PostForm("topic")
	id, err := service.CreateQuestion(token, title, introduce, topic)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseQuestionID(c, id)
}

func createAnswer(c *gin.Context) {
	token := c.GetHeader("Authorization")
	questionIDStr := c.Query("question_id")
	text := c.PostForm("text")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, "question id非法")
		return
	}
	id, err := service.CreateAnswer(token, text, questionID)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseAnswerID(c, id)
}
