package filter

import (
	"encoding/json"
	"reflect"
)

func equalJSON(jsonStr1, jsonStr2 string) (bool, error) {
	var i interface{}
	var i2 interface{}
	err := json.Unmarshal([]byte(jsonStr1), &i)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal([]byte(jsonStr2), &i2)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(i, i2), nil
}

// EqualJSON 判断两个或者多个json字符串是否等价（有相同的键值，不同的顺序）
func EqualJSON(jsonStr1, jsonStr2 string, moreJson ...string) (bool, error) {
	if len(moreJson) == 0 {
		return equalJSON(jsonStr1, jsonStr2)
	}
	equal, err := equalJSON(jsonStr1, jsonStr2)
	if err != nil {
		return false, err
	}
	if !equal {
		return false, nil
	}
	for _, js := range moreJson {
		eq, err := equalJSON(jsonStr1, js)
		if err != nil {
			return false, err
		}
		if !eq {
			return false, err
		}
	}
	return true, nil

}
