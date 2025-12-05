package main

import (
	"fmt"
	"github.com/glebarez/sqlite" // ← 纯 Go sqlite 驱动
	"gorm.io/gorm"
)

/**
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。

基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。

为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

type User struct {
	gorm.Model
	Name      string
	Email     string
	PostCount int
	Posts     []Post
	Comments  []Comment
}

type Post struct {
	gorm.Model
	Title    string
	Content  string
	Comments []Comment
	UserID   uint
	User     User
	Status   string
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户文章数量
	var user User
	tx.Model(&User{}).Where("id", p.UserID).First(&user)
	user.PostCount++
	tx.Model(&User{}).Where("id", p.UserID).Update("post_count", user.PostCount)
	return
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
	Post    Post
	UserID  uint
	User    User
}

// 钩子函数
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	// 更新文章的评论数量
	var post Post
	tx.Model(&Post{}).Where("id", c.PostID).First(&post)
	post.Status = "有评论"
	tx.Model(&Post{}).Where("id", c.PostID).Update("status", post.Status)
	return
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 检查文章的评论数量
	var count int64
	tx.Model(&Comment{}).Where("post_id", c.PostID).Count(&count)
	if count == 0 {
		tx.Model(&Post{}).Where("id", c.PostID).Update("status", "无评论")
	}
	return
}

var DB *gorm.DB

func initDB() *gorm.DB {
	// 连接 SQLite（文件不存在会自动创建）
	conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("链接数据库")
	return conn
}

func InitTable(db *gorm.DB) {
	fmt.Println("初始化数据库")
	defer fmt.Println("初始化数据库完毕")
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	db.Debug().Unscoped().Delete(&User{}, "1=1")
	db.Debug().Unscoped().Delete(&Post{}, "1=1")
	db.Debug().Unscoped().Delete(&Comment{}, "1=1")
}

func InitData(db *gorm.DB) {
	users := []User{
		{Name: "张三", Email: "zhangsan@example.com"},
		{Name: "李四", Email: "lisi@example.com"},
		{Name: "王五", Email: "wangwu@example.com"},
		{Name: "赵六", Email: "zhaoliu@example.com"},
	}
	db.Create(&users)

	posts := []Post{}

	for _, u := range users {
		posts = append(posts, Post{Title: "第一篇博客", Content: "第一篇博客的内容", UserID: u.ID})
		posts = append(posts, Post{Title: "第二篇博客", Content: "第二篇博客的内容", UserID: u.ID})
		posts = append(posts, Post{Title: "第三篇博客", Content: "第三篇博客的内容", UserID: u.ID})
	}
	db.Create(&posts)

	comments := []Comment{}
	for _, p := range posts {
		for _, u := range users {
			comments = append(comments, Comment{Content: "评论啦啦啦啦啦", PostID: p.ID, UserID: u.ID})
		}
	}

	db.Create(&comments)

	fmt.Println("初始化数据")
}

func main() {
	db := initDB()
	InitTable(db)
	InitData(db)

	//var users []User
	//db.Preload("Posts").Preload("Posts.Comments").Preload("Posts.Comments.User").Find(&users)
	//
	//fmt.Println("用户数量：", len(users))
	//fmt.Println("用户列表:")
	//fmt.Println("--------------")
	//for _, u := range users {
	//	fmt.Println("用户:", u.Name)
	//	for _, p := range u.Posts {
	//		fmt.Println("  文章:", p.Title)
	//		for _, c := range p.Comments {
	//			fmt.Println("    评论:", c.Content, "作者:", c.User.Name)
	//		}
	//	}
	//}

	// 使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	fmt.Println("使用Gorm查询某个用户发布的所有文章及其对应的评论信息。")
	var user User
	db.Debug().Model(&User{}).Where("name", "张三").Preload("Posts").Preload("Comments").First(&user)
	fmt.Println("用户名", user.Name, "累计文章数", len(user.Posts), "累计评论数", len(user.Comments), "用户文章数量", user.PostCount)
	fmt.Println("累计文章数", len(user.Posts))
	fmt.Println("累计评论数", len(user.Comments))

	// 使用Gorm查询评论数量最多的文章信息。
	fmt.Println("使用Gorm查询评论数量最多的文章信息。")
	var posts map[string]interface{}
	db.Debug().Model(&Post{}).Joins("left join comments on comments.post_id = posts.id").Select("posts.*, count(comments.id) count").Group("posts.id").Order("count desc").Limit(1).Scan(&posts)
	fmt.Println(posts)

	// 钩子函数
	fmt.Println("钩子函数")
	testPost := Post{Title: "张三的文章", Content: "第一篇博客的内容", UserID: user.ID}

	// 创建文章
	db.Create(&testPost)
	fmt.Println("文章创建成功")
	testComment := []Comment{
		{Content: "张三的评论", PostID: testPost.ID, UserID: user.ID},
		{Content: "张三的评论", PostID: testPost.ID, UserID: user.ID},
		{Content: "张三的评论", PostID: testPost.ID, UserID: user.ID},
	}
	db.Create(&testComment)
	db.Model(&Post{}).Where("id", testPost.ID).First(&testPost)
	fmt.Println("文章状态", testPost.Status)

	fmt.Println("评论创建成功")
	for _, c := range testComment {
		db.Delete(&c)
	}
	fmt.Println("评论删除成功")
	db.Model(&Post{}).Where("id", testPost.ID).First(&testPost)
	fmt.Println("文章状态", testPost.Status)
}
