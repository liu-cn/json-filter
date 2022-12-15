package main

import (
	"encoding/json"
	"fmt"
	"github.com/liu-cn/json-filter/filter"
)

type Slice struct {
	Slices   []**string   `json:"slices,select(test)"`
	Test     []**string   `json:"test,select(),omit(test)"`
	SliceP   *[]**string  `json:"slice_p,select(test)"`
	SlicesPP **[]**string `json:"slices_pp,select(test)"`
}

func TestSlice() {
	s := "值"
	p := &s

	slice := make([]**string, 0, 5)
	slice = append(slice, &p)
	pp := &slice
	ppp := &pp

	test := Slice{
		Slices:   slice,
		SliceP:   pp,
		SlicesPP: ppp,
		Test:     slice,
	}

	fmt.Println("slice select:", filter.SelectMarshal("test", test).MustJSON())
	//{"slice_p":["值"],"slices":["值"],"slices_pp":["值"]}
	fmt.Println("slice omit:", filter.OmitMarshal("test", test).MustJSON())
	//{"slice_p":["值"],"slices":["值"],"slices_pp":["值"]}

	marshal, _ := json.Marshal(test)
	fmt.Println("原生slice json 解析", string(marshal))
	//{"slices":["值"],"test":["值"],"slice_p":["值"],"slices_pp":["值"]}
}
