package filter

import (
	"reflect"
)

var fieldCacheVar fieldCache

func init() {
	fieldCacheVar.maps = make(map[string]*field)
}

type fieldCache struct {
	maps map[string]*field
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
			parserNilInterface(t, key)
		}
	case reflect.Struct:
		parserStructCache(typeOf, valueOf, t, scene, key, isSelect)
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

func parserStructCache(typeOf reflect.Type, valueOf reflect.Value, t *fieldNodeTree, scene string, key string, isSelect bool) {
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
	cacheField, ok := fieldCacheVar.maps[cacheKey]
	if !ok {

		var parentField *field
		if t.isRoot {
			//fieldCacheVar.maps[cacheKey] = &field{}
			parentField = new(field)
			fieldCacheVar.maps[cacheKey] = parentField
		} else {
			parentField = t.fieldCache
		}

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
			tag.IsAnonymous = isAnonymous
			fieldCache := newField(tag, i)
			fieldEl, ok1 := fieldCacheVar.maps[cacheKey]
			if ok1 {
				fieldEl.selectFields = append(fieldEl.selectFields, &fieldCache)
				//fieldCacheVar.maps[cacheKey].selectFields = append(fieldCacheVar.maps[cacheKey].selectFields, &fieldCache)
			}

			tree := &fieldNodeTree{
				Key:         tag.UseFieldName,
				ParentNode:  t,
				IsAnonymous: isAnonymous,
				Tag:         tag,
				fieldCache:  &fieldCache,
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

			tree.parseAny2(tag.UseFieldName, scene, value, isSelect)

			if t.IsAnonymous {
				t.AnonymousAddChild(tree)
			} else {
				t.AddChild(tree)
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

			tree.parseAny2(tag.UseFieldName, scene, value, isSelect)

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
