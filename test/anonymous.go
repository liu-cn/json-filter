package main

type AnonymousChild struct {
	PageInfo int `json:"pageInfo,select($any),omit(Anonymous)"`
	PageNum  int `json:"pageNum,select($any)"`
}

type Anonymous struct {
	Title          string                    `json:"title,select(article),omit(Anonymous)"`
	AnonymousChild `json:",select(article)"` // 这种tag字段名为空的方式会直接把该结构体展开，当作匿名结构体处理
	//AnonymousChild `json:"page,select(article)"` // 注意这里tag里标注了匿名结构体的字段名，所以解析时会解析成对象，不会展开
	Author string `json:"author,select(admin)"`
}

func NewAnonymous() Anonymous {
	anonymous := Anonymous{
		Author: "北洛",
		Title:  "c++从研发到脱发",
		AnonymousChild: AnonymousChild{
			PageInfo: 999,
			PageNum:  1,
		},
	}
	return anonymous
}
