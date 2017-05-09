package controllers

type CategoryController struct {
	BaseController
}

func (this *CategoryController) Get() {
	this.Data["IsCategory"] = true
	this.TplName = "category.html"
}
