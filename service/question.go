package service

import (
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
