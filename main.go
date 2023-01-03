package main

import (
	"go-blog/model"
	"go-blog/router"
)

func main() {
	//引入数据库
	model.InitDb()
	router.InitRouter()
}
