package main

import (
	"encoding/json"
	"github.com/liu-cn/json-filter/filter"
	"reflect"
	"testing"
)

func jsonNav() {
	us := newUs()
	i := filter.Select("h", us)
	marshal, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	no(marshal)
}

func no(v interface{}) {}

func BenchmarkGotoAndFor(b *testing.B) {

	//s := &Us{}
	//of := reflect.TypeOf(s)

	b.Run("_goto", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := &Us{}
			of := reflect.TypeOf(s)
			_goto(of)
		}
	})
	b.Run("_for", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := &Us{}
			of := reflect.TypeOf(s)
			_for(of)
		}
	})
}
