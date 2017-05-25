package controllers

type ArchivesController struct {
	BaseController
}

func (this *ArchivesController) Get() {
	this.Data["IsTag"] = true
	this.TplName = "default/archives.html"
}
