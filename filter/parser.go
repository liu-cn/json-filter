package filter

import (
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"
)

type tagInfo struct {
	tag  tag
	omit bool //表示这个字段忽略
}

func (t *fieldNodeTree) parseAny(key, scene string, valueOf reflect.Value, isSelect bool) {
TakePointerValue: //取指针的值
	if !valueOf.IsValid() {
		parserNilInterface(t, key)
		return
	}

	switch valueOf.Kind() {
	case reflect.Ptr: //如果是指针类型则取值重新判断类型
		if !valueOf.IsNil() {
			valueOf = valueOf.Elem()
			goto TakePointerValue
		} else {
			parserNilInterface(t, key)
			return
		}
	case reflect.Interface:
		if !valueOf.IsNil() {
			valueOf = reflect.ValueOf(valueOf.Interface())
			goto TakePointerValue
		} else {
			parserNilInterface(t, key)
			return
		}
	}

	if v, ok := tryMarshalLeafValue(valueOf); ok {
		parserLeafValue(t, key, v)
		return
	}

	typeOf := valueOf.Type()
	switch typeOf.Kind() {
	case reflect.Struct:
		parserStruct(typeOf, valueOf, t, scene, key, isSelect)
	case reflect.Bool,
		reflect.String,
		reflect.Float64, reflect.Float32,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		parserBaseType(valueOf, t, key)
	case reflect.Map:
		parserMap(valueOf, t, scene, isSelect)
	case reflect.Slice, reflect.Array:
		parserSliceOrArray(typeOf, valueOf, t, scene, key, isSelect)
	}

}

func parserNilInterface(t *fieldNodeTree, key string) {
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
		t.Key = key
		t.IsNil = true
	}
}

func parserLeafValue(t *fieldNodeTree, key string, val interface{}) {
	if t.IsAnonymous {
		tree := &fieldNodeTree{
			Key:        t.Key,
			ParentNode: t,
			Val:        val,
		}
		t.AnonymousAddChild(tree)
		return
	}
	t.Val = val
	t.Key = key
}

func tryMarshalLeafValue(value reflect.Value) (interface{}, bool) {
	if marshaler, ok := asJSONMarshaler(value); ok {
		data, err := marshaler.MarshalJSON()
		if err == nil {
			return json.RawMessage(data), true
		}
	}
	if marshaler, ok := asTextMarshaler(value); ok {
		data, err := marshaler.MarshalText()
		if err == nil {
			return string(data), true
		}
	}
	return nil, false
}

func asJSONMarshaler(value reflect.Value) (json.Marshaler, bool) {
	if !value.IsValid() {
		return nil, false
	}
	if value.CanInterface() {
		if marshaler, ok := value.Interface().(json.Marshaler); ok {
			return marshaler, true
		}
	}
	if ptr := addressableValue(value); ptr.IsValid() && ptr.CanInterface() {
		if marshaler, ok := ptr.Interface().(json.Marshaler); ok {
			return marshaler, true
		}
	}
	return nil, false
}

func asTextMarshaler(value reflect.Value) (encoding.TextMarshaler, bool) {
	if !value.IsValid() {
		return nil, false
	}
	if value.CanInterface() {
		if marshaler, ok := value.Interface().(encoding.TextMarshaler); ok {
			return marshaler, true
		}
	}
	if ptr := addressableValue(value); ptr.IsValid() && ptr.CanInterface() {
		if marshaler, ok := ptr.Interface().(encoding.TextMarshaler); ok {
			return marshaler, true
		}
	}
	return nil, false
}

func addressableValue(value reflect.Value) reflect.Value {
	if !value.IsValid() {
		return reflect.Value{}
	}
	if value.CanAddr() {
		return value.Addr()
	}
	ptr := reflect.New(value.Type())
	ptr.Elem().Set(value)
	return ptr
}

// map的key为数值 bool 和字符串
func isMapKey(t reflect.Value) string {
	switch t.Kind() {
	case reflect.String:
		return t.String()
	case reflect.Bool:
		return fmt.Sprintf("%v", t.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%v", t.Interface())
	default:
		return ""
	}
}
func parserMap(valueOf reflect.Value, t *fieldNodeTree, scene string, isSelect bool) {
	keys := valueOf.MapKeys()
	if len(keys) == 0 { //空map情况下解析为{}
		t.Val = struct{}{}
		return
	}
	for i := 0; i < len(keys); i++ {
		mapIsNil := false
		val := valueOf.MapIndex(keys[i])
		for val.Kind() == reflect.Ptr {
			if val.IsNil() {
				mapIsNil = true
				break
			}
			val = val.Elem()
		}

		key := isMapKey(keys[i])
		if key == "" {
			continue
		}
		k := key
		nodeTree := &fieldNodeTree{
			Key:        k,
			ParentNode: t,
		}
		if mapIsNil {
			nodeTree.IsNil = true
			t.AddChild(nodeTree)
		} else {
			nodeTree.parseAny(k, scene, val, isSelect)
			t.AddChild(nodeTree)
		}
	}
}

func parserBaseType(valueOf reflect.Value, t *fieldNodeTree, key string) {
	parserLeafValue(t, key, valueOf.Interface())
}

func parserStruct(typeOf reflect.Type, valueOf reflect.Value, t *fieldNodeTree, scene string, key string, isSelect bool) {
	if valueOf.CanConvert(timeTypes) { //是time.Time类型或者底层是time.Time类型
		t.Key = key
		t.Val = valueOf.Convert(timeTypes).Interface()
		return
	}
	if typeOf.NumField() == 0 { //如果是一个struct{}{}类型的字段或者是一个空的自定义结构体编码为{}
		t.Key = key
		t.Val = struct{}{}
		return
	}
	for _, meta := range getFieldMetas(typeOf) {
		tagInfo := meta.tagInfo(scene, isSelect)

		if tagInfo.omit {
			continue
		}
		tag := tagInfo.tag
		if tag.IsOmitField || !tag.IsSelect {
			continue
		}
		isAnonymous := meta.anonymous && tag.IsAnonymous ////什么时候才算真正的匿名字段？ Book中Article才算匿名结构体

		tree := &fieldNodeTree{
			Key:         tag.UseFieldName,
			ParentNode:  t,
			IsAnonymous: isAnonymous,
		}
		value := valueOf.Field(meta.index)
		if tag.Function != "" { //解析并调用func选择器
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

		tree.parseAny(tag.UseFieldName, scene, value, isSelect)

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

func parserSliceOrArray(typeOf reflect.Type, valueOf reflect.Value, t *fieldNodeTree, scene string, key string, isSelect bool) {
	val1 := valueOf.Interface()
	ok := valueOf.CanConvert(byteTypes)
	if ok {
		t.Key = key
		t.Val = val1
		return
	}

	if typeOf.Kind() == reflect.Array {
		uid, ok := val1.(encoding.TextMarshaler)
		if ok {
			t.Key = key
			t.Val = uid
			return
		}
	}
	t.IsSlice = true
	l := valueOf.Len()
	if l == 0 {
		t.Val = emptySlice
		return
	}

	for i := 0; i < l; i++ {
		sliceIsNil := false
		node := &fieldNodeTree{
			Key:        "",
			ParentNode: t,
		}
		val := valueOf.Index(i)
		for val.Kind() == reflect.Ptr {
			if val.IsNil() {
				sliceIsNil = true
				break
			}
			val = val.Elem()
		}

		if sliceIsNil {
			node.IsNil = true
			t.AddChild(node)
		} else {
			node.parseAny("", scene, val, isSelect)
			t.AddChild(node)
		}
	}
}
