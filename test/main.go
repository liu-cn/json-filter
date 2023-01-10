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

type UID [3]byte
type UIDs []byte

//func (u UID) String() string {
//	return "uid"
//}

func (u UID) MarshalText() (text []byte, err error) {
	return []byte("uid"), nil
}

type Us struct {
	//Name       string   `json:"name,select(all),omit(h)"`
	B          []byte   `json:"b,select(all),omit(h)"`
	EmptySlice []string `json:"empty_slice,select(all),omit(h)"`

	//H    struct{} `json:"h,select(h)"`
	//Uu   Uu       `json:"uu,select(h),omit(h)"`
	//S    []string `json:"s,select(h)"`

	BB      [3]byte `json:"bb,select(all)"`
	Avatar  []byte  `json:"avatar,select(all),func(GetAvatar)"`
	Avatar2 []byte  `json:"avatar2,select(all),func(GetAvatar2)"`
	UID     UID     `json:"uid,select(all)"`
	UIDs    UIDs    `json:"uids,select(all)"`
}

func (u Us) GetAvatar() string {
	return string(u.Avatar[:]) + ".jpg"
}
func (u *Us) GetAvatar2() string {
	return string(u.Avatar[:]) + ".jpg"
}

func newUs() Us {
	return Us{
		//Name: "1",
		//H:    struct{}{},
		//Uu: Uu{
		//	Name: "uu",
		//},
		//S: []string{"1", "2"},
	}
}

func main() {

	//var bb = []byte(`{"a":"1"}`)
	u := Us{
		BB:         [3]byte{1, 2, 4},
		EmptySlice: make([]string, 0, 1),
		B:          []byte(`{"a":"1"}`),
		UID:        UID{1, 3, 4},
		UIDs:       UIDs{1, 23, 55},
		Avatar:     []byte("uuid"),
		Avatar2:    []byte("uuid2"),
	}

	//fmt.Println(mustJson(u))
	//fmt.Println(filter.Omit("h", u))
	fmt.Println(filter.Select("all", u))
	fmt.Println(filter.Omit("all", u))
	fmt.Println(filter.Select("all", &u))
	fmt.Println(filter.Omit("all", &u))
	TestSlice()
	TestMap()
	TestU()
}
