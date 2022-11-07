package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gin-blogs/utils"
	"gin-blogs/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:", username, ",password:", password)

	id := models.QueryUserWithParam(username, utils.MD5(password))
	fmt.Println("id:", id)
	if id > 0 {
		this.SetSession("loginuser", username)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": " Login Success"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": " Login Invalid"}
	}
	this.ServeJSON()
}

