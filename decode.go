package main

import (
	"reflect"
)

func (t *NTree) SetData(key string, el interface{}) {
	typeOf := reflect.TypeOf(el)
	valueOf := reflect.ValueOf(el)
TakePointerValue: //取指针的指
	switch typeOf.Kind() {
	case reflect.Pointer: //如果是指针类型则取地址重新判断类型
		typeOf = typeOf.Elem() //if el类型为int类型的指针(*int) 则此操作相当于*el,取指
		goto TakePointerValue
	case reflect.Struct:
		for i := 0; i < typeOf.NumField(); i++ {
			tree := NewNTree(typeOf.Field(i).Name)
			jsonTag, ok := typeOf.Field(i).Tag.Lookup("json")
			if !ok {
				continue
			}
			if jsonTag == "-" {
				continue
			}
			tag := NewSelectTag(jsonTag, "req", typeOf.Field(i).Name)
			if tag.IsOmitField || !tag.IsSelect {
				continue
			}

			if valueOf.Kind() == reflect.Pointer {
				valueOf = valueOf.Elem()
			}
			if valueOf.Field(i).Kind() == reflect.Pointer {
				tree.SetData(tag.FieldName, valueOf.Field(i).Elem().Interface())
			} else {
				tree.SetData(tag.FieldName, valueOf.Field(i).Interface())
			}
			t.AddChild(tree)
		}
	case reflect.Bool,
		reflect.String,
		reflect.Float64, reflect.Float32,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		t.Val = valueOf.Interface()
		t.Key = key
	}
}

func SelectMarshal(selectScene string, el interface{}) string {
	tree := NewNTree("")
	tree.SetData("root", el)
	return tree.MustJSON()
}
