package main

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"

	"testing"
)

func TestSlice(t *testing.T) {
	type Tag struct {
		ID   uint   `json:"id,select(all),omit(id)"`
		Name string `json:"name,select(justName|all),omit(name)"`
		Icon string `json:"icon,select(chat|profile|all),omit(icon)"`
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

	fmt.Println(filter.Omit("id", tags))
	//[{"icon":"icon-c","name":"c"},{"icon":"icon-c++","name":"c++"},{"icon":"icon-go","name":"go"}]
	fmt.Println(filter.Omit("name", tags))
	//[{"icon":"icon-c","id":1},{"icon":"icon-c++","id":1},{"icon":"icon-go","id":1}]
	fmt.Println(filter.Omit("icon", tags))
	//[{"id":1,"name":"c"},{"id":1,"name":"c++"},{"id":1,"name":"go"}]

}
