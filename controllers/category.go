package controllers

type CategoryController struct {
	BaseController
}

func (this *CategoryController) Get() {
	this.Data["IsCategory"] = true
	this.TplName = "category.html"
}

func (this *CategoryController) CategoryList() {
	this.Data["catelist"] = "nihao"
}
