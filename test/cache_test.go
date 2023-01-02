package main

import (
	"encoding/json"
	"github.com/liu-cn/json-filter/filter"
	"github.com/liu-cn/pkg/benchmark"
	"testing"
)

func getCacheVal(s string, el interface{}, isSelect bool) string {
	st := ""
	benchmark.TimeAndRes(func() interface{} {
		if isSelect {
			ss := filter.Select(s, el)
			marshal, err := json.Marshal(ss)
			if err != nil {
				panic(err)
			}
			json1 := string(marshal)
			st = json1
			return json1
		} else {
			ss := filter.Omit(s, el)

			marshal, err := json.Marshal(ss)
			if err != nil {
				panic(err)
			}
			json1 := string(marshal)
			st = json1
			return json1
		}
	}, 5)
	return st
}

type result struct {
	s     string
	eq    bool
	json1 string
	json2 string
}

func eq(s string, el interface{}, isSelect bool) result {
	val := getCacheVal(s, el, isSelect)
	var json string
	filter.EnableCache(false)
	defer filter.EnableCache(true)
	if isSelect {
		json = filter.SelectMarshal(s, el).MustJSON()
	} else {
		json = filter.OmitMarshal(s, el).MustJSON()
	}
	return result{
		s:     s,
		json1: val,
		json2: json,
		eq:    json == val,
	}

}

func TestAll(t *testing.T) {

}
