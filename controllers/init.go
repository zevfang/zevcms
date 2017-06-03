package controllers

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func InitApp()  {
	initLocales()
}
//多语言配置
func initLocales() {
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")
	langTypes = make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		beego.Trace("加载语言" + lang)
		if err := i18n.SetMessage(lang, "conf/language/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("语言文件加载失败:" + err.Error())
			return
		}

	}
}

