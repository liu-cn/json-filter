package main

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
)

func TestEqJson(t *testing.T) {
	json1 := `{"author":"北洛","pageNum":1,"age":22}`
	json2 := `{"pageNum":1,"age":22,"author":"北洛"}`
	json3 := `{"age":22,"pageNum":1,"author":"北洛"}`
	json4 := `{"age":22,"pageNum":1,"author":"洛北"}`
	json5 := `{"age":22,"pageNum":1}`

	fmt.Println(filter.EqualJSON(json1, json2))        //结构相等且值相等，返回true
	fmt.Println(filter.EqualJSON(json1, json2, json3)) //判断多个，需要全部等价才返回true
	fmt.Println(filter.EqualJSON(json3, json4))        //结构相等但是值不相等，不等价
	fmt.Println(filter.EqualJSON(json3, json5))        //结构不相等不等价
}
