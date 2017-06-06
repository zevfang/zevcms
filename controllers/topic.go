package controllers

type TopicController struct {
	BaseController
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.TplName = "default/topic.html"
}

func (this *TopicController) TopicList() {
	this.TplName = "admin/topic/list.html"
}
