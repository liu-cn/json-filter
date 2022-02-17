package filter

import "strings"

type ExampleModel struct {
	Name string `json:"name,omitempty,select(req|res),omit(chat|profile|article)"`
}

type Tag struct {

	//执行的场景
	SelectScene string
	//该字段是否需要被忽略？
	IsOmitField bool
	//是选中的情况
	IsSelect bool
	//字段名称
	FieldName string
}

func NewSelectTag(tag, selectScene, fieldName string) Tag {

	tagEl := Tag{
		SelectScene: selectScene,
		IsOmitField: true,
	}
	tags := strings.Split(tag, ",")
	tagEl.FieldName = tags[0]
	for _, s := range tags {
		if strings.HasPrefix(s, "select(") {
			selectStr := s[7 : len(s)-1]
			scene := strings.Split(selectStr, "|")
			for _, v := range scene {
				if v == selectScene {
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

func NewOmitTag(tag, selectScene, fieldName string) Tag {

	tagEl := Tag{
		SelectScene: selectScene,
		IsOmitField: false,
	}
	tags := strings.Split(tag, ",")
	tagEl.FieldName = tags[0]
	for _, s := range tags {
		if strings.HasPrefix(s, "omit(") {
			omitStr := s[5 : len(s)-1]
			scene := strings.Split(omitStr, "|")
			for _, v := range scene {
				if v == selectScene {
					//说明选中了tag里的场景
					//tagEl.IsSelect = false
					tagEl.IsOmitField = true
					return tagEl
				}
			}
		}
	}
	return tagEl
}
