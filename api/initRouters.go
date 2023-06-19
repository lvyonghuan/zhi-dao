package api

import "github.com/gin-gonic/gin"

func InitRouters() {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/register")
		user.GET("/login")
		user.GET("/refresh_token")
	}
	question := r.Group("/question")
	{
		question.POST("/create")
		question.POST("/answer")
		question.GET("/:question_id")
		question.DELETE("/:answer_id") //虽然但是，知乎好像没有去删除提问的功能？从产品角度考虑，删除问题意味着删除回答，这和删帖不一样——回答者某种意义上是付出了成本的。提问者未经商讨就删除问题是对回答者的不尊重。
		question.PUT("/:answer_id")
		question.POST("/:answer_id/like")
	}
	comment := r.Group("/comment")
	{
		comment.POST("/create")
		comment.DELETE("/:comment_id")
		comment.POST("/:comment_id/like")
	}
	r.Run()
}
