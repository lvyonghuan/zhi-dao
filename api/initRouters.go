package api

import "github.com/gin-gonic/gin"

func InitRouters() {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/register", register)
		user.GET("/login", login)
		user.GET("/refresh_token", refreshToken)
	}
	question := r.Group("/question")
	{
		question.POST("/create", createQuestion)
		question.POST("/answer", createAnswer)
		question.GET("/look/:question_id")
		question.GET("/my", getUserQuestionListAndAnswerList) //获取自己的所有问题、回答
		question.PUT("/change_question/:question_id", changeQuestion)
		question.DELETE("/delete_answer/:answer_id", deleteAnswer)
		question.DELETE("/delete_question/:question_id", deleteQuestion) //在问题没有回答的前提下
		question.PUT("/change_answer/:answer_id", changeAnswer)
		question.POST("/:answer_id")
	}
	comment := r.Group("/comment")
	{
		comment.POST("/create")
		comment.DELETE("/:comment_id")
		comment.POST("/:comment_id")
	}
	r.Run()
}
