package filter

import (
	"reflect"
)

var fieldCacheVar fieldCache

func init() {
	fieldCacheVar.filterFieldMap = make(map[string]*field)
}

type fieldCache struct {
	filterFieldMap map[string]*field
}
type field struct {
	index        int //该字段所处结构体的索引位置
	tag          tag
	selectFields []*field
}

func newField(t tag, idx int) field {
	return field{
		tag:   t,
		index: idx,
	}
}

func getCacheKey(pkgInfo string, scene string, isSelect bool) string {
	mode := ".s"
	if !isSelect {
		mode = ".o"
	}
	return pkgInfo + "." + scene + mode
}
func (t *fieldNodeTree) parseAny2(key, scene string, valueOf reflect.Value, isSelect bool) {
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
			t.parserNilInterface()
		}
	case reflect.Struct:
		t.parserStructCache(typeOf, valueOf)
	case reflect.Bool,
		reflect.String,
		reflect.Float64, reflect.Float32,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		t.parserBaseType(valueOf)
	case reflect.Map:
		parserMap(valueOf, t, scene, isSelect)
	case reflect.Slice, reflect.Array:
		parserSliceOrArray(typeOf, valueOf, t, scene, key, isSelect)
	}

}

func (t *fieldNodeTree) parserStructCache(typeOf reflect.Type, valueOf reflect.Value) {
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

	cacheKey := getCacheKey(pkgInfo, t.scene, t.isSelect)
	cacheField, ok := fieldCacheVar.filterFieldMap[cacheKey]

	//说明缓存中不存在该字段的过滤信息
	if !ok {
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
			tagInfo = getFieldSelectTag(fieldType, t.scene)
			//tagInfo = getSelectTag(scene, pkgInfo, i, typeOf)
			if !t.isSelect {
				tagInfo = getFieldOmitTag(fieldType, t.scene)
			}

			if tagInfo.omit {
				continue
			}
			fieldTag := tagInfo.tag
			if fieldTag.IsOmitField || !fieldTag.IsSelect {
				continue
			}
			isAnonymous := fieldType.Anonymous && fieldTag.IsAnonymous
			fieldTag.IsAnonymous = isAnonymous

			//能执行到这一步，说明该字段没被过滤掉，所以缓存应该缓存此字段信息。
			cacheFieldEl := newField(fieldTag, i)
			fieldEl, ok1 := fieldCacheVar.filterFieldMap[cacheKey]
			if ok1 {
				//在该父节点的过滤列表里添加此字段
				fieldEl.selectFields = append(fieldEl.selectFields, &cacheFieldEl)
				//fieldCacheVar.filterFieldMap[cacheKey].selectFields = append(fieldCacheVar.filterFieldMap[cacheKey].selectFields, &fieldCache)
			}

			treeNode := t.newNode(fieldTag.UseFieldName)
			treeNode.IsAnonymous = isAnonymous
			treeNode.fieldCache = &cacheFieldEl
			treeNode.Tag = fieldTag
			//tree := &fieldNodeTree{
			//	Key:         fieldTag.UseFieldName,
			//	ParentNode:  t,
			//	IsAnonymous: isAnonymous,
			//	Tag:         fieldTag,
			//	fieldCache:  &cacheFieldEl,
			//}
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

			treeNode.parseAny3(value)

			if t.IsAnonymous {
				t.AnonymousAddChild(treeNode)
			} else {
				t.AddChild(treeNode)
			}
		}
	} else { //说明缓存取到
		for i := 0; i < len(cacheField.selectFields); i++ {
			fieldInfo := cacheField.selectFields[i]
			tag := fieldInfo.tag
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

			tree.parseAny3(value)

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
