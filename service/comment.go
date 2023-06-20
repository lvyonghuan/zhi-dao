package service

import (
	"errors"
	"zhi-dao/dao"
	"zhi-dao/model"
)

func CreateComment(token, text string, answerID int) (id int, err error) {
	err, uid := checkExp(token, tokenSecret)
	if err != nil {
		return 0, err
	}
	answer, err := dao.SearchAnswerByAnswerID(answerID)
	if err != nil {
		return 0, err
	}
	comment := model.Comment{
		QuestionId:  answer.QuestionId,
		AnswerId:    answerID,
		CommenterId: uid,
		ReplyId:     0,
		Text:        text,
		Like:        0,
	}
	id, err = dao.CreateComment(comment)
	if err != nil {
		return 0, err
	}
	return id, err
}

func ReplyComment(token, text string, replyID int) (commentID int, err error) {
	err, uid := checkExp(token, tokenSecret)
	if err != nil {
		return 0, err
	}
	comment, err := dao.SearchCommentByCommentID(replyID)
	if err != nil {
		return 0, err
	}
	reComment := model.Comment{
		QuestionId:  comment.QuestionId,
		AnswerId:    comment.AnswerId,
		CommenterId: uid,
		ReplyId:     comment.Id,
		Text:        text,
		Like:        0,
	}
	commentID, err = dao.CreateComment(reComment)
	if err != nil {
		return 0, err
	}
	return commentID, nil
}

func DeleteComment(token string, commentID int) (err error) {
	err, uid := checkExp(token, tokenSecret)
	if err != nil {
		return err
	}
	comment, err := dao.SearchCommentByCommentID(commentID)
	if err != nil {
		return err
	}
	if comment.CommenterId != uid {
		return errors.New("用户无权限删除此评论")
	}
	err = dao.DeleteComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func GetCommentList(answerID int) (commentList model.CommentList, err error) {
	commentList, err = dao.GetCommentList(answerID)
	if err != nil {
		return nil, err
	}
	return commentList, nil
}

func GetReplyCommentList(replyID int) (commentList model.CommentList, err error) {
	commentList, err = dao.GetReplyCommentList(replyID)
	if err != nil {
		return nil, err
	}
	return commentList, nil
}
