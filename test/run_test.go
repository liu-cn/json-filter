package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/liu-cn/json-filter/filter"
)

type testCase struct {
	isSelect bool
	scene    string
	want     string
	Struct   interface{}
}

func TestAll(t *testing.T) {
	var tests []testCase
	//添加所有测试用例
	tests = append(tests, getTestArrayCase()...)
	tests = append(tests, getTestCase()...)
	t.Run("v1", func(t *testing.T) {
		runTestAll(tests, t, 1)
		runTestAll(tests, t, 1)
	})
	t.Run("v2", func(t *testing.T) {
		runTestAll(tests, t, 2)
		runTestAll(tests, t, 2)
	})
}

func runTestAll(tests []testCase, t *testing.T, version int) {
	for i, test := range tests {
		var jsonFilter interface{}
		if test.isSelect {
			if version == 1 {
				jsonFilter = filter.Select(test.scene, test.Struct)
			} else if version == 2 {
				jsonFilter = filter.SelectCache(test.scene, test.Struct)
			}
		} else {
			if version == 1 {
				jsonFilter = filter.Omit(test.scene, test.Struct)
			} else if version == 2 {
				jsonFilter = filter.OmitCache(test.scene, test.Struct)
			}
		}
		jsonFilterStr, err := json.Marshal(jsonFilter)
		if err != nil {
			t.Error(err)
		}
		wantOk, err := filter.EqualJSON(string(jsonFilterStr), test.want)
		if err != nil {
			t.Error(err)
		}
		if !wantOk {
			t.Errorf("解析的结果不符合预期，scene:%v\n isSelect:%v\nwant:%v\ngot:%v", test.scene, test.isSelect, test.want, string(jsonFilterStr))
		}
		if i == len(tests)-1 {
			fmt.Printf("共测试%v个case", i)
		}
	}
}

