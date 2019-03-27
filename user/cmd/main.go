package main

import (
	"book/user/cmd/service"
	"book/utils/db"
)

func main() {
	//实例化数据库

	db.New("mysql","root:root@/test?charset=utf8")

	service.Run()
}
