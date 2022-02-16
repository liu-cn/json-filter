package main

import (
	"fmt"
	"reflect"
)

//type Book struct {
//	Page  int    `json:"page,select(req|res|article)"`
//	Price string `json:"price,select(req|res|article)"`
//}
//
//type User struct {
//	Name  string `json:"name,select(req|res|article)"`
//	Age   int    `json:"age,select(req|res|article)"`
//	Hobby string `json:"hobby,select(req|res|article)"`
//	Books []Book `json:"books,select(article)"`
//	B     *Book  `json:"b,select(req|article)"`
//}

//func SelectMarshal(el interface{}) interface{} {
//	k, v := GetKV(el)
//	fmt.Println("k:", k, "v:", v)
//	return v
//}

//func Encode(el interface{}, t *NTree) error {
//	typeOf := reflect.TypeOf(el)
//	if typeOf.Kind() == reflect.Pointer {
//		if typeOf.Elem().Kind() != reflect.Struct {
//			return errors.New("el must be struct")
//		} else {
//			typeOf = typeOf.Elem()
//		}
//	}
//	if typeOf.Kind() != reflect.Struct {
//		return errors.New("el must be struct")
//	}
//
//	valueOf := reflect.ValueOf(el)
//
//	for i := 0; i < typeOf.NumField(); i++ {
//		k, v := GetKV(valueOf.Field(i))
//
//	}
//
//}

func GetKV(el interface{}, tree *NTree) (k string, v interface{}) {
	typeOf := reflect.TypeOf(el)
	valueOf := reflect.ValueOf(el)

TakePointerValue: //取指针的指

	switch typeOf.Kind() {

	case reflect.Pointer: //如果是指针类型则取地址重新判断类型
		typeOf = typeOf.Elem() //if el类型为int类型的指针(*int) 则此操作相当于*el,取指
		goto TakePointerValue

	case reflect.Struct:
		for i := 0; i < typeOf.NumField(); i++ {
			fmt.Print(typeOf.Field(i).Name, "\t")
			fmt.Print(valueOf.Field(i).Kind(), "\n")
		}
		return typeOf.Name(), valueOf.Interface()
	case reflect.Bool,
		reflect.String,
		reflect.Float64, reflect.Float32,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return typeOf.Name(), valueOf.Interface()
	}

	return "unknown", nil
}

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

//func (t *NTree) ()  {
//
//}

func SelectMarshal(selectScene string, el interface{}) string {
	tree := NewNTree("")
	tree.SetData("root", el)
	return tree.MustJSON()
}
