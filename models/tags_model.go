package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Tags struct {
	Id         int64 `orm:"pk;auto"`
	TagName    string
	TopicCount int64 `orm:"default(0)"`
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
}

func GetTagsList() ([]*Tags, error) {
	o := orm.NewOrm()
	tags := make([]*Tags, 0)
	qt := o.QueryTable("tags")
	_, err := qt.All(&tags)
	return tags, err
}

func AddTags(tagname string) error {
	tag := &Tags{TagName: tagname}
	o := orm.NewOrm()
	_, err := o.Insert(tag)
	return err
}
