package controllers

import (
	"fmt"
	"gin-blogs/models"
)

type UpdateArticleController struct {
	BaseController
}

func (this *UpdateArticleController) Get() {
	id, _ := this.GetInt("id")
	fmt.Println(id)

	art := models.QueryArticleWithId(id)

	this.Data["Title"] = art.Title
	this.Data["Tags"] = art.Tags
	this.Data["Short"] = art.Short
	this.Data["Content"] = art.Content
	this.Data["Id"] = art.Id

	this.TplName = "write_article.html"
}

func (this *UpdateArticleController) Post() {
	id, _ := this.GetInt("id")
	fmt.Println("postid:", id)

	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	art := models.Article{id, title, tags, short, content, "", 0}
	_, err := models.UpdateArticle(art)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "Update Article Success"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Update Article Invalid"}
	}

	this.ServeJSON()
}
