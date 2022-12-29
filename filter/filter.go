package filter

import "fmt"

type Filter struct {
	node *fieldNodeTree
}

func (f Filter) MarshalJSON() ([]byte, error) {
	return f.node.Bytes()
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

// SelectMarshal 不建议使用，第一个参数填你结构体select标签里的场景，第二个参数是你需要过滤的结构体对象，如果字段的select标签里标注的有该场景那么该字段会被选中。
func SelectMarshal(selectScene string, el interface{}) Filter {
	if enableCache {
		return selectWithCache(selectScene, el)
	}
	return selectMarshal(selectScene, el)
}
func selectMarshal(selectScene string, el interface{}) Filter {
	tree := &fieldNodeTree{
		Key:        "",
		ParentNode: nil,
	}
	tree.ParseSelectValue("", selectScene, el)
	return Filter{
		node: tree,
	}
}

// Select 直接返回过滤后的数据结构，相当于直接SelectMarshal后再调用Interface方法
func Select(selectScene string, el interface{}) interface{} {
	if enableCache {
		return selectWithCache(selectScene, el)
	}
	return selectMarshal(selectScene, el)
}

// selectWithCache 直接返回过滤后的数据结构，相当于直接SelectMarshal后再调用Interface方法
func selectWithCache(selectScene string, el interface{}) Filter {
	tree := &fieldNodeTree{
		Key:        "",
		ParentNode: nil,
	}
	tree.ParseSelectValueWithCache("", selectScene, el)
	return Filter{
		node: tree,
	}
}

// Omit 直接返回过滤后的数据结构，相当于直接OmitMarshal后再调用Interface方法
func Omit(omitScene string, el interface{}) interface{} {
	if enableCache {
		return omitWithCache(omitScene, el)
	}
	return omitMarshal(omitScene, el)
}

// OmitMarshal 不建议使用，第一个参数填你结构体omit标签里的场景，第二个参数是你需要过滤的结构体对象，如果字段的omit标签里标注的有该场景那么该字段会被过滤掉
func OmitMarshal(omitScene string, el interface{}) Filter {
	if enableCache {
		return omitWithCache(omitScene, el)
	}
	return omitMarshal(omitScene, el)
}

func omitMarshal(omitScene string, el interface{}) Filter {
	tree := &fieldNodeTree{
		Key:        "",
		ParentNode: nil,
	}
	tree.ParseOmitValue("", omitScene, el)
	return Filter{
		node: tree,
	}
}

// omitWithCache 第一个参数填你结构体omit标签里的场景，第二个参数是你需要过滤的结构体对象，如果字段的omit标签里标注的有该场景那么该字段会被过滤掉
func omitWithCache(omitScene string, el interface{}) Filter {
	tree := &fieldNodeTree{
		Key:        "",
		ParentNode: nil,
	}
	tree.ParseOmitValueWithCache("", omitScene, el)
	return Filter{
		node: tree,
	}
}

// EnableCache 决定是否启用缓存，默认开启（强烈建议，除非万一缓存模式下出现bug，可以关闭缓存退回曾经的无缓存过滤模式），开启缓存后会有30%-40%的性能提升，开启缓存并没有副作用，只是会让结构体的字段tag常驻内存减少tag字符串处理操作
func EnableCache(enable bool) {
	enableCache = false
}
