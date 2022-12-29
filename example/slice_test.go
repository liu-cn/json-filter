package main

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
)

func TestSlice(t *testing.T) {
	type Tag struct {
		ID   uint   `json:"id,select(all)"`
		Name string `json:"name,select(justName|all)"`
		Icon string `json:"icon,select(chat|profile|all)"`
	}

	tags := []Tag{ //切片和数组都支持 slice or array
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

	fmt.Println(filter.Select("justName", tags))
	//--->输出结果： [{"name":"c"},{"name":"c++"},{"name":"go"}]

	fmt.Println(filter.Select("all", tags))
	//--->输出结果： [{"icon":"icon-c","id":1,"name":"c"},{"icon":"icon-c++","id":1,"name":"c++"},{"icon":"icon-go","id":1,"name":"go"}]

	fmt.Println(filter.Select("chat", tags))
	//--->输出结果： [{"icon":"icon-c"},{"icon":"icon-c++"},{"icon":"icon-go"}]
}
