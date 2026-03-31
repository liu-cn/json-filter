package filter

import (
	"fmt"
	"reflect"
)

type Filter struct {
	node *fieldNodeTree
}

// Select returns a filtered value that can be passed directly to json.Marshal.
//
// For backward compatibility this function keeps returning interface{}, but the
// concrete value is always Filter. Use SelectFilter when you want the typed API.
func Select(selectScene string, el interface{}) interface{} {
	return SelectFilter(selectScene, el)
}

func jsonFilter(selectScene string, el interface{}, isSelect bool) Filter {
	tree := &fieldNodeTree{
		Key:        "",
		ParentNode: nil,
	}
	tree.parseAny("", selectScene, reflect.ValueOf(el), isSelect)
	return Filter{
		node: tree,
	}
}

// Omit returns a filtered value that can be passed directly to json.Marshal.
//
// For backward compatibility this function keeps returning interface{}, but the
// concrete value is always Filter. Use OmitFilter when you want the typed API.
func Omit(omitScene string, el interface{}) interface{} {
	return OmitFilter(omitScene, el)
}

// EnableCache 决定是否启用缓存，默认开启（强烈建议，除非万一缓存模式下出现bug，可以关闭缓存退回曾经的无缓存过滤模式），开启缓存后会有30%-40%的性能提升，开启缓存并没有副作用，只是会让结构体的字段tag常驻内存减少tag字符串处理操作
func EnableCache(enable bool) {
	enableCache = enable
}

// SelectFilter returns the typed Filter result for the given select scene.
func SelectFilter(selectScene string, el interface{}) Filter {
	return jsonFilter(selectScene, el, true)
}

// OmitFilter returns the typed Filter result for the given omit scene.
func OmitFilter(omitScene string, el interface{}) Filter {
	return jsonFilter(omitScene, el, false)
}

// Deprecated: use SelectFilter.
func SelectMarshal(selectScene string, el interface{}) Filter {
	return SelectFilter(selectScene, el)
}

// Deprecated: use OmitFilter.
func OmitMarshal(omitScene string, el interface{}) Filter {
	return OmitFilter(omitScene, el)
}

func (f Filter) MarshalJSON() ([]byte, error) {
	return useJSONMarshalFunc(f.node.Marshal())
}

// Deprecated: use MustBytes.
func (f Filter) MastMarshalJSON() []byte {
	return f.MustBytes()
}

// Deprecated: use MustBytes.
func (f Filter) MustMarshalJSON() []byte {
	return f.MustBytes()
}

// Interface returns the filtered result as a Go value ready for JSON encoding.
func (f Filter) Interface() interface{} {
	return f.node.Marshal()
}

// Bytes returns the filtered JSON bytes.
func (f Filter) Bytes() ([]byte, error) {
	return f.node.Bytes()
}

// MustBytes returns the filtered JSON bytes or panics on error.
func (f Filter) MustBytes() []byte {
	return f.node.MustBytes()
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

// Map 过滤后的map结构
func (f Filter) Map() map[string]interface{} {
	return f.node.Map()
}

// Slice 过滤后的切片结构
func (f Filter) Slice() []interface{} {
	slices := make([]interface{}, 0, len(f.node.Children))
	for i := 0; i < len(f.node.Children); i++ {
		v, ok := f.node.Children[i].GetValue()
		if ok {
			slices = append(slices, v)
		}
	}
	return slices
}
