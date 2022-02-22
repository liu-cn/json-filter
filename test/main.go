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

func main() {

	article := Article{
		Title: "c++从研发到脱发",
		Page: Page{
			PageInfo: 999,
			PageNum:  1,
		},
	}

	articleJson := filter.SelectMarshal("article", article)
	fmt.Println(articleJson)
}
