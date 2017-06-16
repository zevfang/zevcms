package models

import "time"

type Tags struct {
	Id         int64 `orm:"index"`
	TagName    string
	TopicCount int64
	Created    time.Time `orm:"index"`
}
