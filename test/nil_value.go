package main

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"
)

func TestNilValue() {
	var a *T
	var c *T
	var el = map[string]interface{}{
		"a": a, // a: nil ptr
		"b": 1,
		"c": map[string]interface{}{
			"bb": nil,
			"dd": c, // dd: nil ptr
		},
	}

	fmt.Println(filter.Select("test", el))
	//{"a":null,"b":1,"c":{"bb":null,"dd":null}}

}

type T struct {
	A string `json:"a,select(test)"`
}
