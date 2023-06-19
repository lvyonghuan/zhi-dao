package dao

import (
	"errors"
	"gorm.io/gorm"
	"zhi-dao/model"
)

func FindUserByUsername(username string) (err error, user model.User) {
	err = DB.Where("username=?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.User{}
	} else if err != nil {
		return errors.New("查找用户错误：" + err.Error()), model.User{}
	}
	return nil, user
}

func Register(user model.User) (err error) {
	err = DB.Create(&user).Error
	if err != nil {
		errors.New("数据库注册失败：" + err.Error())
	}
	return nil
}

func Login(username string) (user model.User, err error) {
	err = DB.Where("name=?", username).First(&user).Error
	return user, err
}

func FindUserByUid(uid int) (user model.User, err error) {
	err = DB.Where("id=?", uid).First(&user).Error
	return user, err
}
