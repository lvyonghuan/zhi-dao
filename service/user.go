package service

import (
	"errors"
	"strconv"
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

func Login(username, password string) (token, refreshToken string, err error) {
	err, user := dao.FindUserByUsername(username)
	if err != nil {
		return "", "", err
	} else if user == (model.User{}) {
		return "", "", errors.New("用户未注册")
	}
	err = verifyPassword(password, []byte(user.Password))
	if err != nil {
		return "", "", err
	}
	token, refreshToken, err = createTokenAndRefreshToken(strconv.Itoa(user.Id))
	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil
}

func RefreshToken(refreshToken string) (token, newRefreshToken string, err error) {
	err, id := checkExp(refreshToken, refreshTokenSecret)
	if err != nil {
		return "", "", err
	}
	token, newRefreshToken, err = createTokenAndRefreshToken(strconv.Itoa(id))
	if err != nil {
		return "", "", err
	}
	return token, newRefreshToken, err
}
