package main

import (
	"encoding/json"
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
)

type (
	Tag struct {
		Name string `json:"name,select($any)"`
		Icon string `json:"icon,omit($any)"`
	}

	Array struct {
		A   [1]Tag   `json:"A,select(A|all),omit(A|all)"`
		B   [2]Tag   `json:"B,select(B|all),omit(B|all)"`
		C   [3]Tag   `json:"C,select(C|all),omit(C|all)"`
		AP  *[1]Tag  `json:"AP,select(AP|all),omit(AP|all)"`
		BP  *[2]Tag  `json:"BP,select(BP|all),omit(BP|all)"`
		CP  *[3]Tag  `json:"CP,select(CP|all),omit(CP|all)"`
		APP *[1]*Tag `json:"APP,select(APP|all),omit(APP|all)"`
		BPP *[2]*Tag `json:"BPP,select(BPP|all),omit(BPP|all)"`
		CPP *[3]*Tag `json:"CPP,select(CPP|all),omit(CPP|all)"`
	}
)

var arrayWants = []string{
	"A",
	"B",
	"C",
	"AP",
	"BP",
	"CP",
	"APP",
	"BPP",
	"CPP",
}

func getKeys(jsonStr string) string {
	maps := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &maps)
	if err != nil {
		panic(err)
	}
	keys := ""
	for k := range maps {
		keys += k + ","
	}
	return keys
}

func newArray() *Array {

	tag := Tag{Name: "tag"}
	tags1 := [1]Tag{tag}
	tags2 := [2]Tag{tag, tag}
	tags3 := [3]Tag{tag, tag, tag}
	tags1p := &[1]Tag{tag}
	tags2p := &[2]Tag{tag, tag}
	tags3p := &[3]Tag{tag, tag, tag}
	tags1pp := &[1]*Tag{&tag}
	tags2pp := &[2]*Tag{&tag, &tag}
	tags3pp := &[3]*Tag{&tag, &tag, &tag}

	arr := &Array{
		A:   tags1,
		B:   tags2,
		C:   tags3,
		AP:  tags1p,
		BP:  tags2p,
		CP:  tags3p,
		APP: tags1pp,
		BPP: tags2pp,
		CPP: tags3pp,
	}
	return arr
}

func TestSelectArray(t *testing.T) {
	for _, want := range arrayWants {
		fmt.Println(want, ":", filter.Select(want, newArray()))
	}
	//=== RUN   TestSelectArray
	//A : {"A":[{"name":"tag"}]}
	//B : {"B":[{"name":"tag"},{"name":"tag"}]}
	//C : {"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}
	//AP : {"AP":[{"name":"tag"}]}
	//BP : {"BP":[{"name":"tag"},{"name":"tag"}]}
	//CP : {"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}
	//APP : {"APP":[{"name":"tag"}]}
	//BPP : {"BPP":[{"name":"tag"},{"name":"tag"}]}
	//CPP : {"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}
	//--- PASS: TestSelectArray (0.00s)
	//PASS

}
func TestSelectOmitTag(t *testing.T) {
	fmt.Println(filter.Select("any", newArray().CPP))
	fmt.Println(filter.Select("any", *newArray().CPP))
	fmt.Println(filter.Omit("any", newArray().CPP))
	fmt.Println(filter.Omit("any", *newArray().CPP))
	//=== RUN   TestSelectOmitTag
	//[{"name":"tag"},{"name":"tag"},{"name":"tag"}]
	//[{"name":"tag"},{"name":"tag"},{"name":"tag"}]
	//[{"name":"tag"},{"name":"tag"},{"name":"tag"}]
	//[{"name":"tag"},{"name":"tag"},{"name":"tag"}]
	//--- PASS: TestSelectOmitTag (0.00s)
	//PASS
}

func TestOmitArray(t *testing.T) {
	for _, want := range arrayWants {
		omit := filter.Omit(want, newArray())
		fmt.Println(want, ":", "keys:", getKeys(mustJson(omit)), omit)
	}
	//=== RUN   TestOmitArray
	//A : keys: CP,CPP,AP,APP,B,BP,BPP,C, {"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}
	//B : keys: BPP,C,CP,CPP,A,AP,APP,BP, {"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}
	//C : keys: A,AP,APP,B,BP,BPP,CP,CPP, {"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}
	//AP : keys: CPP,A,APP,B,BP,BPP,C,CP, {"A":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}
	//BP : keys: APP,B,BPP,C,CP,CPP,A,AP, {"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}
	//CP : keys: CPP,A,AP,APP,B,BP,BPP,C, {"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}
	//APP : keys: BP,BPP,C,CP,CPP,A,AP,B, {"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}
	//BPP : keys: A,AP,APP,B,BP,C,CP,CPP, {"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CPP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}
	//CPP : keys: BPP,C,CP,A,AP,APP,B,BP, {"A":[{"name":"tag"}],"AP":[{"name":"tag"}],"APP":[{"name":"tag"}],"B":[{"name":"tag"},{"name":"tag"}],"BP":[{"name":"tag"},{"name":"tag"}],"BPP":[{"name":"tag"},{"name":"tag"}],"C":[{"name":"tag"},{"name":"tag"},{"name":"tag"}],"CP":[{"name":"tag"},{"name":"tag"},{"name":"tag"}]}
	//--- PASS: TestOmitArray (0.00s)
	//PASS

}

func TestEQ(t *testing.T) {

	t.Run("select_eq", func(t *testing.T) {
		for _, want := range arrayWants {
			arr := newArray()
			ptr := mustJson(filter.Select(want, arr))
			val := mustJson(filter.Select(want, *arr))
			if ptr != val {
				t.Errorf("select传递指针和值结果不相等,want%s,ptr:%s,val:%s", want, ptr, val)
			}
		}
	})
	t.Run("omit_eq", func(t *testing.T) {
		for _, want := range arrayWants {
			arr := newArray()
			ptr := mustJson(filter.Omit(want, arr))
			val := mustJson(filter.Omit(want, *arr))
			if ptr != val {
				t.Errorf("omit传递指针和值结果不相等,want%s,ptr:%s,val:%s", want, ptr, val)
			}
		}
	})
	//=== RUN   TestEQ
	//--- PASS: TestEQ (0.00s)
	//=== RUN   TestEQ/select_eq
	//--- PASS: TestEQ/select_eq (0.00s)
	//=== RUN   TestEQ/omit_eq
	//--- PASS: TestEQ/omit_eq (0.00s)
	//PASS

}
