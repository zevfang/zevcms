package controllers

import (
	"zevcms/models"
	"github.com/astaxie/beego"
)

type TagsController struct {
	BaseController
}

func (this *TagsController) Get() {
	this.Data["IsTag"] = true
	this.TplName = "default/tags.html"
}

func (this *TagsController) TagsList() {
	this.TplName = "admin/tags/list.html"
	tags, err := models.GetTagsList()
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["TagsList"] = tags
}

func (this *TagsController) AddTags() {

	if this.Ctx.Request.Method == "POST" {

		tagname := this.Input().Get("TagName")
		if len(tagname) == 0 {
			this.Redirect("/admin/tags/list", 302)
			return
		}
		tagcolor := this.Input().Get("TagColor")
		if len(tagcolor) == 0 {
			this.Redirect("/admin/tags/list", 302)
			return
		}

		err := models.AddTags(tagname, tagcolor)
		if err != nil {
			beego.Error(err)
			return
		}

	}
	this.Redirect("/admin/tags/list", 302)
}

func (this *TagsController) UpdTags() {
	if this.Ctx.Request.Method == "GET" {
		tagId := this.Input().Get("id")
		if len(tagId) == 0 {
			beego.Error("非法访问")
			this.Redirect("/admin/tags/list", 302)
			return
		}
		tag, err := models.GetTagsSingle(tagId)
		if err != nil {
			beego.Error("没有数据")
			this.Redirect("/admin/tags/list", 302)
			return
		}
		this.Data["Tags"] = tag
	}

	if this.Ctx.Request.Method == "POST" {
		tagId := this.Input().Get("Id")
		tagName := this.Input().Get("TagName")
		tagColor := this.Input().Get("TagColor")
		err := models.UpdTags(tagId, tagName, tagColor)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/admin/tags/list", 302)

	}
	this.TplName = "admin/tags/edit.html"
}
func (this *TagsController) DelTags() {
	tagId := this.Input().Get("id")
	if len(tagId) == 0 {
		beego.Error("非法访问")
		this.Ctx.WriteString("{\"state\":\"n\",\"msg\":\"非法访问.\"}")
		return
	}
	err := models.DelTags(tagId)
	if err != nil {
		beego.Error("服务器错误")
		this.Ctx.WriteString("{\"state\":\"n\",\"msg\":\"服务器错误.\"}")
		return
	}
	this.Ctx.WriteString("{\"state\":\"y\",\"msg\":\"删除成功.\"}")
}
