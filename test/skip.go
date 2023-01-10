package main

import (
	"reflect"
)

//func skip() {
//	info := benchmark.TimeAndRes(func() interface{} {
//		//selectSkip := filter.SelectSkip("h", Us{})
//		return selectSkip
//	}, 5)
//	info.Print(info.LastRes)
//}

func _for(p reflect.Type) {
	for p.Kind() == reflect.Ptr {
		p = p.Elem()
	}
}
func _goto(p reflect.Type) {
take:
	if p.Kind() == reflect.Ptr {
		p = p.Elem()
		goto take
	}
}
