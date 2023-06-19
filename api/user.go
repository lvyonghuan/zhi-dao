package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"zhi-dao/service"
	"zhi-dao/util/resp"
)

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	err := service.Register(username, password)
	if err != nil {
		log.Println(err)
		resp.NormErr(c, 400, err.Error())
		return
	}
	resp.ResponseOK(c)
}
