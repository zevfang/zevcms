package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Tags struct {
	Id         int64 `orm:"pk;auto"`
	TagName    string
	TagColor   string
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

func AddTags(tagname, tagcolor string) error {
	tag := &Tags{TagName: tagname, TagColor: tagcolor}
	o := orm.NewOrm()
	_, err := o.Insert(tag)
	return err
}

func GetTagsSingle(tagId string) (*Tags, error) {
	tagid, err := strconv.ParseInt(tagId, 10, 64)
	if err != nil {
		return nil, err
	}
	tag := new(Tags)
	o := orm.NewOrm()
	qt := o.QueryTable("tags")
	err = qt.Filter("Id", tagid).One(tag)
	return tag, err
}

func UpdTags(tagId, tagName, tagColor string) error {
	tagid, err := strconv.ParseInt(tagId, 10, 64)
	if err != nil {
		return err
	}
	tag := &Tags{Id: tagid, }
	o := orm.NewOrm()
	if err := o.Read(tag); err == nil {
		tag.TagName = tagName
		tag.TagColor = tagColor
		_, err = o.Update(tag)
	}
	return err
}

func DelTags(tagId string) error {
	tagid, err := strconv.ParseInt(tagId, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	_, err = o.Delete(&Tags{Id: tagid})
	return err

}
