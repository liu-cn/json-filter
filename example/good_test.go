package main

import "github.com/liu-cn/json-filter/filter"

type UserModel struct {
	UID    uint        `json:"uid,select($any)"` //标记了$any无论选择任何场景都会解析该参数
	Name   string      `json:"name,select(article|profile|chat)"`
	Avatar interface{} `json:"data,select(profile|chat)"`
}

func (u UserModel) ArticleResp() interface{} {
	//这样当你后面想要优化性能时可以在这里进行优化，
	return filter.Select("article", u)
}

func (u UserModel) ProfileResp() interface{} {
	//这样当你后面想要优化性能时可以在这里进行优化，
	return filter.Select("profile", u)
}

func (u UserModel) ChatResp() interface{} {
	//假如性能出现瓶颈，想要优化
	chat := struct {
		UID  uint   `json:"uid"`
		Name string `json:"name"`
	}{
		UID:  u.UID,
		Name: u.Name,
	}
	return chat
}
