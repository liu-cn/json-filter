package filter

type Filter struct {
	node *fieldNodeTree
}

func (f Filter) MarshalJSON() ([]byte, error) {
	return f.node.Bytes()
}
func (f Filter) MastMarshalJSON() []byte {
	return f.node.MustBytes()
}

func (f Filter) Interface() interface{} {
	return f.node.Marshal()
}

func (f Filter) MustJSON() string {
	return f.node.MustJSON()
}

func (f Filter) JSON() (string, error) {
	return f.node.JSON()
}

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
