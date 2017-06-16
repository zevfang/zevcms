package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

func InitModels() {
	//orm.DefaultTimeLoc = time.Local
	orm.RegisterModel(new(Category), new(Topic), new(Tags))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:fangwei0505@/zevcms?charset=utf8&parseTime=true&loc=Local")
}
