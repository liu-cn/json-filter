package filter

import (
	"encoding/json"
)

type FieldNodeTree struct {
	Key   string
	Val   interface{}
	Child []*FieldNodeTree
}

func (t *FieldNodeTree) GetVal() interface{} {
	if t.Child == nil {
		return t.Val
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
