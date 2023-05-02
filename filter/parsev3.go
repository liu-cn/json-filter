package filter

import (
	"encoding"
	"fmt"
	"reflect"
)

func (t *fieldNodeTree) parseAny_V3(key, scene string, valueOf reflect.Value, isSelect bool) {
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
			parserNilInterfaceV3(t, key)
		}

	case reflect.Struct:

		parserStructV3(typeOf, valueOf, t, scene, key, isSelect)
	case reflect.Bool,
		reflect.String,
		reflect.Float64, reflect.Float32,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		parserBaseTypeV3(valueOf, t, key)
	case reflect.Map:
		parserMapV3(valueOf, t, scene, isSelect)
	case reflect.Slice, reflect.Array:
		parserSliceOrArrayV3(typeOf, valueOf, t, scene, key, isSelect)
	}

}

func parserNilInterfaceV3(t *fieldNodeTree, key string) {
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

func parserMapV3(valueOf reflect.Value, t *fieldNodeTree, scene string, isSelect bool) {
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
			// nodeTree.parseAny(k, scene, val, isSelect)
			nodeTree.parseAny_V3(k, scene, val, isSelect)
			t.AddChild(nodeTree)
		}
	}
}

func parserBaseTypeV3(valueOf reflect.Value, t *fieldNodeTree, key string) {

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

func parserStructV3(typeOf reflect.Type, valueOf reflect.Value, t *fieldNodeTree, scene string, key string, isSelect bool) {
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

	cacheKey := getCacheKey(pkgInfo, scene, isSelect)
	cacheField, ok := fieldCacheVar.filterFieldMap[cacheKey]

	//说明缓存中不存在该字段的过滤信息
	if !ok {
		// fmt.Println("缓存未命中")
		var parentField *field
		//如果是最外层的结构
		if t.isRoot {
			//fieldCacheVar.filterFieldMap[cacheKey] = &field{}
			parentField = new(field)
			fieldCacheVar.filterFieldMap[cacheKey] = parentField
		} else {
			parentField = t.fieldCache
		}

		//遍历这个结构体的所有字段
		for i := 0; i < typeOf.NumField(); i++ {
			var tagInfo tagInfo
			fieldType := typeOf.Field(i)

			//tagInfo = getSelectTag(scene, pkgInfo, i, typeOf)
			if t.isSelect {
				tagInfo = getFieldSelectTag(fieldType, scene)

			} else {
				tagInfo = getFieldOmitTag(fieldType, scene)
				// fmt.Printf("%+v",tagInfo.tag)
			}

			if tagInfo.omit {
				continue
			}
			fieldTag := tagInfo.tag
			if fieldTag.IsOmitField || !fieldTag.IsSelect {
				continue
			}
			// key=fieldTag.UseFieldName
			isAnonymous := fieldType.Anonymous && fieldTag.IsAnonymous
			fieldTag.IsAnonymous = isAnonymous

			//能执行到这一步，说明该字段没被过滤掉，所以缓存应该缓存此字段信息。
			cacheFieldEl := newField(fieldTag, i)
			// fmt.Printf("---------%+v:%v",fieldTag,cacheKey)
			fieldEl, ok1 := fieldCacheVar.filterFieldMap[cacheKey]
			if ok1 {
				//在该父节点的过滤列表里添加此字段
				fieldEl.selectFields = append(fieldEl.selectFields, &cacheFieldEl)
				//fieldCacheVar.filterFieldMap[cacheKey].selectFields = append(fieldCacheVar.filterFieldMap[cacheKey].selectFields, &fieldCache)
			}

			// treeNode := t.newNode(fieldTag.UseFieldName)
			// treeNode.IsAnonymous = isAnonymous
			// treeNode.fieldCache = &cacheFieldEl
			// treeNode.Tag = fieldTag
			treeNode := &fieldNodeTree{
				Key:         fieldTag.UseFieldName,
				ParentNode:  t,
				IsAnonymous: isAnonymous,
				Tag:         fieldTag,
				fieldCache:  &cacheFieldEl,
			}
			value := valueOf.Field(i)
			if fieldTag.Function != "" {
				function := valueOf.MethodByName(fieldTag.Function)
				if !function.IsValid() {
					if valueOf.CanAddr() {
						function = valueOf.Addr().MethodByName(fieldTag.Function)
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
						if fieldTag.Omitempty {
							continue
						}
						treeNode.IsNil = true
						t.AddChild(treeNode)
						continue
					} else {
						value = value.Elem()
						goto TakeFieldValue
					}
				}
			} else {
				if fieldTag.Omitempty {
					if value.IsZero() { //为零值忽略
						continue
					}
				}
			}
			treeNode.parseAny_V3(fieldTag.UseFieldName, scene, value, isSelect)
			if t.IsAnonymous {
				t.AnonymousAddChild(treeNode)
			} else {
				t.AddChild(treeNode)
			}
		}
	} else { //说明缓存取到
		// fmt.Printf("缓存命中:%v\n",cacheKey)
		// fmt.Printf("%+v",cacheField)
		for i := 0; i < len(cacheField.selectFields); i++ {
			fieldInfo := cacheField.selectFields[i]
			tag := fieldInfo.tag
			// tree := t.newNode(tag.UseFieldName)
			tree := &fieldNodeTree{
				Key:         tag.UseFieldName,
				ParentNode:  t,
				IsAnonymous: tag.IsAnonymous,
				Tag:         tag,
				fieldCache:  fieldInfo,
			}
			value := valueOf.Field(fieldInfo.index)
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
			TakeFieldValue1:
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
						goto TakeFieldValue1
					}
				}
			} else {
				if tag.Omitempty {
					if value.IsZero() { //为零值忽略
						continue
					}
				}
			}

			// tree.parseAnyV3(value)
			tree.parseAny_V3(tag.UseFieldName, scene, value, isSelect)
			if t.IsAnonymous {
				t.AnonymousAddChild(tree)
			} else {
				t.AddChild(tree)
			}
		}
	}

	if t.Children == nil && !t.IsAnonymous {
		t.Val = struct{}{} //这样表示返回{}
	}

}

func parserSliceOrArrayV3(typeOf reflect.Type, valueOf reflect.Value, t *fieldNodeTree, scene string, key string, isSelect bool) {

takeV:
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
		goto takeV
	}
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
			node.parseAny_V3("", scene, val, isSelect)
			// node.parseAny("", scene, val, isSelect)
			t.AddChild(node)
		}
	}
}

func EchoCache() {

	for k, v := range fieldCacheVar.filterFieldMap {
		fmt.Println("k", k)
		fmt.Println("kv", v.selectFields)
		// for _, v := range v.selectFields {
		// 	fmt.Println()
		// }
	}

}
