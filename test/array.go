package main

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

type Tags struct {
	ID   uint   `json:"id,select(all),omit(id)"`
	Name string `json:"name,select(justName|all),omit(name)"`
	Icon string `json:"icon,select(chat|profile|all),omit(icon)"`
}

func newTags() []Tags {

	tags := []Tags{ //切片和数组都支持 slice or array
		{
			ID:   1,
			Name: "c",
			Icon: "icon-c",
		},
		{
			ID:   1,
			Name: "c++",
			Icon: "icon-c++",
		},
		{
			ID:   1,
			Name: "go",
			Icon: "icon-go",
		},
	}
	return tags
}
