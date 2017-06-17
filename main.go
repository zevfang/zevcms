package main

import (
	"github.com/astaxie/beego"
	"zevcms/controllers"
	"github.com/beego/i18n"
	"zevcms/models"
	"github.com/astaxie/beego/orm"
)

func main() {
	models.InitModels()
	controllers.InitApp()

	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/tags", &controllers.TagsController{})
	beego.Router("/archives", &controllers.ArchivesController{})

	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/", &controllers.AdminController{}),
		beego.NSRouter("/main", &controllers.AdminController{}),
		beego.NSRouter("/login", &controllers.LoginController{}),
		beego.NSNamespace("/category",
			beego.NSRouter("/list", &controllers.CategoryController{}, "get:CategoryList"),
			beego.NSRouter("/add", &controllers.CategoryController{}, "post:AddCategory"),
			beego.NSRouter("/edit", &controllers.CategoryController{}, "get,post:UpdCategory"),
		),
		beego.NSNamespace("/topic",
			beego.NSRouter("/list", &controllers.TopicController{}, "get:TopicList"),
		),
	)

	beego.AddNamespace(ns)

	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
