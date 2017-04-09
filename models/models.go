package models

import (
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"path"
	"os"
	"strconv"
	"fmt"
)

const (
	_DB_NAME = "data/blog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct{
	Id int64
	Title string
	Ctime time.Time     `orm:"index"`
	Views int64		`orm:"index`
	TopicTime time.Time `orm:"index`
	TopicCount int64
	TopicLastUser int64
}

type Topic struct {
	Id int64
	Uid int64
	Title string
	Category string
	Content string 		`orm:"size(5000)"`
	Attachment string
	Ctime time.Time	`orm:"index"`
	Updated time.Time   `orm:"index"`
	Views int64  		`orm:"index"`
	Authour string
	ReplyTime time.Time
	ReplyCount int64
	ReplyLastId int64
}

func RegisterDB()  {
	//检察数据库文件
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	//注册驱动
	orm.RegisterModel(new(Category), new(Topic) )
	// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
	orm.RegisterDriver(_DB_NAME, orm.DRSqlite)
	//注册默认数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}


/*添加分类*/
func AddCategory(title string) error  {
	fmt.Println("AddCategory: the catetory is :", title)
	o := orm.NewOrm()

	entries := &Category{
		Title: title,
		Ctime: time.Now(),
		TopicTime: time.Now(),
	}

	//查询数据
	result := o.QueryTable("category")
	err := result.Filter("title", title).One(entries)
	if err == nil{  	//err为nil，说明数据库中已经存在相同值，返回
		return err
	}

	//插入数据
	_, err = o.Insert(entries)
	if err != nil {
		return err
	}

	return nil
}

/*删除分类*/
func DeleteCategory(id string) error {
	readId, err := strconv.ParseInt(id, 10, 64);
	if err!= nil{
		return err
	}
	o := orm.NewOrm()
	entry := &Category{Id:readId}

	_, err = o.Delete(entry)
	return err
}

func UpdateCategory(title string) error  {
	o := orm.NewOrm()

	category := &Category{Title: title}
	fmt.Println("UpdateCategory:", category)

	if o.Read(category) == nil{
		category.TopicCount++
		o.Update(category)
	}

	return nil
}

/*获取所有分类*/
func GetCategories() ([]*Category, error) {
	o := orm.NewOrm()

	categories := make([]*Category, 0)
	result := o.QueryTable("category")

	_, err := result.All(&categories)
	return categories, err
}


/*添加文章*/
func AddTopic(title, content, category string)  error {
	o := orm.NewOrm()

	topic := &Topic{
		Title: title,
		Category: category,
		Content: content,
		Ctime: time.Now(),
		Updated: time.Now(),
		ReplyTime: time.Now(),
	}

	_, err := o.Insert(topic)
	if err!=nil{
		return err
	}

	err = UpdateCategory(category)  //更新文章分类
	fmt.Println("AddTopic", title)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTopic(id string)  error {
	realId, err := strconv.ParseInt(id, 10, 64);
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	topic := &Topic{Id: realId}

	_, err = o.Delete(topic)
	return err
}

func ModifyTopic(id,title, content string) error  {
	realId, err := strconv.ParseInt(id, 10, 64);
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	topic := &Topic{Id: realId}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}

	return nil
}

func GetAllTopics(desc bool) ([]*Topic, error)  {
	topics :=make([]*Topic, 0)
	o := orm.NewOrm()

	result := o.QueryTable("topic")

	var err error
	if desc {  //按最近修改时间排序
		_, err = result.OrderBy("-updated").All(&topics)
	}else {
		_, err = result.All(&topics)
	}

	return topics, err
}

func GetTopic(id string) (*Topic, error) {
	realId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topic := new(Topic)
	result := o.QueryTable("topic")
	err = result.Filter("id", realId).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++    //更新文章访问量
	_, err = o.Update(topic)

	return topic, err
}