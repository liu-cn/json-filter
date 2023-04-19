package main

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"testing"
)

type testField struct {
	number int
	key    string
	val    interface{}
	Fields map[string]*testField
}

func f(x, y interface{}) bool {
	b := reflect.DeepEqual(x, y)
	fmt.Println(b)
	return b
}
func TestF(t *testing.T) {

	s := `{"LangAge":{"Name":"go"},"Name":"boyan"}`
	s2 := `{"Name":"boyan","LangAge":{"Name":"go"}}`

	jsonStr1 := `{"name":"John","age":30,"city":"New York"}`
	jsonStr2 := `{"age":30,"name":"John","city":"New York"}`
	fmt.Println(isEqualJSON(s, s2))
	fmt.Println(isEqualJSON(jsonStr1, jsonStr2))

}

func isEqualJSON(jsonStr1, jsonStr2 string) bool {
	// 两个 JSON 字符串
	//jsonStr1 := `{"name":"John","age":30,"city":"New York"}`
	//jsonStr2 := `{"age":30,"name":"John","city":"New York"}`

	// 解析 JSON 字符串为 map[string]interface{}
	var jsonObj1 map[string]interface{}
	var jsonObj2 map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr1), &jsonObj1); err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(jsonStr2), &jsonObj2); err != nil {
		panic(err)
	}
	return isEqual(jsonObj1, jsonObj2)
}

// 判断两个 map 是否相等
func isEqual(x, y map[string]interface{}) bool {
	if len(x) != len(y) {
		return false
	}
	for k, vx := range x {
		vy, ok := y[k]
		if !ok || !isEqualValue(vx, vy) {
			return false
		}
	}
	return true
}

// 判断两个值是否相等
func isEqualValue(x, y interface{}) bool {
	switch x := x.(type) {
	case nil:
		return y == nil
	case float64:
		// 这里对 float64 进行了特别处理，以避免因精度问题导致的误判。
		// 如果需要更高精度的计算，可以使用 decimal 等库。
		y, ok := y.(float64)
		if !ok {
			return false
		}
		return math.Abs(x-y) < 1e-9
	case []interface{}:
		if y, ok := y.([]interface{}); ok {
			if len(x) != len(y) {
				return false
			}
			for i := range x {
				if !isEqualValue(x[i], y[i]) {
					return false
				}
			}
			return true
		}
	case map[string]interface{}:
		if y, ok := y.(map[string]interface{}); ok {
			return isEqual(x, y)
		}
	}
	return false
}
