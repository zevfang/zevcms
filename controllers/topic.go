package controllers

type TopicController struct {
	BaseController
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
}
