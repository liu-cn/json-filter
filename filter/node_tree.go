package filter

import (
	"encoding/json"
	"fmt"
)

type fieldNodeTree struct {
	Key string
	//字段名

	Val interface{}
	//字段值，基础数据类型，int string，bool 等类型直接存在这里面，如果是struct,切片数组map 类型则字段所有k v会存在Child里

	IsSlice bool //是否是切片，或者数组，

	//是否是匿名结构体，内嵌结构体，需要把所有字段展开
	IsAnonymous bool

	//父节点指针，可以为nil，
	ParentNode *fieldNodeTree

	//如果是struct则保存所有字段名和值的指针，如果是切片就保存切片的所有值
	Child []*fieldNodeTree
}

func (t *fieldNodeTree) GetValue() interface{} {
	if t.IsAnonymous {
		//如果是匿名字段则不需要再追加这个字段
		return nil
	}
	if t.Child == nil {
		return t.Val
	}
	if t.IsSlice { //为切片和数组时候key为空
		slices := make([]interface{}, 0, len(t.Child))
		for i := 0; i < len(t.Child); i++ {
			slices = append(slices, t.Child[i].GetValue())
		}
		//t.Val = slices
		return slices
	}
	maps := make(map[string]interface{})
	for _, v := range t.Child {
		value := (*v).GetValue()
		if value != nil {
			maps[(*v).Key] = value
		}
	}
	//t.Val = maps
	return maps
}

func (t *fieldNodeTree) Map() map[string]interface{} {
	maps := make(map[string]interface{})
	for _, v := range t.Child {
		value := (*v).GetValue()
		if value != nil {
			maps[(*v).Key] = value
		}
	}
	return maps
}

func (t *fieldNodeTree) Marshal() interface{} {
	if t.IsSlice {
		slices := make([]interface{}, 0, len(t.Child))
		for i := 0; i < len(t.Child); i++ {
			v := t.Child[i].GetValue()
			if v != nil {
				slices = append(slices, v)
			}
		}
		return slices
	} else {
		return t.Map()
	}
}

func (t *fieldNodeTree) AddChild(tree *fieldNodeTree) *fieldNodeTree {
	t.Child = append(t.Child, tree)
	return t
}

//如果是以下这种情况，层层无限嵌入匿名字段，最深层Page的字段也需要添加到最上层User字段里，
//User正确解析的结果应该是：{"bookName":"book","pageInfo":10,"userName":"boyan"}，这里根据匿名结构体是否命名来决定是否需要展开结构体字段
//type User struct {
//	UserName string `json:"userName,select(all)"`
//	Book     `json:",select(all)"`
//}
//
//type Book struct {
//	BookName string `json:"bookName,select(all)"`
//	Page     `json:",select(all)"`
//}
//
//type Page struct {
//	PageInfo int `json:",select(all)"`
//}

// GetParentNodeInsertPosition 递归找到最上层可以插入的节点
func (t *fieldNodeTree) GetParentNodeInsertPosition() *fieldNodeTree {
	if t.ParentNode == nil {
		panic(fmt.Sprintf("父节点为nil %+v", t))
	}

	//层层向父节点递归，直到寻找到不是匿名字段的节点，向该节点的child中添加数据
	if t.ParentNode.IsAnonymous {
		return t.ParentNode.GetParentNodeInsertPosition()
	}
	return t.ParentNode
}

// AnonymousAddChild 匿名字段向父节点追加操作
func (t *fieldNodeTree) AnonymousAddChild(tree *fieldNodeTree) *fieldNodeTree {
	t.GetParentNodeInsertPosition().AddChild(tree)
	return t
}

// MustJSON 如果解析失败会直接panic掉
func (t *fieldNodeTree) MustJSON() string {
	j, err := json.Marshal(t.Marshal())

	//j, err := sonic.Marshal(t.Marshal()) //这个目前兼容性不是特别好，先用官方库
	if err != nil {
		panic(err)
	}
	return string(j)
}

func newFieldNodeTree(key string, parentNode *fieldNodeTree, val ...interface{}) *fieldNodeTree {
	f := &fieldNodeTree{
		Key:        key,
		ParentNode: parentNode,
	}
	if len(val) > 0 {
		f.Val = val[0]
	}
	return f
}
