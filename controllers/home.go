package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["Website"] = "zevfang.com"
	c.Data["Email"] = "fangwei0505@gmail.com"
	c.TplName = "home.html"
}
