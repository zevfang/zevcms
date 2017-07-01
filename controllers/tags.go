package controllers

type TagsController struct {
	BaseController
}

func (this *TagsController) Get() {
	this.Data["IsTag"] = true
	this.TplName = "default/tags.html"
}
func (this *TagsController) TagsList()  {
	this.TplName = "admin/tags/list.html"
}
func (this *TagsController) GetTagsList()  {
	this.TplName = "admin/tags/list.html"
}
func (this *TagsController) AddTags()  {
	this.TplName = "admin/tags/list.html"
}
func (this *TagsController) UpdTags()  {
	this.TplName = "admin/tag/edit.html"
}
func (this *TagsController) DelTags()  {
	
}