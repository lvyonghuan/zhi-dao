package dao

import (
	"errors"
	"gorm.io/gorm"
	"zhi-dao/model"
)

func CreateQuestion(question model.Question) (id int, err error) {
	err = DB.Create(&question).Error
	if err != nil {
		return 0, errors.New("数据库错误，问题创建失败：" + err.Error())
	}
	return question.Id, nil
}

func CreateAnswer(answer model.Answer) (id int, err error) {
	err = DB.Create(&answer).Error
	if err != nil {
		return 0, errors.New("数据库错误，回答创建失败：" + err.Error())
	}
	return answer.Id, nil
}

func SearchQuestionByQuestionID(id int) (question model.Question, err error) {
	err = DB.Where("id=?", id).First(&question).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Question{}, nil
	} else if err != nil {
		return model.Question{}, errors.New("查找用户错误：" + err.Error())
	}
	return question, nil
}
