package filter

import (
	"strings"
)

const (
	anySelect = "$any"
	//empty ="$empty"
)

type Tag struct {
	//执行的场景
	SelectScene string
	//该字段是否需要被忽略？
	IsOmitField bool
	//是选中的情况,标识该字段是否需要被解析
	IsSelect bool
	//字段名称
	UseFieldName string
	//IsAnonymous 标识该字段是否是匿名字段
	IsAnonymous bool
	////为空忽略
	Omitempty bool
}

func newSelectTag(tag, selectScene, fieldName string) Tag {

	tagEl := Tag{
		SelectScene: selectScene,
		IsOmitField: true,
	}
	tags := strings.Split(tag, ",")
	tagEl.UseFieldName = fieldName

	if len(tags) < 2 {
		return tagEl
	} else {
		if tags[0] == "" {
			tagEl.IsAnonymous = true
		} else {
			tagEl.UseFieldName = tags[0]
		}
	}
	if tags[1] == "omitempty" {
		tagEl.Omitempty = true
	}

	for _, s := range tags {
		if strings.HasPrefix(s, "select(") {
			selectStr := s[7 : len(s)-1]
			scene := strings.Split(selectStr, "|")
			for _, v := range scene {
				if v == selectScene || v == anySelect {
					//说明选中了tag里的场景,不应该被忽略
					tagEl.IsOmitField = false
					tagEl.IsSelect = true
					return tagEl
				}
			}
		}
	}
	return tagEl
}

func newOmitTag(tag, omitScene, fieldName string) Tag {
	tagEl := Tag{
		SelectScene: omitScene,
		IsOmitField: false,
		IsSelect:    true,
	}
	tags := strings.Split(tag, ",")
	tagEl.UseFieldName = fieldName

	if len(tags) < 2 {
		if len(tags) == 1 {
			if tags[0] != "" {
				tagEl.UseFieldName = tags[0]
			}
		}
		return tagEl
	} else {
		if tags[0] == "" {
			tagEl.IsAnonymous = true
		} else {
			tagEl.UseFieldName = tags[0]
		}
	}
	if tags[1] == "omitempty" {
		tagEl.Omitempty = true
	}

	for _, s := range tags {
		if strings.HasPrefix(s, "omit(") {
			selectStr := s[5 : len(s)-1]
			scene := strings.Split(selectStr, "|")
			for _, v := range scene {
				if v == omitScene || v == anySelect {
					//说明选中了tag里的场景,应该被忽略
					tagEl.IsOmitField = true
					tagEl.IsSelect = false
					return tagEl
				}
			}
		}
	}
	return tagEl
}
