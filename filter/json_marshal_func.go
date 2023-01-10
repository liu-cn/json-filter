package filter

import "encoding/json"

type jsonMarshalFunc = func(v interface{}) ([]byte, error)

// var useJSONMarshal = false
var useJSONMarshalFunc = json.Marshal

// SetJSONMarshal 默认使用官方的json.Marshal 解析过滤后的数据结构，可以自定义json解析方法 比如换成字节的github.com/bytedance/sonic 这个进行json解析，
// filter.SetJSONMarshal(sonic.Marshal)就可以了，但是sonic(速度是官方的2-3倍)这个库好像需要go 1.15以上，可以根据自己的需要选择自己需要的json解析库
// 如果对性能要求没有那么高没必要换，这个方法在任意位置调用一次就可以了
func SetJSONMarshal(fn jsonMarshalFunc) {
	if fn != nil {
		useJSONMarshalFunc = fn
	}
}
