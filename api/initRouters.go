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
		question.GET("/:question_id", getQuestionAndAnswer)   //查看问题与问题下所有回答。不需要登陆。
		question.GET("/my", getUserQuestionListAndAnswerList) //获取自己的所有问题、回答
		question.PUT("/change_question/:question_id", changeQuestion)
		question.DELETE("/delete_answer/:answer_id", deleteAnswer)
		question.DELETE("/delete_question/:question_id", deleteQuestion) //在问题没有回答的前提下
		question.PUT("/change_answer/:answer_id", changeAnswer)
		question.POST("/like", likeAnswer)
	}
	comment := r.Group("/comment")
	{
		comment.POST("/:answer_id", creatComment)
		comment.POST("/reply", replyComment)
		comment.DELETE("/:comment_id", deleteComment)
		comment.GET("/:answer_id", getCommentList)
		comment.GET("/reply/:comment_id", getReplyCommentList)
		comment.POST("/like")
	}
	r.Run()
}
