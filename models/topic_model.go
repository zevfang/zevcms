package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Topic struct {
	Id        int64 `orm:"pk;auto"`
	CateId    int64
	TagsId    string
	Title     string
	Content   string `orm:"size(5000)"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time
	ViewCount int64 `orm:"default(0)"`
}

func AddTopic(cateid, tagid, title, content string) error {
	cateId, err := strconv.ParseInt(cateid, 10, 64)
	if err != nil {
		return err
	}

	topic := &Topic{CateId: cateId, TagsId: tagid, Title: title, Content: content}
	cate := &Category{Id: cateId}

	o := orm.NewOrm()

	err = o.Begin()

	//添加文章
	_, err = o.Insert(topic)

	//更新类型数量
	if err := o.Read(cate); err == nil {
		cate.Id = cateId
		cate.TopicCount ++
		_, err = o.Update(cate)
	}

	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return err
}
