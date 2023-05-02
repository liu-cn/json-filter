package main

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
)

func TestTimes(t *testing.T) {
	v := filter.Select("Times", NewTimes())
	fmt.Println(v)

	v = filter.Omit("Times", NewTimes())
	fmt.Println(v)
}
