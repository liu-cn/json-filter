package filter

import (
	"encoding"
	"reflect"
)

func (t *fieldNodeTree) parseAny3(valueOf reflect.Value) {
	typeOf := valueOf.Type()
TakePointerValue: //取指针的值
	switch typeOf.Kind() {
	case reflect.Ptr: //如果是指针类型则取值重新判断类型
		valueOf = valueOf.Elem()
		typeOf = typeOf.Elem()
		goto TakePointerValue
	case reflect.Interface:
		if !valueOf.IsNil() {
			valueOf = reflect.ValueOf(valueOf.Interface())
			typeOf = valueOf.Type()
			goto TakePointerValue
		} else {
			//parserNilInterface(t, key)
			t.parserNilInterface()
		}

	case reflect.Struct:
		//parserStruct(typeOf, valueOf, t, t.scene, key, t.isSelect)
		t.parserStruct(typeOf, valueOf)
	case reflect.Bool,
		reflect.String,
		reflect.Float64, reflect.Float32,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		//parserBaseType(valueOf, t, key)
		t.parserBaseType(valueOf)
	case reflect.Map:
		//parserMap(valueOf, t, scene, isSelect)
		t.parserMap(valueOf)
	case reflect.Slice, reflect.Array:
		//parserSliceOrArray(typeOf, valueOf, t, scene, key, isSelect)
		t.parserSliceOrArray(typeOf, valueOf)
	}

}

func (t *fieldNodeTree) parserNilInterface() {
	if t.IsAnonymous {
		tree := &fieldNodeTree{
			Key:        t.Key,
			ParentNode: t,
			Val:        t.Val,
			IsNil:      true,
		}
		t.AnonymousAddChild(tree)
	} else {
		t.Val = nil
		//t.Key = key
		t.IsNil = true
	}
}

func (t *fieldNodeTree) getFieldSelectTag(field reflect.StructField, scene string) tagInfo {
	tagInfoEl := tagInfo{}
	//没开缓存就获取tag
	jsonTag, ok := field.Tag.Lookup("json")
	var tag tag
	if !ok {
		tagInfoEl.omit = true
		return tagInfoEl
	} else {
		if jsonTag == "-" {
			tagInfoEl.omit = true
			return tagInfoEl
		}
		tag = newSelectTag(jsonTag, scene, field.Name)
	}
	tagInfoEl.tag = tag
	return tagInfoEl
}

func (t *fieldNodeTree) parserMap(valueOf reflect.Value) {
	keys := valueOf.MapKeys()
	if len(keys) == 0 { //空map情况下解析为{}
		t.Val = struct{}{}
		return
	}
	for i := 0; i < len(keys); i++ {
		mapIsNil := false
		val := valueOf.MapIndex(keys[i])
	takeValMap:
		if val.Kind() == reflect.Ptr {
			if val.IsNil() {
				mapIsNil = true
				continue
			} else {
				val = val.Elem()
				goto takeValMap
			}
		}
		k := keys[i].String()
		nodeTree := &fieldNodeTree{
			Key:        k,
			ParentNode: t,
		}
		if mapIsNil {
			nodeTree.IsNil = true
			t.AddChild(nodeTree)
		} else {
			nodeTree.parseAny(k, t.scene, val, t.isSelect)
			t.AddChild(nodeTree)
		}
	}
}

func (t *fieldNodeTree) parserBaseType(valueOf reflect.Value) {

	if t.IsAnonymous {
		tree := &fieldNodeTree{
			Key:        t.Key,
			ParentNode: t,
			Val:        t.Val,
		}
		t.AnonymousAddChild(tree)
	} else {
		t.Val = valueOf.Interface()
		//t.Key = key
	}
}

func (t *fieldNodeTree) parserStruct(typeOf reflect.Type, valueOf reflect.Value) {
	if valueOf.CanConvert(timeTypes) { //是time.Time类型或者底层是time.Time类型
		//t.Key = key
		t.Val = valueOf.Interface()
		return
	}
	if typeOf.NumField() == 0 { //如果是一个struct{}{}类型的字段或者是一个空的自定义结构体编码为{}
		//t.Key = key
		t.Val = struct{}{}
		return
	}
	pkgInfo := typeOf.PkgPath() + "." + typeOf.Name()
	for i := 0; i < typeOf.NumField(); i++ {

		var tagInfo tagInfo
		tagInfo = getSelectTag(t.scene, pkgInfo, i, typeOf)
		if !t.isSelect {
			tagInfo = getOmitTag(t.scene, pkgInfo, i, typeOf)
		}

		if tagInfo.omit {
			continue
		}
		tag := tagInfo.tag
		if tag.IsOmitField || !tag.IsSelect {
			continue
		}
		isAnonymous := typeOf.Field(i).Anonymous && tag.IsAnonymous ////什么时候才算真正的匿名字段？ Book中Article才算匿名结构体

		tree := &fieldNodeTree{
			isSelect:    t.isSelect,
			scene:       t.scene,
			Key:         tag.UseFieldName,
			ParentNode:  t,
			IsAnonymous: isAnonymous,
		}
		value := valueOf.Field(i)
		if tag.Function != "" {
			function := valueOf.MethodByName(tag.Function)
			if !function.IsValid() {
				if valueOf.CanAddr() {
					function = valueOf.Addr().MethodByName(tag.Function)
				}
			}
			if function.IsValid() {
				values := function.Call([]reflect.Value{})
				if len(values) != 0 {
					value = values[0]
				}
			}
		}
		if value.Kind() == reflect.Ptr {
		TakeFieldValue:
			if value.Kind() == reflect.Ptr {
				if value.IsNil() {
					if tag.Omitempty {
						continue
					}
					tree.IsNil = true
					t.AddChild(tree)
					continue
				} else {
					value = value.Elem()
					goto TakeFieldValue
				}
			}
		} else {
			if tag.Omitempty {
				if value.IsZero() { //为零值忽略
					continue
				}
			}
		}

		tree.parseAny(tag.UseFieldName, t.scene, value, t.isSelect)

		if t.IsAnonymous {
			t.AnonymousAddChild(tree)
		} else {
			t.AddChild(tree)
		}
	}
	if t.Children == nil && !t.IsAnonymous {
		t.Val = struct{}{} //这样表示返回{}
	}

}

// 如果是切片或者是数组
func (t *fieldNodeTree) parserSliceOrArray(typeOf reflect.Type, valueOf reflect.Value) {
	val1 := valueOf.Interface()

	//先判断是否是byte的切片，byte的切片会被base64编码
	ok := valueOf.CanConvert(byteTypes)
	if ok {
		//t.Key = key
		t.Val = val1
		return
	}

	if typeOf.Kind() == reflect.Array {
		uid, ok := val1.(encoding.TextMarshaler)
		if ok {
			//t.Key = key
			t.Val = uid
			return
		}
	}

	l := valueOf.Len()
	if l == 0 {
		t.Val = emptySlice
		return
	}

	t.IsSlice = true
	for i := 0; i < l; i++ {
		sliceIsNil := false
		node := &fieldNodeTree{
			Key:        "",
			ParentNode: t,
		}
		val := valueOf.Index(i)
	takeValSlice:
		if val.Kind() == reflect.Ptr {
			if val.IsNil() {
				sliceIsNil = true
				continue
			} else {
				val = val.Elem()
				goto takeValSlice
			}
		}

		if sliceIsNil {
			node.IsNil = true
			t.AddChild(node)
		} else {
			node.parseAny3(val)
			t.AddChild(node)
		}
	}
}