package main

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"
)

type Map struct {
	//Name string              `json:"name,select(article)"`
	//Age  int                 `json:"age,select(article)"`
	//ID   int                 `json:"id,select(article)"`
	M   map[string]**string   `json:"m,select(test)"`
	T   map[string]**string   `json:"t,select(),omit(test)"`
	MP  *map[string]**string  `json:"mp,select(test)"`
	MPP **map[string]**string `json:"mpp,select(test)"`
}

func TestMap() {

	str := "c++从研发到脱发"
	ptr := &str
	maps := make(map[string]**string)
	maps["test"] = &ptr
	mp := &maps
	mpp := &mp
	fmt.Println("select:", filter.Select("test", Map{M: maps, T: maps, MP: mp, MPP: mpp}))
	//	{"m":{"test":"c++从研发到脱发"},"mp":{"test":"c++从研发到脱发"},"mpp":{"test":"c++从研发到脱发"}}

	fmt.Println("omit:", filter.Select("test", Map{M: maps, T: maps, MP: mp, MPP: mpp}))
	//{"m":{"test":"c++从研发到脱发"},"mp":{"test":"c++从研发到脱发"},"mpp":{"test":"c++从研发到脱发"}}
}
