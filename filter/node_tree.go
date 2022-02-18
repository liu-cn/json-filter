package filter

import (
	"encoding/json"
)

type FieldNodeTree struct {
	Key string
	//字段名

	Val interface{}
	//字段值，基础数据类型，int string，bool 等类型直接存在这里面，如果是struct类型则字段所有k v会存在Child里

	IsSlice bool //是否是切片，或者数组，

	//如果是struct则保存所有字段k v，如果是切片就保存切片的所有值
	Child []*FieldNodeTree
}

func (t *FieldNodeTree) GetVal() interface{} {
	if t.Child == nil {
		return t.Val
	}
	if t.IsSlice { //为切片和数组时候key为空
		slices := make([]interface{}, 0, len(t.Child))
		for i := 0; i < len(t.Child); i++ {
			slices = append(slices, t.Child[i].GetVal())
		}
		t.Val = slices
		return slices
	}
	maps := make(map[string]interface{})
	for _, v := range t.Child {
		v.Val = (*v).GetVal()
		maps[(*v).Key] = (*v).Val
	}
	t.Val = maps
	return maps
}

func (t *FieldNodeTree) Map() map[string]interface{} {
	maps := make(map[string]interface{})
	for _, v := range t.Child {
		maps[(*v).Key] = (*v).GetVal()
	}
	return maps
}
func (t *FieldNodeTree) AddChild(tree *FieldNodeTree) *FieldNodeTree {
	t.Child = append(t.Child, tree)
	return t
}

func (t *FieldNodeTree) MustJSON() string {
	j, err := json.Marshal(t.Map())
	if err != nil {
		panic(err)
	}
	return string(j)
}

func NewFieldNodeTree(key string, val ...interface{}) *FieldNodeTree {
	if len(val) > 0 {
		return &FieldNodeTree{
			Key: key,
			Val: val[0],
		}
	}
	return &FieldNodeTree{
		Key: key,
	}
}
