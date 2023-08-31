package main

import (
	"encoding/json"
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"time"
)

func mustJson(v interface{}) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(marshal)
}

type UID [3]byte
type UIDs []byte

func (u UID) MarshalText() (text []byte, err error) {
	return []byte("uid"), nil
}

type Us struct {
	//Name       string   `json:"name,select(all),omit(h)"`
	B          []byte   `json:"b,select(all),omit(h)"`
	EmptySlice []string `json:"empty_slice,select(all),omit(h)"`
	BB         [3]byte  `json:"bb,select(all)"`
	Avatar     []byte   `json:"avatar,select(all),func(GetAvatar)"`
	Avatar2    []byte   `json:"avatar2,select(all),func(GetAvatar2)"`
	UID        UID      `json:"uid,select(all)"`
	UIDs       UIDs     `json:"uids,select(all)"`
}

func (u Us) GetAvatar() string {
	return string(u.Avatar[:]) + ".jpg"
}
func (u *Us) GetAvatar2() string {
	return string(u.Avatar[:]) + ".jpg"
}

func goOmit() {
	type (
		Tag struct {
			Icon string `json:"icon,omit(article)"`
			Name string `json:"name,omit(profile)"`
		}
		Articles struct {
			Password int   `json:"password,omit($any)"` //$any表示任何场景都排除该字段
			Tags     []Tag `json:"tags"`
			Hot      int   `json:"hot,select(img),func(GetHot)"` //热度 过滤时会调用GetHot方法获取该字段的值
			Like     int   `json:"-"`
			Collect  int   `json:"-"`
		}
	)

	articles := Articles{Like: 100, Collect: 20, Tags: []Tag{{"icon", "foo"}, {"icon", "bar"}}}

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(filter.Omit("article", &articles)) //尽量传指针，不传指针func选择器中的用指针接收的方法无法被调用
		}()
	}

}
func goSelect() {
	type (
		Tag struct {
			Icon string `json:"icon,omit(article)"`
			Name string `json:"name,omit(profile)"`
		}
		Articles struct {
			Password int   `json:"password,omit($any)"` //$any表示任何场景都排除该字段
			Tags     []Tag `json:"tags"`
			Hot      int   `json:"hot,select(img),func(GetHot)"` //热度 过滤时会调用GetHot方法获取该字段的值
			Like     int   `json:"-"`
			Collect  int   `json:"-"`
		}
	)

	articles := Articles{Like: 100, Collect: 20, Tags: []Tag{{"icon", "foo"}, {"icon", "bar"}}}

	for i := 0; i < 10000; i++ {
		go func() {
			fmt.Println(filter.Select("article", &articles)) //尽量传指针，不传指针func选择器中的用指针接收的方法无法被调用
		}()
	}

}

func main() {

	//goOmit()
	goSelect()
	time.Sleep(time.Second * 10)
	//var bb = []byte(`{"a":"1"}`)
	//u := Us{
	//	BB:         [3]byte{1, 2, 4},
	//	EmptySlice: make([]string, 0, 1),
	//	B:          []byte(`{"a":"1"}`),
	//	UID:        UID{1, 3, 4},
	//	UIDs:       UIDs{1, 23, 55},
	//	Avatar:     []byte("uuid"),
	//	Avatar2:    []byte("uuid2"),
	//}
	//list := []Us{u, u, u}
	//fmt.Println(filter.Omit("1", &list))
	//fmt.Println(mustJson(u))

	//TestMap()
	//TestMap()
	//for i := 0; i < 3; i++ {
	//	ExampleOmit()
	//}
}

func ExampleOmit() {
	type (
		Tag struct {
			Icon string `json:"icon,omit(article)"`
			Name string `json:"name,omit(profile)"`
		}
		Articles struct {
			Password int   `json:"password,omit($any)"` //$any表示任何场景都排除该字段
			Tags     []Tag `json:"tags"`
			Hot      int   `json:"hot,select(img),func(GetHot)"` //热度 过滤时会调用GetHot方法获取该字段的值
			Like     int   `json:"-"`
			Collect  int   `json:"-"`
		}
	)

	//func (a Articles) GetHot() int { //这个方法里可以对字段进行处理，处理后可以返回一个任意值
	//	return a.Like + a.Collect
	//}

	articles := Articles{Like: 100, Collect: 20, Tags: []Tag{{"icon", "foo"}, {"icon", "bar"}}}
	article := filter.Omit("article", &articles) //尽量传指针，不传指针func选择器中的用指针接收的方法无法被调用
	profile := filter.Omit("profile", &articles)
	articleJSON, _ := json.Marshal(article)
	fmt.Println(string(articleJSON))
	fmt.Println(profile) //可以直接打印，打印会直接输出过滤后的json

	//Output:
	//{"hot":120,"tags":[{"name":"foo"},{"name":"bar"}]}
	//{"hot":120,"tags":[{"icon":"icon"},{"icon":"icon"}]}
}
