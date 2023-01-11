package main

import (
	"github.com/liu-cn/json-filter/filter"
	"testing"
	"time"
)

type Users struct {
	UID    uint   `json:"uid,select(article),omit(article)"`    //select中表示选中的场景(该字段将会使用到的场景)
	Avatar string `json:"avatar,select(article),omit(article)"` //和上面一样此字段在article接口时才会解析该字段

	Nickname string `json:"nickname,select(article|profile)"` //"｜"表示有多个场景都需要这个字段 article接口需要 profile接口也需要

	Sex        int       `json:"sex,select(profile)"`                        //该字段是仅仅profile才使用
	VipEndTime time.Time `json:"vip_end_time,select(profile),omit(article)"` //同上
	Price      string    `json:"price,select(profile)"`                      //同上

	Hobby string    `json:"hobby,omitempty,select($any)"`              //任何场景下为空忽略
	Lang  []LangAge `json:"lang,omitempty,select($any),omit(article)"` //任何场景下为空忽略
}

type LangAge struct {
	Name string `json:"name,omitempty,select($any)"`
	Art  string `json:"art,omitempty,select($any)"`
}

func newUsers() Users {
	return Users{
		UID:        1,
		Nickname:   "boyan",
		Avatar:     "avatar",
		Sex:        1,
		VipEndTime: time.Now().Add(time.Hour * 24 * 365),
		Price:      "999.9",
		Lang: []LangAge{
			{
				Name: "1",
				Art:  "24",
			},
			{
				Name: "2",
				Art:  "35",
			},
		},
	}
}

var str string

func BenchmarkOmitPointerWithCache(b *testing.B) {
	user := newUsers()
	filter.EnableCache(true)
	for i := 0; i < b.N; i++ {
		_ = filter.Omit("article", &user)
	}
}

func BenchmarkSelectPointerWithCache(b *testing.B) {
	user := newUsers()
	filter.EnableCache(true)
	for i := 0; i < b.N; i++ {
		_ = filter.Select("article", &user)
	}
}

func BenchmarkOmitVal(b *testing.B) {
	user := newUsers()
	filter.EnableCache(false)
	for i := 0; i < b.N; i++ {
		_ = filter.Omit("article", user)
	}
}
func BenchmarkSelectVal(b *testing.B) {
	user := newUsers()
	filter.EnableCache(false)
	for i := 0; i < b.N; i++ {
		_ = filter.Select("article", user)
	}
}

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Run("BenchmarkOmitPointerWithCache", BenchmarkOmitPointerWithCache)
		b.Run("BenchmarkOmitVal", BenchmarkOmitVal)
		b.Run("BenchmarkSelectPointerWithCache", BenchmarkSelectPointerWithCache)
		b.Run("BenchmarkSelectVal", BenchmarkSelectVal)
	}
}