func getTestArrayCase() []testCase {
	return []testCase{
		{
			isSelect: true,
			scene:    "A",
			Struct:   newArray(),
			want:     `{"A":[{"name":"tag"}]}`,
		}, {
			isSelect: true,
			scene:    "B",
			Struct:   newArray(),
			want:     `{"B":[{"name":"tag"},{"name":"tag"}]}`,
		}, {
			isSelect: true,
			scene:    "C",
			Struct:   newArray(),
			want:     `{"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
		}, {
			isSelect: true,
			scene:    "AP",
			Struct:   newArray(),
			want:     `{"AP":[{"name":"tag"}]}`,
		}, {
			isSelect: true,
			scene:    "BP",
			Struct:   newArray(),
			want:     `{"BP":[{"name":"tag"},{"name":"tag"}]}`,
		}, {
			isSelect: true,
			scene:    "CP",
			Struct:   newArray(),
			want:     `{"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
		}, {
			isSelect: true,
			scene:    "APP",
			Struct:   newArray(),
			want:     `{"APP":[{"name":"tag"}]}`,
		}, {
			isSelect: true,
			scene:    "BPP",
			Struct:   newArray(),
			want:     `{"BPP":[{"name":"tag"},{"name":"tag"}]}`,
		}, {
			isSelect: true,
			scene:    "CPP",
			Struct:   newArray(),
			want:     `{"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
		},
		//omit
		{
			isSelect: false,
			scene:    "A",
			Struct:   newArray(),
			want:     `{"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
		}, {
			isSelect: false,
			scene:    "B",
			Struct:   newArray(),
			want:     `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
		}, {
			isSelect: false,
			scene:    "C",
			Struct:   newArray(),
			want:     `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
		}, {
			isSelect: false,
			scene:    "AP",
			Struct:   newArray(),
			want:     `{"A":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
		}, {
			isSelect: false,
			scene:    "BP",
			Struct:   newArray(),
			want:     `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
		}, {
			isSelect: false,
			scene:    "CP",
			Struct:   newArray(),
			want:     `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
		}, {
			isSelect: false,
			scene:    "APP",
			Struct:   newArray(),
			want:     `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
		}, {
			isSelect: false,
			scene:    "BPP",
			Struct:   newArray(),
			want:     `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
		}, {
			isSelect: false,
			scene:    "CPP",
			Struct:   newArray(),
			want:     `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
		},
	}
}
func getTestCase() []testCase {
	return []testCase{
		{
			isSelect: true,
			scene:    "lang",
			want:     `{"langAge":[{"name":"c"},{"name":"c++"},{"name":"Go"}],"uid":1}`,
			Struct:   NewUser(),
		},
		{
			isSelect: true,
			scene:    "lookup",
			want:     `{"langAge":[{"arts":[{"profile":{"c":"clang"},"values":["1","2"]}]},{"arts":[{"profile":{"c++":"cpp"},"values":["cpp1","cpp2"]}]},{"arts":[{"profile":{"Golang":"go"},"values":["Golang","Golang1"]}]}],"uid":1}`,
			Struct:   NewUser(),
		}, {
			isSelect: true,
			scene:    "test",
			want:     `{"slice_p":["值"],"slices":["值"],"slices_pp":["值"]}`,
			Struct:   newSliceTest(),
		}, {
			isSelect: false,
			scene:    "test",
			want:     `{"slice_p":["值"],"slices":["值"],"slices_pp":["值"]}`,
			Struct:   newSliceTest(),
		}, {
			isSelect: true,
			scene:    "test",
			want:     `{"m":{"test":"c++从研发到脱发"},"mp":{"test":"c++从研发到脱发"},"mpp":{"test":"c++从研发到脱发"}}`,
			Struct:   newTestMap(),
		}, {
			isSelect: false,
			scene:    "test",
			want:     `{"m":{"test":"c++从研发到脱发"},"mp":{"test":"c++从研发到脱发"},"mpp":{"test":"c++从研发到脱发"}}`,
			Struct:   newTestMap(),
		}, {
			isSelect: true,
			scene:    "a",
			want:     `{"a":"","b":null,"c":null,"d":null}`,
			Struct: struct {
				A interface{} `json:"a,select(a)"`
				B interface{} `json:"b,select(a)"`
				C interface{} `json:"c,select(a)"`
				D interface{} `json:"d,select(a)"`
			}{
				A: "",
			},
		}, {
			isSelect: true,
			scene:    "justName",
			want:     `[{"name":"c"},{"name":"c++"},{"name":"go"}]`,
			Struct:   newTags(),
		}, {
			isSelect: true,
			scene:    "all",
			want:     `[{"icon":"icon-c","id":1,"name":"c"},{"icon":"icon-c++","id":1,"name":"c++"},{"icon":"icon-go","id":1,"name":"go"}]`,
			Struct:   newTags(),
		}, {
			isSelect: true,
			scene:    "chat",
			want:     `[{"icon":"icon-c"},{"icon":"icon-c++"},{"icon":"icon-go"}]`,
			Struct:   newTags(),
		}, {
			isSelect: false,
			scene:    "id",
			want:     `[{"icon":"icon-c","name":"c"},{"icon":"icon-c++","name":"c++"},{"icon":"icon-go","name":"go"}]`,
			Struct:   newTags(),
		}, {
			isSelect: false,
			scene:    "name",
			want:     `[{"icon":"icon-c","id":1},{"icon":"icon-c++","id":1},{"icon":"icon-go","id":1}]`,
			Struct:   newTags(),
		}, {
			isSelect: false,
			scene:    "icon",
			want:     `[{"id":1,"name":"c"},{"id":1,"name":"c++"},{"id":1,"name":"go"}]`,
			Struct:   newTags(),
		}, {
			isSelect: true,
			scene:    "Times",
			want:     `{"birth_time":"2016-01-02 15:04:05","birth_time2":"2016-01-02 15:04:05","nil_birth_time":"0001-01-01 00:00:00","timer":"2016-01-02T15:04:05Z"}`,
			Struct:   NewTimes(),
		}, {
			isSelect: false,
			scene:    "Times",
			want:     `{"birth_time2":"2016-01-02 15:04:05","nil_birth_time":"0001-01-01 00:00:00","timer":"2016-01-02T15:04:05Z"}`,
			Struct:   NewTimes(),
		}, {
			isSelect: true,
			scene:    "article",
			want:     `{"pageInfo":999,"pageNum":1,"title":"c++从研发到脱发"}`,
			Struct:   NewAnonymous(),
		}, {
			isSelect: false,
			scene:    "Anonymous",
			want:     `{"author":"北洛","pageNum":1}`,
			Struct:   NewAnonymous(),
		},
	}
}
