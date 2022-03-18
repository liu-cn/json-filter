package main

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
)

func TestAnonymous(t *testing.T) {

	type Page struct {
		PageInfo int `json:"pageInfo,select($any)"`
		PageNum  int `json:"pageNum,select($any)"`
	}

	type Article struct {
		Title string                    `json:"title,select(article)"`
		Page  `json:",select(article)"` // 这种tag字段名为空的方式会直接把该结构体展开，当作匿名结构体处理
		//Page `json:"page,select(article)"` // 注意这里tag里标注了匿名结构体的字段名，所以解析时会解析成对象，不会展开
		Author string `json:"author,select(admin)"`
	}

	article := Article{
		Title: "c++从研发到脱发",
		Page: Page{
			PageInfo: 999,
			PageNum:  1,
		},
	}

	articleJson := filter.SelectMarshal("article", article)
	fmt.Println(articleJson.MustJSON())
	//输出结果--->  {"pageInfo":999,"pageNum":1,"title":"c++从研发到脱发"}

}
