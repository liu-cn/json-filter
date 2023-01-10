package main

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
)

type Anyway struct {
	Name string `json:"name,select($any),omit(user)"`
	Age  int    `json:"age,select(chat),omit(profile)"`
	Sex  int    `json:"sex,select(article),omit($any)"`
}

func newAnyway() Anyway {
	return Anyway{
		Age:  10,
		Sex:  10,
		Name: "boyan",
	}
}

func TestSelectAny(t *testing.T) {
	fmt.Println(filter.Select("chat", newAnyway()))
	//{"age":10,"name":"boyan"}
	fmt.Println(filter.Select("article", newAnyway()))
	//{"name":"boyan","sex":10}
}

func TestOmitAny(t *testing.T) {
	fmt.Println(filter.Omit("user", newAnyway()))
	//{"age":10}
	fmt.Println(filter.Omit("profile", newAnyway()))
	//{"name":"boyan"}
}
