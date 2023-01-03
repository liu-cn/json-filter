package main

import (
	"encoding/json"
	"fmt"
	"github.com/liu-cn/json-filter/filter"
)

func mustJson(v interface{}) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(marshal)
}

type Uu struct {
	Name string `json:"name"`
}

type Us struct {
	Name string   `json:"name,select(h)"`
	H    struct{} `json:"h,select(h)"`
	Uu   Uu       `json:"uu,select(h),omit(h)"`
	S    []string `json:"s,select(h)"`
}

func main() {
	//fmt.Println(filter.Select("intAll", All{}))
	fmt.Println(filter.Select("h", Us{}))
	//fmt.Println(filter.Select("h", struct{}{}))
	//fmt.Println(filter.Omit("h", Us{}))
	//fmt.Println(filter.Omit("h", struct{}{}))
}
