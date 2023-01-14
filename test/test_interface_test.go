package main

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
)

type TestInterface struct {
	A interface{} `json:"a,select(a)"`
	B interface{} `json:"b,select(a)"`
	C interface{} `json:"c,select(a)"`
	D interface{} `json:"d,select(a)"`
}

func TestTestInterface(t *testing.T) {
	tt := TestInterface{
		A: "",
	}

	fmt.Println(mustJson(filter.Select("a", tt)))
}
