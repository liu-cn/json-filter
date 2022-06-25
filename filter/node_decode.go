package filter

import (
	"encoding/json"
)

type fieldNodeTree struct {
	Key         string           //字段名
	Val         interface{}      //字段值，基础数据类型，int string，bool 等类型直接存在这里面，如果是struct,切片数组map 类型则字段所有k v会存在ChildNodes里
	IsSlice     bool             //是否是切片，或者数组，
	IsAnonymous bool             //是否是匿名结构体，内嵌结构体，需要把所有字段展开
	IsNil       bool             //该字段值是否为nil
	ParentNode  *fieldNodeTree   //父节点指针，根节点为nil，
	ChildNodes  []*fieldNodeTree //如果是struct则保存所有字段名和值的指针，如果是切片就保存切片的所有值
}

func (t *fieldNodeTree) GetValue() (val interface{}, ok bool) {
	if t.IsAnonymous {
		//如果是匿名字段则不需要再追加这个字段
		return nil, false
	}
	if t.IsNil {
		return nil, true
	}
	if t.ChildNodes == nil {
		return t.Val, true
	}
	if t.IsSlice { //为切片和数组时候key为空
		slices := make([]interface{}, 0, len(t.ChildNodes))
		for i := 0; i < len(t.ChildNodes); i++ {
			value, ok0 := t.ChildNodes[i].GetValue()
			if ok0 {
				slices = append(slices, value)
			}
		}
		return slices, true
	}
	maps := make(map[string]interface{})
	for _, v := range t.ChildNodes {
		value, ok1 := v.GetValue()
		if ok1 {
			maps[v.Key] = value
		}
	}
	return maps, true
}

func (t *fieldNodeTree) Map() map[string]interface{} {
	maps := make(map[string]interface{})
	for _, v := range t.ChildNodes {
		value, ok := (*v).GetValue()
		if ok {
			maps[(*v).Key] = value
		}
	}
	return maps
}
func (t *fieldNodeTree) Slice() interface{} {
	slices := make([]interface{}, 0, len(t.ChildNodes))
	for i := 0; i < len(t.ChildNodes); i++ {
		v, ok := t.ChildNodes[i].GetValue()
		if ok {
			slices = append(slices, v)
		}
	}
	return slices
}

func (t *fieldNodeTree) Marshal() interface{} {
	if t.IsSlice {
		return t.Slice()
	} else { //说明是结构体或者map
		return t.Map()
	}
}

func (t *fieldNodeTree) AddChild(tree *fieldNodeTree) *fieldNodeTree {
	if t.ChildNodes == nil {
		t.ChildNodes = make([]*fieldNodeTree, 0, 3)
	}
	t.ChildNodes = append(t.ChildNodes, tree)
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
		return t
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

func (t *fieldNodeTree) JSON() (string, error) {
	j, err := json.Marshal(t.Marshal())
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func (t *fieldNodeTree) Bytes() ([]byte, error) {
	j, err := json.Marshal(t.Marshal())
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (t *fieldNodeTree) MustBytes() []byte {
	j, err := json.Marshal(t.Marshal())
	if err != nil {
		panic(err)
	}
	return j
}
