package main

import (
	"encoding/json"
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
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
	for _, v := range getTestArrayCase() {
		tests = append(tests, v)
	}
	for _, v := range getTestCase() {
		tests = append(tests, v)
	}

	for i, test := range tests {
		var jsonFilter interface{}
		if test.isSelect {
			jsonFilter = filter.Select(test.scene, test.Struct)
		} else {
			jsonFilter = filter.Omit(test.scene, test.Struct)
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
		},
	}
}

//func TestArray(t *testing.T) {
//	tests:=[]testCase{
//		{
//			isSelect: true,
//			scene: "A",
//			Struct: newArray(),
//			want: `{"A":[{"name":"tag"}]}`,
//		},{
//			isSelect: true,
//			scene: "B",
//			Struct: newArray(),
//			want: `{"B":[{"name":"tag"},{"name":"tag"}]}`,
//		},{
//			isSelect: true,
//			scene: "C",
//			Struct: newArray(),
//			want: `{"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
//		},{
//			isSelect: true,
//			scene: "AP",
//			Struct: newArray(),
//			want: `{"AP":[{"name":"tag"}]}`,
//		},{
//			isSelect: true,
//			scene: "BP",
//			Struct: newArray(),
//			want: `{"BP":[{"name":"tag"},{"name":"tag"}]}`,
//		},{
//			isSelect: true,
//			scene: "CP",
//			Struct: newArray(),
//			want: `{"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
//		},{
//			isSelect: true,
//			scene: "APP",
//			Struct: newArray(),
//			want: `{"APP":[{"name":"tag"}]}`,
//		},{
//			isSelect: true,
//			scene: "BPP",
//			Struct: newArray(),
//			want: `{"BPP":[{"name":"tag"},{"name":"tag"}]}`,
//		},{
//			isSelect: true,
//			scene: "CPP",
//			Struct: newArray(),
//			want: `{"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
//		},
//		//omit
//		{
//			isSelect: false,
//			scene: "A",
//			Struct: newArray(),
//			want: `{"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
//		},{
//			isSelect: false,
//			scene: "B",
//			Struct: newArray(),
//			want: `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
//		},{
//			isSelect: false,
//			scene: "C",
//			Struct: newArray(),
//			want: `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
//		},{
//			isSelect: false,
//			scene: "AP",
//			Struct: newArray(),
//			want: `{"A":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
//		},{
//			isSelect: false,
//			scene: "BP",
//			Struct: newArray(),
//			want: `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
//		},{
//			isSelect: false,
//			scene: "CP",
//			Struct: newArray(),
//			want: `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
//		},{
//			isSelect: false,
//			scene: "APP",
//			Struct: newArray(),
//			want: `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
//		},{
//			isSelect: false,
//			scene: "BPP",
//			Struct: newArray(),
//			want: `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
//		},{
//			isSelect: false,
//			scene: "CPP",
//			Struct: newArray(),
//			want: `{"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}`,
//		},
//	}
//}
