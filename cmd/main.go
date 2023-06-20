package main

import (
	"zhi-dao/api"
	"zhi-dao/dao"
)

func main() {
	dao.InitDB()
	dao.InitRedis()
	api.InitRouters()
}
