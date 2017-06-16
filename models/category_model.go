package models

import (
	"time"
)

type Category struct {
	Id         int64 `orm:"index""`
	CateName   string
	TopicCount int64
	Created    time.Time `orm:"index""`
}

func GetCategoryList() {

}
