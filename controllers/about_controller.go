package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {

	c.Data["wa"] = "No：0882005424842"
	c.Data["qq"] = "QQ：123456789"
	c.Data["tel"] = "Tel：13167582311"

	c.TplName = "aboutme.html"
}
