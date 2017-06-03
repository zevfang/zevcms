package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"strings"
)

var langTypes []*langType

type langType struct {
	Lang, Name string
}

type BaseController struct {
	beego.Controller
	i18n.Locale
}

func (this *BaseController) Prepare() {
	if this.setLangVer() {
		i := strings.Index(this.Ctx.Request.RequestURI, "?")
		this.Redirect(this.Ctx.Request.RequestURI[:i], 302)
		return
	}
}

func (this *BaseController) setLangVer() bool {
	isNeedRedir := false
	hasCookie := false

	lang := this.Input().Get("lang")
	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	if len(lang) == 0 {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5]
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	if len(lang) == 0 {
		lang = "zh-CN"
		isNeedRedir = false
	}

	culLang := langType{
		Lang: lang,
	}

	if !hasCookie {
		this.Ctx.SetCookie("lang", culLang.Lang, 1<<31-1, "/")
	}

	restLangs := make([]*langType, 0, len(langTypes)-1)
	for _, v := range langTypes {
		if lang != v.Lang {
			restLangs = append(restLangs, v)
		} else {
			culLang.Name = v.Name
		}
	}

	this.Lang = lang
	this.Data["Lang"] = culLang.Lang
	this.Data["CurLang"] = culLang.Name
	this.Data["RestLangs"] = restLangs

	return isNeedRedir
}
