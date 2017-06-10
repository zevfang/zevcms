package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"time"
)

type Category struct {
	CateId     int64 `orm:"index""`
	CateName   string
	TopicCount int64
	Created    time.Time `orm:"index""`
}

type Topic struct {
	TopicId   int64 `orm:index`
	Title     string
	Content   string `orm:"size(5000)"`
	Created   time.Time `orm:"index"`
	Updated   time.Time `orm:"index"`
	ViewCount int64
}

type Tags struct {
	TagId      int64 `orm:"index"`
	TagName    string
	TopicCount int64
	Created    time.Time `orm:"index"`
}

func InitModels() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("zevcms","mysql","root:root@/zevcms?charset=utf8")
}
