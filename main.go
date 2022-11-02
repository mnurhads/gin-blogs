package main

import (
	_ "gin-blogs/routers"
	"github.com/astaxie/beego"
	"gin-blogs/utils"
)

func main() {

	utils.InitMysql()
	beego.Run()
}

