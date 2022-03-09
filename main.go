package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/liu-cn/json-filter/filter"
)

//同一个结构体，你可能想要在article 接口下只返回UID Avatar Nickname这三个字段就够了，其他字段不想要暴露
//另外你还想在profile 接口下返回 Nickname Sex VipEndTime Price 这四个字段，其他字段不想暴露
//这样的情况有很多，想要复用一个结构体来随心所欲的构造自己想要的json数据结构，可以看一个简单的demo

type User struct {
	UID    uint   `json:"uid,select(article)"`    //select中表示选中的场景(该字段将会使用到的场景)
	Avatar string `json:"avatar,select(article)"` //和上面一样此字段在article接口时才会解析该字段

	Nickname string `json:"nickname,select(article|profile)"` //"｜"表示有多个场景都需要这个字段 article接口需要 profile接口也需要

	Sex        int       `json:"sex,select(profile)"`          //该字段是仅仅profile才使用
	VipEndTime time.Time `json:"vip_end_time,select(profile)"` //同上
	Price      string    `json:"price,select(profile)"`        //同上

	Hobby string  `json:"hobby,omitempty,select($any)"` //任何场景下为空忽略
	Lang  LangAge `json:"lang,omitempty,select($any)"`  //任何场景下为空忽略
}

type LangAge struct {
	Name string `json:"name,omitempty,select($any)"`
	Art  string `json:"art,omitempty,select($any)"`
}

func main() {

	user := User{
		UID:        1,
		Nickname:   "boyan",
		Avatar:     "avatar",
		Sex:        1,
		VipEndTime: time.Now().Add(time.Hour * 24 * 365),
		Price:      "999.9",
	}

	marshal, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal)) //以下是官方的json解析输出结果：可以看到所有的字段都被解析了出来
	//{"uid":1,"avatar":"avatar","nickname":"boyan","sex":1,"vip_end_time":"2023-03-09T20:52:19.524188+08:00","price":"999.9","lang":{}}

	fmt.Println(filter.SelectMarshal("article", user).MustJSON()) //以下是通过json-filter 过滤后的json，此输出是article接口下的json
	//{"avatar":"avatar","nickname":"boyan","uid":1}

	fmt.Println(filter.SelectMarshal("profile", user).MustJSON()) //profile接口下
	//{"nickname":"boyan","price":"999.9","sex":1,"vip_end_time":"2023-03-09T20:52:19.524188+08:00"}

	user2 := user
	user2.Hobby = "code"
	user2.Lang = LangAge{
		Name: "Go",
	}

	fmt.Println(filter.SelectMarshal("profile", user2).MustJSON())
	//{"hobby":"code","lang":{"name":"Go"},"nickname":"boyan","price":"999.9","sex":1,"vip_end_time":"2023-03-09T20:58:45.946248+08:00"}
}
