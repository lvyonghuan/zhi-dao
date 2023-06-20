package dao

import (
	"errors"
	"zhi-dao/model"
)

func CreateComment(comment model.Comment) (id int, err error) {
	err = DB.Create(&comment).Error
	if err != nil {
		return 0, errors.New("创建评论错误：" + err.Error())
	}
	return comment.Id, nil
}
