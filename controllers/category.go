package controllers

import (
	"zevcms/models"
	"github.com/astaxie/beego"
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

	if this.Ctx.Request.Method == "GET" {
		cateId := this.Input().Get("id")
		if len(cateId) == 0 {
			beego.Error("非法访问")
			this.Redirect("/admin/category/list", 302)
			return
		}
		cate, err := models.GetCategorySingle(cateId)
		if err != nil {
			beego.Error("没有数据")
			this.Redirect("/admin/category/list", 302)
			return
		}
		this.Data["Category"] = cate
	}

	if this.Ctx.Request.Method == "POST" {
		cateId := this.Input().Get("Id")
		cateName := this.Input().Get("CateName")
		err := models.UpdCategory(cateId, cateName)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/admin/category/list", 302)

	}
	this.TplName = "admin/category/edit.html"

}

func (this *CategoryController) DelCategory() {
	cateId := this.Input().Get("id")
	if len(cateId) == 0 {
		beego.Error("非法访问")
		this.Ctx.WriteString("{\"state\":\"n\",\"msg\":\"非法访问.\"}")
	}
	err := models.DelCategory(cateId)
	if err != nil {
		beego.Error(err)
		this.Ctx.WriteString("{\"state\":\"n\",\"msg\":\"服务器错误.\"}")
	}
	this.Ctx.WriteString("{\"state\":\"y\",\"msg\":\"删除成功.\"}")
}
