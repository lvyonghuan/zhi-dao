package main

import (
	"zhi-dao/api"
	"zhi-dao/dao"
)

func main() {
	api.InitRouters()
	dao.InitDB()
}
