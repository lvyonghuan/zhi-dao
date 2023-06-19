package service

import (
	"errors"
	"zhi-dao/dao"
	"zhi-dao/model"
)

func CreateQuestion(token, title, introduce, topic string) (questionID int, err error) {
	err, id := checkExp(token, tokenSecret)
	if err != nil {
		return 0, err
	}
	question := model.Question{
		QuestionerId: id,
		Title:        title,
		Introduce:    introduce,
		Topic:        topic,
	}
	questionID, err = dao.CreateQuestion(question)
	if err != nil {
		return 0, err
	}
	return questionID, nil
}

func CreateAnswer(token, text string, questionID int) (answerID int, err error) {
	err, id := checkExp(token, tokenSecret)
	if err != nil {
		return 0, err
	}

	temp, err := dao.SearchQuestionByQuestionID(questionID)
	if err != nil {
		return 0, err
	} else if temp == (model.Question{}) {
		return 0, errors.New("没有该问题")
	}

	answer := model.Answer{
		QuestionId: questionID,
		AnswererId: id,
		Text:       text,
		Like:       0,
	}
	answerID, err = dao.CreateAnswer(answer)
	if err != nil {
		return 0, err
	}
	return answerID, err
}

func SearchUserQuestionAndAnswer(token string) (questionList model.QuestionList, answerList model.AnswerList, err error) {
	err, id := checkExp(token, tokenSecret)
	if err != nil {
		return nil, nil, err
	}
	questionList, err = dao.SearchQuestionListByUserID(id)
	if err != nil {
		return nil, nil, err
	}
	answerList, err = dao.SearchAnswerListByUserID(id)
	if err != nil {
		return nil, nil, err
	}
	return questionList, answerList, nil
}

func ChangeQuestion(token, title, introduce, topic string, questionID int) (err error) {
	err, _ = checkExp(token, tokenSecret) //TODO:利用uid编写问题修改日志
	if err != nil {
		return err
	}
	question, err := dao.SearchQuestionByQuestionID(questionID)
	if err != nil {
		return err
	} else if question == (model.Question{}) {
		return errors.New("没有该问题")
	}
	question.Title = title
	question.Introduce = introduce
	question.Topic = topic
	err = dao.ChangeQuestion(question)
	if err != nil {
		return err
	}
	return nil
}

func ChangeAnswer(token, text string, answerID int) (err error) {
	err, id := checkExp(token, tokenSecret)
	if err != nil {
		return err
	}
	answer, err := dao.SearchAnswerByAnswerID(answerID)
	if err != nil {
		return err
	}
	if answer.AnswererId != id {
		return errors.New("用户无权限修改该回答")
	}
	answer.Text = text
	err = dao.ChangeAnswer(answer)
	if err != nil {
		return err
	}
	return nil
}
