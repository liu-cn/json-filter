package main

import (
	"encoding/json"
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
)

func TestMapWithValue(t *testing.T) {
	type F struct {
		A string `json:"a,select(a)"`
	}
	//filterMap:=filter.Select("a",&F{A: "a"}).(filter.Filter).Map()
	filterMap := filter.SelectFilter("a", &F{A: "a"}).Map() // 推荐使用 typed API
	filterMap["b"] = "b"
	filterMap["cc"] = struct {
		CC string
	}{
		CC: "CC",
	}

	marshal, err := json.Marshal(filterMap)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshal))
	//{
	//	"a": "a",
	//	"b": "b",
	//	"cc": {"CC": "CC"}
	//}

}

func TestSliceWithValue(t *testing.T) {
	type F struct {
		A string `json:"a,select(a)"`
	}

	list := []F{
		F{A: "a"},
	}

	//slices:=filter.Select("a",&F{A: "a"}).(filter.Filter).Slice()
	slices := filter.SelectFilter("a", list).Slice() // 推荐使用 typed API

	slices = append(slices, F{A: "b"})
	slices = append(slices, F{A: "c"})

	marshal, err := json.Marshal(slices)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshal))
	//[
	//	{
	//	"a": "a"
	//	},
	//	{
	//	"a": "b"
	//	},
	//	{
	//	"a": "c"
	//	}
	//]

}
