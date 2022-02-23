package main

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"
)

type User struct {
	UserName string `json:"userName,select(resp)"`
	Book     `json:",select(resp)"`
}

type Book struct {
	BookName string `json:"bookName,select(resp)"`
	*Page    `json:"page,select(resp)"`
}

type Page struct {
	PageInfo int `json:"pageInfo,select($any)"`
	PageNum  int `json:"pageNum,select($any)"`
}

type Article struct {
	Title  string `json:"title,select(article)"`
	Page   `json:"page,select(article)"`
	Author string `json:"author,select(admin)"`
}

type Tag struct {
	ID   uint   `json:"id,select(all)"`
	Name string `json:"name,select(justName|all)"`
	Icon string `json:"icon,select(chat|profile|all)"`
}

func main() {

	//article := Article{
	//	Title: "c++从研发到脱发",
	//	Page: Page{
	//		PageInfo: 999,
	//		PageNum:  1,
	//	},
	//}
	//
	//articleJson := filter.SelectMarshal("article", article)
	//fmt.Println(articleJson)

	type Tag struct {
		ID   uint   `json:"id,select(all)"`
		Name string `json:"name,select(justName|all)"`
		Icon string `json:"icon,select(chat|profile|all)"`
	}

	tags := [3]Tag{
		{
			ID:   1,
			Name: "c",
			Icon: "icon-c",
		},
		{
			ID:   1,
			Name: "c++",
			Icon: "icon-c++",
		},
		{
			ID:   1,
			Name: "go",
			Icon: "icon-go",
		},
	}

	fmt.Println(filter.SelectMarshal("justName", tags))
	//--->输出结果： [{"name":"c"},{"name":"c++"},{"name":"go"}]

	fmt.Println(filter.SelectMarshal("all", tags))
	//--->输出结果： [{"icon":"icon-c","id":1,"name":"c"},{"icon":"icon-c++","id":1,"name":"c++"},{"icon":"icon-go","id":1,"name":"go"}]

	fmt.Println(filter.SelectMarshal("chat", tags))
	//--->输出结果： [{"icon":"icon-c"},{"icon":"icon-c++"},{"icon":"icon-go"}]

}
