package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhi-dao/model"
)

func ResponseQuestionID(c *gin.Context, id int) {
	c.JSON(http.StatusOK, gin.H{
		"status":      200,
		"question_id": id,
	})
}

func ResponseAnswerID(c *gin.Context, id int) {
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"answer_id": id,
	})
}

func ResponseQuestionListAndAnswerList(c *gin.Context, questionList model.QuestionList, answerList model.AnswerList) {
	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"question_list": questionList,
		"answer_list":   answerList,
	})
}

func ResponseQuestionAndAnswerList(c *gin.Context, question model.Question, answerList model.AnswerList) {
	c.JSON(http.StatusOK, gin.H{
		"status":      200,
		"question":    question,
		"answer_list": answerList,
	})
}
