package main

import (
	"encoding/json"
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

type IntMap struct {
	IntMaps   map[int]string    `json:"int_maps,select(IntMaps)"`
	UintMap   map[uint]string   `json:"uint_map,select(UintMap)"`
	StringMap map[string]string `json:"string_map,select(StringMap)"`
	//BoolMap    map[bool]string      `json:"bool_map,select(BoolMap)"`
	//FloatMap   map[float32]string   `json:"float_map,select(FloatMap)"`
	//ComplexMap map[complex64]string `json:"complex_map,select(ComplexMap)"`
}

func TestMap() {
	//a := 1
	//m := map[interface{}]string{
	//	a:   "ooo",
	//	"b": "bbb",
	//}
	//marshal, err := json.Marshal(m)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(marshal))

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

	mmm := IntMap{
		IntMaps:   map[int]string{-1: "一", 1: "二"},
		UintMap:   map[uint]string{1: "一", 2: "二"},
		StringMap: map[string]string{"s": "s"},
		//FloatMap: map[float32]string{1.12: "一.12", 2.67657: "二.67657"},
		//BoolMap:    map[bool]string{true: "true", false: "false"},
		//ComplexMap: map[complex64]string{complex64(1): "1", complex64(2): "2"},
	}
	marshal, err := json.Marshal(mmm)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))

	fmt.Println(filter.Select("IntMaps", mmm))
	fmt.Println(filter.Select("UintMap", mmm))
	fmt.Println(filter.Select("StringMap", mmm))

}
