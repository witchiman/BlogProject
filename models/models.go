package models

import (
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"path"
	"os"
)

const (
	_DB_NAME = "/data/blog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct{
	Id int64
	Title string
	Ctime time.Duration     `orm:"index"`
	Views int64		`orm:"index`
	TopicTime time.Duration `orm:"index`
	TopicCount int64
	TopicLastUser int64
}

type Topic struct {
	Id int64
	Uid int64
	Title string
	Content string 		`orm:"size(5000)"`
	Attachment string
	Ctime time.Duration 	`orm:"index"`
	Updated time.Duration   `orm:"index"`
	Views int64  		`orm:"index"`
	Authour string
	ReplyTime time.Duration	 `orm:"index"`
	ReplyCount int64
	ReplyLastId int64
}

func RegisterDB()  {
	if !com.IsDir(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic) )
	orm.RegisterDriver(_DB_NAME, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, "10")
}
