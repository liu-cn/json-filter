package main

//假如在文章(article)场景下，你只需要返回age，name，avatar，uid这四个字段就够了，
//不想暴露出其他不需要的字段，也不想重新建一个struct一个一个字段的赋值，
//返回更多字段在不安全的同时意味着需要传输更多的数据，这就意味着会浪费带宽资源，编解码也更加耗时，
//有时候我看到过有人为了偷懒把一个毫无意义的结构体序列化后返回，上面带着很多无用的字段，可能只有4-5个字段有用，
//其他大多数字段都没有用，不仅影响阅读还浪费带宽，所以或许可以尝试用json-filter的过滤器来过滤你想要的字段吧，
//不仅简单，更重要的是很强大，很复杂的结构体也可以过滤出你想要的字段。

import (
	"fmt"

	"github.com/liu-cn/json-filter/filter"
)

type User struct {
	UID      uint     `json:"uid,select(article|profile)"`  //这个在select()里添加了article场景，
	Name     string   `json:"name,select(article|profile)"` //会被article场景解析
	Age      int      `json:"age,select(article|profile)"`  //会被article场景解析
	Sex      int      `json:"sex,select(profile)"`          //不会被article场景解析，在article场景过滤时会直接忽略该字段。
	Avatar   string   `json:"avatar,select(article)"`       //会被article场景解析
	Password string   `json:"password"`
	Slat     string   `json:"-"` //任何场景都会被忽略
	Lang     []string `json:"lang,select(lang)"`
}

func main() {
	user := User{
		Name:     "boyan",
		UID:      1,
		Age:      20,
		Sex:      1,
		Avatar:   "https://www.avatar.com",
		Password: "pwd",
		Slat:     "slat",
		Lang:     []string{"c", "c++", "Go"},
	}

	articleJson := filter.SelectMarshal("article", user)
	fmt.Println(articleJson) //输出以下json：
	//{"age":20,"avatar":"https://www.avatar.com","name":"boyan","uid":1}

	//需求来了有一个需要展示编程语言的场景，要求只展示你的编程语言的字段，其他字段都不要展示，那怎么办呢？
	//很简单
	langJson := filter.SelectMarshal("lang", &user) //user传递指针和值都无所谓。
	fmt.Println(langJson)                           //输出以下json：
	//{"lang":["c","c++","Go"]}
}
