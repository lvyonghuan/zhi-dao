package dao

import (
	"errors"
	"zhi-dao/model"
)

func CreateQuestion(question model.Question) (id int, err error) {
	err = DB.Create(&question).Error
	if err != nil {
		return 0, errors.New("数据库错误，问题创建失败：" + err.Error())
	}
	return question.Id, nil
}
