package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Category struct {
	Id         int64 `orm:"pk;auto"`
	CateName   string
	CateColor   string
	TopicCount int64 `orm:"default(0)"`
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
}

func GetCategoryList() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qt := o.QueryTable("category")
	_, err := qt.All(&cates)
	return cates, err
}

func GetCategorySingle(id string) (*Category, error) {
	cateid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	cate := new(Category)
	o := orm.NewOrm()
	qt := o.QueryTable("category")
	err = qt.Filter("Id", cateid).One(cate)
	return cate, err

}

func AddCategory(cateName, cateColor string) error {
	cate := &Category{CateName: cateName, CateColor: cateColor }
	o := orm.NewOrm()
	_, err := o.Insert(cate)
	return err
}

func UpdCategory(id, cateName, cateColor string) error {
	cateId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	cate := &Category{Id: cateId}
	o := orm.NewOrm()

	if o.Read(cate) == nil {
		cate.CateName = cateName
		cate.CateColor = cateColor
		_, err = o.Update(cate)
	}
	return err
}

func DelCategory(id string) error {
	cateid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	cate := &Category{Id: cateid}
	o := orm.NewOrm()
	_, err = o.Delete(cate)
	return err
}
