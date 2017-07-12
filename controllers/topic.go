package controllers

import (
	"zevcms/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	BaseController
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.TplName = "default/topic.html"
}




func (this *TopicController) AddTopic() {
	this.TplName="admin/topic/add.html"
}

func (this *TopicController) TopicList() {
	this.TplName = "admin/topic/list.html"
	var (
		err   error
		tags  []*models.Tags
		cates []*models.Category
	)
	tags, err = models.GetTagsList()
	if err != nil {
		beego.Error(err)
	}
	this.Data["TagsList"] = tags;
	cates, err = models.GetCategoryList()
	if err != nil {
		beego.Error(err)
	}
	this.Data["CategoryList"] = cates;
}
