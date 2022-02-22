package filter

import (
	"reflect"
)

var nilSlice = make([]int, 0, 0)

// ParseValue 解析字段值
func (t *fieldNodeTree) ParseValue(key, selectScene string, el interface{}) {
	typeOf := reflect.TypeOf(el)
	valueOf := reflect.ValueOf(el)
TakePointerValue: //取指针的值
	switch typeOf.Kind() {
	case reflect.Pointer: //如果是指针类型则取地址重新判断类型
		typeOf = typeOf.Elem()
		goto TakePointerValue
	case reflect.Struct: //如果是字段结构体需要继续递归解析结构体字段所有值
		if typeOf.NumField() == 0 { //如果是一个struct{}{}类型的字段或者是一个空的自定义结构体编码为{}
			t.Key = key
			t.Val = struct{}{}
			return
		}
		for i := 0; i < typeOf.NumField(); i++ {
			jsonTag, ok := typeOf.Field(i).Tag.Lookup("json")
			if !ok || jsonTag == "-" {
				continue
			}
			tag := newSelectTag(jsonTag, selectScene, typeOf.Field(i).Name)
			if tag.IsOmitField || !tag.IsSelect {
				continue
			}
			if valueOf.Kind() == reflect.Pointer {
				valueOf = valueOf.Elem()
			}

			//是否是匿名结构体
			isAnonymous := typeOf.Field(i).Anonymous && tag.IsAnonymous //什么时候才算真正的匿名字段？ Book中Title才算匿名结构体
			//type Book struct {
			//	BookName string `json:"bookName,select(resp)"`
			//	*Page    `json:"page,select(resp)"` // 这个不算匿名字段，为什么？因为tag里打了字段名表示要当作一个字段来对待，
			//	Article    `json:",select(resp)"` //这种情况才是真正的匿名字段，因为tag里字段名为空字符串
			//}
			//

			tree := &fieldNodeTree{
				Key:         tag.UseFieldName,
				ParentNode:  t,
				IsAnonymous: isAnonymous,
			}
			if valueOf.Field(i).Kind() == reflect.Pointer {
				tree.ParseValue(tag.UseFieldName, selectScene, valueOf.Field(i).Elem().Interface())
			} else {
				tree.ParseValue(tag.UseFieldName, selectScene, valueOf.Field(i).Interface())
			}

			if t.IsAnonymous {
				t.AnonymousAddChild(tree)
			} else {
				t.AddChild(tree)
			}
		}
		if t.Child == nil && !t.IsAnonymous {
			//t.Val = struct{}{} //这样表示返回{}

			t.IsAnonymous = true //给他搞成匿名字段的处理方式，直接忽略字段
			//TODO 说明该结构体上没有选择任何字段 应该返回"字段名:{}"？还是直接连字段名都不显示？ 我也不清楚怎么好，后面再说
			//算了反正你啥也不选这字段留着也没任何意义，要就不显示了，至少还能节省一点空间
		}
	case reflect.Bool,
		reflect.String,
		reflect.Float64, reflect.Float32,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		if t.IsAnonymous {
			tree := newFieldNodeTree(t.Key, t, t.Val)
			t.AnonymousAddChild(tree)
		} else {
			t.Val = valueOf.Interface()
			t.Key = key
		}

	case reflect.Map:
		keys := valueOf.MapKeys()
		if len(keys) == 0 { //空map情况下解析为{}
			t.Val = struct{}{}
			return
		}
		for i := 0; i < len(keys); i++ {
			k := keys[i].String()
			nodeTree := newFieldNodeTree(k, t)
			nodeTree.ParseValue(k, selectScene, valueOf.MapIndex(keys[i]).Interface())
			t.AddChild(nodeTree)
		}

	case reflect.Slice, reflect.Array:
		l := valueOf.Len()
		if l == 0 {
			t.Val = nilSlice //空数组空切片直接解析为[],原生的json解析空的切片和数组会被解析为null，真的很烦，遇到脾气暴躁的前端直接跟你开撕。
			return
		}
		t.IsSlice = true
		for i := 0; i < l; i++ {
			node := newFieldNodeTree("", t)
			node.ParseValue("", selectScene, valueOf.Index(i).Interface())
			t.AddChild(node)
		}
	}
}

func SelectMarshal(selectScene string, el interface{}) string {
	tree := newFieldNodeTree("", nil)
	tree.ParseValue("root", selectScene, el)
	return tree.MustJSON()
}

//func decodeAnonymousField(el interface{})map[string]interface{} {
//
//}
