package filter

import (
	"fmt"
	"testing"
)

type User struct {
	Name string `json:"name,select(justName|req|foo)"`
	Age  int    `json:"select(req|res|article)"`

	LongName string `json:"long_name,select(foo)"`
	Hobby    string `json:"hobby,select(req|res|foo)"`
	Books    []Book `json:"books,select()"`
	Book     *Book  `json:"book,select(res|foo)"`
}

type Book struct {
	Page  int    `json:"page,select(req|foo)"`
	Price string `json:"price,select(res|foo)"`
	Title string `json:"title"`
}

func TestFilter(t *testing.T) {
	model := User{
		Name:  "boyan",
		Age:   20,
		Hobby: "coding",
		Books: []Book{
			{Page: 10, Price: "199.9"},
			{Page: 100, Price: "1999.9"},
		},
		LongName: "long name",
		Book: &Book{
			Price: "18.8",
			Page:  19,
			Title: "c++从研发到脱发",
		},
	}
	fmt.Println(SelectMarshal("res", &model))
	//---->>output 输出结果： {"Age":20,"book":{"price":"18.8"},"hobby":"coding"}

	fmt.Println(SelectMarshal("justName", &model))
	//---->>output 输出结果： {"name":"boyan"}

	fmt.Println(SelectMarshal("foo", &model))
	//---->>output 输出结果： {"book":{"page":19,"price":"18.8"},"hobby":"coding","long_name":"long name","name":"boyan"}
}

func BenchmarkFilter(b *testing.B) {
	model := User{
		Name:  "boyan",
		Age:   20,
		Hobby: "coding",
		Books: []Book{
			{Page: 10, Price: "199.9"},
			{Page: 100, Price: "1999.9"},
		},
		Book: &Book{
			Price: "18.8",
			Page:  19,
		},
	}
	for i := 0; i < b.N; i++ {
		_ = SelectMarshal("req", model)
	}

	//goos: darwin
	//goarch: amd64
	//pkg: filter
	//cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
	//BenchmarkFilter
	//BenchmarkFilter-16    	  176220	      6421 ns/op
	//PASS

}
