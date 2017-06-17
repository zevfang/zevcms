package controllers

import (
	"zevcms/models"
	"github.com/astaxie/beego"
	"fmt"
)

type CategoryController struct {
	BaseController
}

func (this *CategoryController) Get() {
	this.Data["IsCategory"] = true
	this.TplName = "category.html"
}

func (this *CategoryController) CategoryList() {
	this.TplName = "admin/category/list.html"
	cates, err := models.GetCategoryList()
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["CateList"] = cates
}

func (this *CategoryController) AddCategory() {
	cateName := this.Input().Get("CateName")
	if len(cateName) == 0 {
		this.Redirect("/admin/category/list", 302)
		return
	}
	err := models.AddCategory(cateName)
	if err != nil {
		beego.Error(err)
		return
	}
	this.Redirect("/admin/category/list", 302)
}

func (this *CategoryController) UpdCategory() {
	fmt.Println(this.Input().Get("id"))
	this.TplName = "admin/category/edit.html"
}
