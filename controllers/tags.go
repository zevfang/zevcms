package controllers

type TagsController struct {
	BaseController
}

func (this *TagsController) Get() {
	this.Data["IsTag"] = true
	this.TplName = "default/tags.html"
}
