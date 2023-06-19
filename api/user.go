package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"zhi-dao/service"
	"zhi-dao/util/resp"
)

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		resp.NormErr(c, 514, errors.New("用户名、密码不能为空").Error())
	}
	err := service.Register(username, password)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseOK(c)
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		resp.NormErr(c, 514, errors.New("用户名、密码不能为空").Error())
	}
	token, refreshToken, err := service.Login(username, password)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseTokenOK(c, token, refreshToken)
}

func refreshToken(c *gin.Context) {
	refreshToken := c.GetHeader("Authorization")
	token, newRefreshToken, err := service.RefreshToken(refreshToken)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseTokenOK(c, token, newRefreshToken)
}
