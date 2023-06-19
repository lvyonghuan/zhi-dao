package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func encryption(str string) (encodingStr []byte, err error) {
	encodingStr, err = bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码机密失败：" + err.Error())
	}
	return encodingStr, nil
}

func verifyPassword(password string, encodingPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(encodingPassword, []byte(password))
	if err != nil {
		return errors.New("密码验证失败：" + err.Error())
	}
	return nil
}
