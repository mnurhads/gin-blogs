package controllers

import (
	"github.com/astaxie/beego"
	"gin-blogs/utils"
	"fmt"
	"github.com/matrix/go-matrix/log"
	"time"
	"gin-blogs/models"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

func (this *RegisterController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Println(username, password, repassword)
	log.INFO(username, password, repassword)

	id := models.QueryUserWithUsername(username)
	fmt.Println("id:", id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Testing - testing"}
		this.ServeJSON()
		return
	}

	password = utils.MD5(password)
	fmt.Println("cek password：", password)

	user := models.User{0, username, password, 0, time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Register Success"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "Register Invalid"}
	}
	this.ServeJSON()

}

