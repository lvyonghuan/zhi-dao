package api

import (
	"github.com/gin-gonic/gin"
	"log"
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
