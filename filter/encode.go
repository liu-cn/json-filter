package filter

import (
	"reflect"
)

func (t *FieldNodeTree) SetData(key, selectScene string, el interface{}) {
	typeOf := reflect.TypeOf(el)
	valueOf := reflect.ValueOf(el)
TakePointerValue: //取指针的值
	switch typeOf.Kind() {
	case reflect.Pointer: //如果是指针类型则取地址重新判断类型
		typeOf = typeOf.Elem()
		goto TakePointerValue
	case reflect.Struct:

		if typeOf.NumField() == 0 {
			t.Key = key
			t.Val = struct{}{}
			return
		}

		for i := 0; i < typeOf.NumField(); i++ {
			jsonTag, ok := typeOf.Field(i).Tag.Lookup("json")
			if !ok || jsonTag == "-" {
				continue
			}
			tag := NewSelectTag(jsonTag, selectScene, typeOf.Field(i).Name)
			if tag.IsOmitField || !tag.IsSelect {
				continue
			}
			if valueOf.Kind() == reflect.Pointer {
				valueOf = valueOf.Elem()
			}
			tree := NewFieldNodeTree(tag.FieldName)
			if valueOf.Field(i).Kind() == reflect.Pointer {
				tree.SetData(tag.FieldName, selectScene, valueOf.Field(i).Elem().Interface())
			} else {
				tree.SetData(tag.FieldName, selectScene, valueOf.Field(i).Interface())
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

	case reflect.Map:

	}
}

func SelectMarshal(selectScene string, el interface{}) string {
	tree := NewFieldNodeTree("")
	tree.SetData("root", selectScene, el)
	return tree.MustJSON()
}
