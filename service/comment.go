package service

import (
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
