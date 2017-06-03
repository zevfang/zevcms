package controllers

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "default/home.html"
}
