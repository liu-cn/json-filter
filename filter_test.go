package main

import (
	"fmt"
	"testing"
)

type Book struct {
	Page  int    `json:"page,select(req|res|article)"`
	Price string `json:"price,select(req|res|article)"`
}

type User struct {
	Name  string `json:"name,select(req|res|article)"`
	Age   int    `json:"age,select(req|res|article)"`
	Hobby string `json:"hobby,select(req|res|article)"`
	Books []Book `json:"books,select(article)"`
	B     *Book  `json:"b,select(req|article)"`
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
		B: &Book{
			Price: "18.8",
			Page:  19,
		},
	}
	fmt.Println(SelectMarshal("req", &model))

	//=== RUN   TestFilter
	//{"B":{"page":19,"price":"18.8"},"age":20,"hobby":"coding","name":"boyan"}
	//--- PASS: TestFilter (0.00s)
	//PASS
	//_ = pkg.SelectMarshal("req", model)
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
		B: &Book{
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
