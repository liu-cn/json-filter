package main

import (
	"encoding/json"
	"github.com/liu-cn/json-filter/filter"
	"github.com/liu-cn/pkg/benchmark"
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
	s         string
	eq        bool
	cacheJson string
	noCache   string
}

func eq(s string, el interface{}, isSelect bool) result {
	val := getCacheVal(s, el, isSelect)
	var json string
	filter.EnableCache(false)
	defer filter.EnableCache(true)
	if isSelect {
		json = mustJson(filter.Select(s, el))
		//json = filter.SelectMarshal(s, el).MustJSON()
	} else {
		json = mustJson(filter.Omit(s, el))
		//json = filter.OmitMarshal(s, el).MustJSON()
	}
	return result{
		s:         s,
		cacheJson: val,
		noCache:   json,
		eq:        json == val,
	}

}
