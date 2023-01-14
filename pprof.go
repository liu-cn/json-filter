package main

import (
	"github.com/liu-cn/json-filter/filter"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

var j interface{}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	filter.EnableCache(true)
	for {
		j = filter.Select("article", newUser())
	}
}

type User struct {
	UID    uint   `json:"uid,select(article)"`    //select中表示选中的场景(该字段将会使用到的场景)
	Avatar string `json:"avatar,select(article)"` //和上面一样此字段在article接口时才会解析该字段

	Nickname string `json:"nickname,select(article|profile)"` //"｜"表示有多个场景都需要这个字段 article接口需要 profile接口也需要

	Sex        int       `json:"sex,select(profile)"`          //该字段是仅仅profile才使用
	VipEndTime time.Time `json:"vip_end_time,select(profile)"` //同上
	Price      string    `json:"price,select(profile)"`        //同上

	Hobby string    `json:"hobby,omitempty,select($any)"` //任何场景下为空忽略
	Lang  []LangAge `json:"lang,omitempty,select($any)"`  //任何场景下为空忽略
}

type LangAge struct {
	Name string `json:"name,omitempty,select($any)"`
	Art  string `json:"art,omitempty,select($any)"`
}

func newUser() User {
	return User{
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
