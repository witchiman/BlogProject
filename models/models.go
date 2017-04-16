package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"path"
	"os"
	"strconv"
	"fmt"
	"strings"
	"github.com/astaxie/beego"
)

const (
	_MYSQL_DRIVER = "mysql"
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
	Label string
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

type Reply struct{
	Id int64
	TopicId int64
	NickName string
	Content string `orm:size(1000)`
	Ctime time.Time `orm: index`
}

func RegisterDB()  {

	//注册驱动
	orm.RegisterModel(new(Category), new(Topic), new(Reply) )
	// 注册驱动,属于默认注册，此处代码可省略）
	orm.RegisterDriver(_MYSQL_DRIVER, orm.DRMySQL)
	//注册默认数据库
	// 最后两个参数分别为最大空闲连接和最大数据库连接
	orm.RegisterDataBase("default", _MYSQL_DRIVER, "test:1234@tcp(localhost:3306)/testDB?charset=utf8", 10,10)
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



/*获取所有分类*/
func GetCategories() ([]*Category, error) {
	o := orm.NewOrm()

	categories := make([]*Category, 0)
	result := o.QueryTable("category")

	_, err := result.All(&categories)
	return categories, err
}


/*添加文章*/
func AddTopic(title, content, category, label, attachment string)  error {
	//标签处理，便于查找，如‘network game’处理后变成‘$network#$game#’
	realLabel := "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	o := orm.NewOrm()

	topic := &Topic{
		Title: title,
		Category: category,
		Label: realLabel,
		Attachment: attachment,
		Content: content,
		Ctime: time.Now(),
		Updated: time.Now(),
		ReplyTime: time.Now(),
	}

	_, err := o.Insert(topic)
	if err!=nil{
		return err
	}

	cate := new(Category) //更新文章分类
	result := o.QueryTable("category")
	err = result.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_,err = o.Update(cate)
	}

	return err
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

func ModifyTopic(id,title, content, category, label, attachment string) error  {
	realId, err := strconv.ParseInt(id, 10, 64);
	if err != nil {
		return err
	}

	realLabel := "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	o := orm.NewOrm()

	 var oldcate, oldattachment string
	//使用Read函数时，传入的的参数必须含有主键值，如Topic的Id。查找成功，没有错误返回则进行下一步处理
	topic := &Topic{Id: realId}
	if o.Read(topic) == nil {
		oldcate = topic.Category    //获取原来的分类
		oldattachment = topic.Attachment //获取原来的文件路径
		topic.Title = title
		topic.Label = realLabel
		topic.Content = content
		topic.Category = category
		topic.Attachment = attachment
		topic.Updated = time.Now()
		_, err = o.Update(topic)
		if err != nil {
			return err
		}
	}

	if !strings.EqualFold(category, oldcate) {   //分类有发生变动时才更新分类
		if len(oldcate) > 0 {
			result := o.QueryTable("category")   //更新旧的文章分类
			cate := new(Category)
			err = result.Filter("title", oldcate).One(cate)
			if err == nil {
				cate.TopicCount--
				_, err = o.Update(cate)
			}
		}

		if len(category)>0{
			result := o.QueryTable("category")   //更新新的文章分类
			cate := new(Category)
			err = result.Filter("title", category).One(cate)
			if err == nil {
				cate.TopicCount++
				_, err = o.Update(cate)
			}
		}
	}

	beego.Info("ModifyTopic: 现在的附件：", attachment, "原附件： ",oldattachment)
	if !strings.EqualFold(attachment, oldattachment) {  //如果前后存储文件不一致，需要把原来的文件删除
		if len(oldattachment) > 0 {
			err = os.Remove(path.Join("attachment", oldattachment))
		}
	}

	return err
}

func GetAllTopics(category, label string, desc bool) ([]*Topic, error)  {
	topics :=make([]*Topic, 0)
	o := orm.NewOrm()

	result := o.QueryTable("topic")
	var err error
	if desc {  //按最近修改时间排序,OrderBy参数前加‘-’，可使之排倒序排列
		if len(category) > 0{   //按分类筛选
			result = result.Filter("category", category)
		}

		if len(label)>0 {  //按标签进行精确查找
			result = result.Filter("label__contains","$"+label+"#")
		}

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

	//把处理过的标签语句还原
	topic.Label = strings.Replace(strings.Replace(topic.Label, "#", " ", -1), "$", "",-1)

	return topic, err
}

func AddReply(topicId, nickName, content string) error {
	tid, err := strconv.ParseInt(topicId, 10, 64)
	if err != nil {
		return err
	}

	reply := &Reply{
		TopicId: tid,
		NickName: nickName,
		Content: content,
		Ctime: time.Now(),
	}

	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err != nil {
		return err
	}

	topic := &Topic{Id: tid} //更新文章评论数和最新回复时间
	if o.Read(topic) == nil {
		topic.ReplyTime = time.Now()
		topic.ReplyCount++
		_, err = o.Update(topic)
	}

	return err
}

func DeleteReply(replyid string) error {
	rid, err := strconv.ParseInt(replyid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	var topicId int64
	reply := &Reply{Id: rid}
	if o.Read(reply) == nil {
		topicId = reply.TopicId
		_, err = o.Delete(reply)
		if err != nil {
			return err
		}
	}

	//查看剩下的回复，按创建时间倒序查看，取最近的评论
	replies := make([]*Reply, 0)
	result := o.QueryTable("reply")
	_, err = result.Filter("topicid", topicId).OrderBy("-ctime").All(&replies)
	if err != nil {
		return err
	}


	topic := &Topic{Id: topicId} //更新文章评论数和最新回复时间
	if o.Read(topic) == nil {
		topic.ReplyTime = replies[0].Ctime
		topic.ReplyCount = int64(len(replies))
		_,err = o.Update(topic)
	}


	return err
}

func GetAllReplies(topicId string) ([]*Reply, error) {
	tid, err := strconv.ParseInt(topicId, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	replies := make([]*Reply, 0)

	result := o.QueryTable("reply")
	_, err = result.Filter("topicid", tid).All(&replies)

	return replies, err
}