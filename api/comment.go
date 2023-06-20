package api

import (
	"errors"
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

func replyComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	replyIDStr := c.Query("reply_id")
	text := c.PostForm("text")
	replyID, err := strconv.Atoi(replyIDStr)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, "reply id非法")
		return
	}
	id, err := service.ReplyComment(token, text, replyID)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseCommentID(c, id)
}

func deleteComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	commentIDStr := c.Query("comment_id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, "comment id非法")
		return
	}
	err = service.DeleteComment(token, commentID)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseOK(c)
}

func getCommentList(c *gin.Context) {
	answerIDStr := c.Param("answer_id")
	answerID, err := strconv.Atoi(answerIDStr)
	if err != nil {
		resp.NormErr(c, 400, errors.New("answer id非法").Error())
		return
	}
	commentList, err := service.GetCommentList(answerID)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseCommentList(c, commentList)
}

func getReplyCommentList(c *gin.Context) {
	replyIDStr := c.Param("comment_id")
	replyID, err := strconv.Atoi(replyIDStr)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, "reply id非法")
		return
	}
	commentList, err := service.GetReplyCommentList(replyID)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseCommentList(c, commentList)
}
