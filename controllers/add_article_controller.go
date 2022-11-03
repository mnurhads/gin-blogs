package controllers

import (
	"fmt"
	"time"
	"gin-blogs/models"
)

type AddArticleController struct {
	BaseController
}

func (this *AddArticleController) Get() {

	this.TplName = "add_article.html"
}

func (this *AddArticleController) Post() {

	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	art := models.Article{0, title, tags, short, content, "testing", time.Now().Unix()}
	_, err := models.AddArticle(art)

	var response map[string]interface{}

	if err == nil {
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}

	this.Data["json"] = response
	this.ServeJSON()
}
