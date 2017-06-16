package models

import "time"

type Topic struct {
	Id        int64 `orm:index`
	Title     string
	Content   string `orm:"size(5000)"`
	Created   time.Time `orm:"index"`
	Updated   time.Time `orm:"index"`
	ViewCount int64
}
