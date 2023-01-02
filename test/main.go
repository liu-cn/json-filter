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

func main() {

	fmt.Println(filter.Select("intAll", All{}))
}
