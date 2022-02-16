package main

import (
	"encoding/json"
	_ "encoding/json"
)

type NTree struct {
	Key   string
	Val   interface{}
	Child []*NTree
}

func (t *NTree) GetVal() interface{} {
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

func (t *NTree) Map() map[string]interface{} {
	maps := make(map[string]interface{})
	for _, v := range t.Child {
		maps[(*v).Key] = (*v).GetVal()
	}
	return maps
}
func (t *NTree) AddChild(tree *NTree) *NTree {
	t.Child = append(t.Child, tree)
	return t
}

func (t *NTree) MustJSON() string {
	j, err := json.Marshal(t.Map())
	if err != nil {
		panic(err)
	}
	return string(j)
}

func NewNTree(key string, val ...interface{}) *NTree {
	if len(val) > 0 {
		return &NTree{
			Key: key,
			Val: val[0],
		}
	}

	return &NTree{
		Key: key,
	}
}
