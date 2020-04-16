package main

import (
	_ "admin-api/routers"
	_ "admin-api/task"
	"admin-api/tests"

	"github.com/astaxie/beego"
)

func main() {
	go tests.Test()
	beego.Run()
}
