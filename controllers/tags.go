package controllers

import (
	"zevcms/models"
	"github.com/astaxie/beego"
	"fmt"
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
		fmt.Println("123")
		tagname := this.Input().Get("TagName")
		if len(tagname) == 0 {
			this.Redirect("/admin/tags/list", 302)
			return
		}
		err:=models.AddTags(tagname)
		if err != nil {
			beego.Error(err)
			return
		}

	}
	this.Redirect("/admin/tags/list", 302)
}

func (this *TagsController) UpdTags() {
	this.TplName = "admin/tag/edit.html"
}
func (this *TagsController) DelTags() {

}
