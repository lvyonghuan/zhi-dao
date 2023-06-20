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

func getUserQuestionListAndAnswerList(c *gin.Context) {
	token := c.GetHeader("Authorization")
	questionList, answerList, err := service.SearchUserQuestionAndAnswer(token)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseQuestionListAndAnswerList(c, questionList, answerList)
}

// 知乎是他人也可以编辑问题。不过暂时没啥精力去做编辑日志。
func changeQuestion(c *gin.Context) {
	token := c.GetHeader("Authorization")
	questionIDStr := c.Param("question_id")
	title := c.PostForm("title")
	introduce := c.PostForm("introduce")
	topic := c.PostForm("topic")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, "question id非法")
		return
	}
	err = service.ChangeQuestion(token, title, introduce, topic, questionID)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseOK(c)
}

func changeAnswer(c *gin.Context) {
	token := c.GetHeader("Authorization")
	answerIDStr := c.Param("answer_id")
	text := c.PostForm("text")
	answerID, err := strconv.Atoi(answerIDStr)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, "answer id非法")
		return
	}
	err = service.ChangeAnswer(token, text, answerID)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseOK(c)
}

func deleteQuestion(c *gin.Context) {
	token := c.GetHeader("Authorization")
	questionIDStr := c.Param("question_id")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, "question id非法")
		return
	}
	err = service.DeleteQuestion(token, questionID)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseOK(c)
}

func deleteAnswer(c *gin.Context) {
	token := c.GetHeader("Authorization")
	answerIDStr := c.Param("answer_id")
	answerID, err := strconv.Atoi(answerIDStr)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, "answer id非法")
		return
	}
	err = service.DeleteAnswer(token, answerID)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseOK(c)
}

func getQuestionAndAnswer(c *gin.Context) {
	questionIDStr := c.Param("question_id")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, "question id非法")
		return
	}
	question, answerList, err := service.GetQuestionAndAnswer(questionID)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseQuestionAndAnswerList(c, question, answerList)
}
