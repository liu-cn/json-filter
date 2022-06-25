package filter

type Filter struct {
	node *fieldNodeTree
}

func (f Filter) MarshalJSON() ([]byte, error) {
	return f.node.Bytes()
}

//Deprecated
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
		return "[Filter Err]"
	}
	return json
}

// SelectMarshal 第一个参数填你结构体select标签里的场景，第二个参数是你需要过滤的结构体对象，如果字段的select标签里标注的有该场景那么该字段会被选中。
func SelectMarshal(selectScene string, el interface{}) Filter {
	tree := &fieldNodeTree{
		Key:        "",
		ParentNode: nil,
	}
	tree.ParseSelectValue("", selectScene, el)
	return Filter{
		node: tree,
	}
}

// OmitMarshal 第一个参数填你结构体omit标签里的场景，第二个参数是你需要过滤的结构体对象，如果字段的omit标签里标注的有该场景那么该字段会被过滤掉
func OmitMarshal(omitScene string, el interface{}) Filter {
	tree := &fieldNodeTree{
		Key:        "",
		ParentNode: nil,
	}
	tree.ParseOmitValue("", omitScene, el)
	return Filter{
		node: tree,
	}
}
