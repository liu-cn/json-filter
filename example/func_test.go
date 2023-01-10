package main

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
)

type Image struct {
	Url     []byte `json:"url,select(img),func(GetUrl)"`
	Path    string `json:"path,select(img),func(GetImagePath)"`
	Name    string `json:"name"`
	Hot     int    `json:"hot,select(img),func(GetHot)"` //热度
	Like    int
	Collect int
	Forward int
}

func (i Image) GetUrl() string {
	return string(i.Url) + ".jpg"
}

// 指针接收器的方法只有在过滤时候传送指针才可以保证此方法被正常调用
func (i *Image) GetImagePath() string {
	return i.Path + i.Name + ".png"
}

// 计算热度
func (i Image) GetHot() int {
	return i.Like * i.Forward * i.Collect
}

func TestFunc(t *testing.T) {
	img := Image{
		Url:     []byte("url"),
		Path:    "path",
		Name:    "_golang",
		Collect: 10,
		Like:    100,
		Forward: 50,
	}
	fmt.Println(filter.Select("img", img))
	//{"hot":50000,"path":"path","url":"url.jpg"}

	fmt.Println(filter.Select("img", &img)) //只有传入指针才可以调用绑定指针接收器方法
	//{"hot":50000,"path":"path_golang.png","url":"url.jpg"}
}
