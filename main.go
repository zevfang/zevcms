package main

import (
	"github.com/astaxie/beego"
	"zevcms/controllers"
	"github.com/beego/i18n"
)

func main() {
	controllers.InitApp()

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/tags", &controllers.TagsController{})
	beego.Router("/archives", &controllers.ArchivesController{})

	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/", &controllers.AdminController{}),
		beego.NSRouter("/main", &controllers.AdminController{}),
		beego.NSRouter("/login", &controllers.LoginController{}),
		beego.NSRouter("/category", &controllers.CategoryController{}, "get:CategoryList"),
		beego.NSNamespace("/topic",
			beego.NSRouter("/list", &controllers.TopicController{}, "get:TopicList"),
		),
	)

	beego.AddNamespace(ns)

	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
