package controllers

import (
	"gin-blogs/models"
	"github.com/astaxie/beego"
)

type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		beego.BeeLogger.Error(err.Error())
	}
	this.Data["Album"] = albums
	this.TplName = "album.html"
	//this.TplName = "album.html"
}
