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

	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
