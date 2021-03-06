package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"

)

func InitModels() {
	//orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Beijing")
	orm.RegisterModel(new(Category), new(Topic), new(Tags))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:password@/zevcms?charset=utf8&parseTime=true&loc=Local")
}
