package filter

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Filter struct {
	node *fieldNodeTree
}

// Select 直接返回过滤后的数据结构，它可以被json.Marshal解析，直接打印会以过滤后的json字符串展示
func Select(selectScene string, el interface{}) interface{} {
	return jsonFilter(selectScene, el, true)
}

func jsonFilter(selectScene string, el interface{}, isSelect bool) Filter {
	tree := &fieldNodeTree{
		Key:        "",
		isRoot:     true,
		ParentNode: nil,
	}
	tree.parseAny("", selectScene, reflect.ValueOf(el), isSelect)
	return Filter{
		node: tree,
	}
}

// Omit 直接返回过滤后的数据结构，它可以被json.Marshal解析，直接打印会以过滤后的json字符串展示
func Omit(omitScene string, el interface{}) interface{} {
	return jsonFilter(omitScene, el, false)
}

// EnableCache 决定是否启用缓存，默认开启（强烈建议，除非万一缓存模式下出现bug，可以关闭缓存退回曾经的无缓存过滤模式），开启缓存后会有30%-40%的性能提升，开启缓存并没有副作用，只是会让结构体的字段tag常驻内存减少tag字符串处理操作
func EnableCache(enable bool) {
	enableCache = enable
}

// Deprecated
// SelectMarshal 不建议使用，第一个参数填你结构体select标签里的场景，第二个参数是你需要过滤的结构体对象，如果字段的select标签里标注的有该场景那么该字段会被选中。
func SelectMarshal(selectScene string, el interface{}) Filter {
	return jsonFilter(selectScene, el, true)
}

// Deprecated
// OmitMarshal 不建议使用，第一个参数填你结构体omit标签里的场景，第二个参数是你需要过滤的结构体对象，如果字段的omit标签里标注的有该场景那么该字段会被过滤掉
func OmitMarshal(omitScene string, el interface{}) Filter {
	return jsonFilter(omitScene, el, false)
}
func (f Filter) MarshalJSON() ([]byte, error) {
	return useJSONMarshalFunc(f.node.Marshal())
}

// Deprecated
func (f Filter) MastMarshalJSON() []byte {
	return f.node.MustBytes()
}
func (f Filter) MustMarshalJSON() []byte {
	return f.node.MustBytes()
}

// Interface 解析为过滤后待json序列化的map[string]interface{}
func (f Filter) Interface() interface{} {
	return f.node.Marshal()
}

// MustJSON 获取解析过滤后的json字符串，如果中途有错误会panic
func (f Filter) MustJSON() string {
	return f.node.MustJSON()
}

// JSON 获取解析过滤后的json字符串，如果中途有错误会返回错误
func (f Filter) JSON() (string, error) {
	return f.node.JSON()
}

// String fmt.Println() 打印时输出json字符串
func (f Filter) String() string {
	json, err := f.JSON()
	if err != nil {
		return fmt.Sprintf("Filter Err: %s", err.Error())
	}
	return json
}

// EqualJSON 判断两个json字符串是否等价（有相同的键值，不同的顺序）
func EqualJSON(jsonStr1, jsonStr2 string) (bool, error) {
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

//// SelectCache 直接返回过滤后的数据结构，它可以被json.Marshal解析，直接打印会以过滤后的json字符串展示
//func SelectCache(selectScene string, el interface{}) interface{} {
//	return jsonFilterCache(selectScene, el, true)
//}
//
//func jsonFilterCache(selectScene string, el interface{}, isSelect bool) Filter {
//	tree := &fieldNodeTree{
//		Key:        "",
//		isRoot:     true,
//		ParentNode: nil,
//	}
//	tree.parseAny2("", selectScene, reflect.ValueOf(el), isSelect)
//	return Filter{
//		node: tree,
//	}
//}
//
//// OmitCache 直接返回过滤后的数据结构，它可以被json.Marshal解析，直接打印会以过滤后的json字符串展示
//func OmitCache(selectScene string, el interface{}) interface{} {
//	return jsonFilterCache(selectScene, el, false)
//}
