package filter

import (
	"encoding"
	"reflect"
)

type tagInfo struct {
	tag  tag
	omit bool //表示这个字段忽略
}

func (t *fieldNodeTree) parseAny(key, scene string, valueOf reflect.Value, isSelect bool) {
	typeOf := valueOf.Type()
TakePointerValue: //取指针的值
	switch typeOf.Kind() {
	case reflect.Ptr: //如果是指针类型则取值重新判断类型
		valueOf = valueOf.Elem()
		typeOf = typeOf.Elem()
		goto TakePointerValue
	case reflect.Interface:
		valueOf = reflect.ValueOf(valueOf.Interface())
		typeOf = valueOf.Type()
		goto TakePointerValue
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

func getFieldOmitTag(field reflect.StructField, scene string) tagInfo {
	tagInfoEl := tagInfo{}
	//没开缓存就获取tag
	jsonTag, ok := field.Tag.Lookup("json")
	var tag tag
	if !ok {
		tag = newOmitNotTag(scene, field.Name)
	} else {
		if jsonTag == "-" {
			tagInfoEl.omit = true
			return tagInfoEl
		}
		tag = newOmitTag(jsonTag, scene, field.Name)
	}
	tagInfoEl.tag = tag
	return tagInfoEl
}
func getFieldSelectTag(field reflect.StructField, scene string) tagInfo {
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

func getOmitTag(scene string, pkgInfo string, i int, typeOf reflect.Type) tagInfo {
	omitTag := tagInfo{}

	if !enableCache { //没开缓存就获取tag
		omitTag = getFieldOmitTag(typeOf.Field(i), scene)
		return omitTag
	}
	fieldName := typeOf.Field(i).Name
	cacheKey := tagCache.key(pkgInfo, scene, fieldName, false)
	tagEl, exist := tagCache.c[cacheKey]
	if !exist { //如果缓存里没取到
		omitTag = getFieldOmitTag(typeOf.Field(i), scene)
		tagCache.c[cacheKey] = omitTag.tag
		return omitTag
	}
	omitTag.tag = tagEl

	return omitTag
}

func getSelectTag(scene string, pkgInfo string, i int, typeOf reflect.Type) tagInfo {
	selectTag := tagInfo{}

	if !enableCache {
		return getFieldSelectTag(typeOf.Field(i), scene)
	}

	fieldName := typeOf.Field(i).Name
	cacheKey := tagCache.key(pkgInfo, scene, fieldName, true)
	tagEl, exist := tagCache.c[cacheKey]
	if !exist { //如果缓存里没取到
		selectTag = getFieldSelectTag(typeOf.Field(i), scene)
		tagCache.c[cacheKey] = selectTag.tag
		return selectTag
	}
	selectTag.tag = tagEl
	return selectTag
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
			nodeTree.parseAny(k, scene, val, isSelect)
			t.AddChild(nodeTree)
		}
	}
}

func parserBaseType(valueOf reflect.Value, t *fieldNodeTree, key string) {

	if t.IsAnonymous {
		tree := &fieldNodeTree{
			Key:        t.Key,
			ParentNode: t,
			Val:        t.Val,
		}
		t.AnonymousAddChild(tree)
	} else {
		t.Val = valueOf.Interface()
		t.Key = key
	}
}

func parserStruct(typeOf reflect.Type, valueOf reflect.Value, t *fieldNodeTree, scene string, key string, isSelect bool) {
	if valueOf.CanConvert(timeTypes) { //是time.Time类型或者底层是time.Time类型
		t.Key = key
		t.Val = valueOf.Interface()
		return
	}
	if typeOf.NumField() == 0 { //如果是一个struct{}{}类型的字段或者是一个空的自定义结构体编码为{}
		t.Key = key
		t.Val = struct{}{}
		return
	}
	pkgInfo := typeOf.PkgPath() + "." + typeOf.Name()
	for i := 0; i < typeOf.NumField(); i++ {

		var tagInfo tagInfo
		tagInfo = getSelectTag(scene, pkgInfo, i, typeOf)
		if !isSelect {
			tagInfo = getOmitTag(scene, pkgInfo, i, typeOf)
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
			node.parseAny("", scene, val, isSelect)
			t.AddChild(node)
		}
	}
}
