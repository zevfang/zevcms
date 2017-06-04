package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "admin/login.html"
}

func (this *LoginController) Post() {
	this.Redirect("main", 302)
}
