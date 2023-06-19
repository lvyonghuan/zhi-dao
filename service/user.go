package service

import (
	"errors"
	"zhi-dao/dao"
	"zhi-dao/model"
)

func Register(username, password string) (err error) {
	err, user := dao.FindUserByUsername(username)
	if err != nil {
		return err
	} else if user != (model.User{}) {
		return errors.New("用户名已注册")
	}
	encodingPassword, err := encryption(password) //密码加密
	if err != nil {
		return err
	}
	user.Username = username
	user.Password = string(encodingPassword)
	user.Administrator = false
	err = dao.Register(user) //数据写入数据库
	if err != nil {
		return err
	}
	return nil
}

//func Login(username, password string) (token, refreshToken string, err error) {
//
//}
