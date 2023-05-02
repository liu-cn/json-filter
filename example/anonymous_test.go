package main

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
)

func TestAnonymous(t *testing.T) {

	type AnonymousChild struct {
		PageInfo int `json:"pageInfo,select($any),omit(Anonymous)"`
		PageNum  int `json:"pageNum,select($any)"`
	}

	type Anonymous struct {
		Title          string                    `json:"title,select(article),omit(Anonymous)"`
		AnonymousChild `json:",select(article)"` // 这种tag字段名为空的方式会直接把该结构体展开，当作匿名结构体处理
		//AnonymousChild `json:"page,select(article)"` // 注意这里tag里标注了匿名结构体的字段名，所以解析时会解析成对象，不会展开
		Author string `json:"author,select(admin)"`
	}

	article := Anonymous{
		Author: "北洛",
		Title:  "c++从研发到脱发",
		AnonymousChild: AnonymousChild{
			PageInfo: 999,
			PageNum:  1,
		},
	}

	Json := filter.Select("article", article)
	fmt.Println(Json)
	//输出结果--->  {"pageInfo":999,"pageNum":1,"title":"c++从研发到脱发"}

	Json = filter.Omit("Anonymous", article)
	fmt.Println(Json)
	//输出结果--->  {"author":"北洛","pageNum":1}
}
