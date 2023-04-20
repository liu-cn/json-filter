package main

import (
	"encoding/json"
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
)

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
