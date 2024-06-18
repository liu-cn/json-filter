package main

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/liu-cn/json-filter/filter"
	"testing"
	"time"
)

type GTime struct {
	Create *gtime.Time `json:"create,select(timeTest)"`
	Test   string      `json:"test,select(timeTest)"`
}

func TestGTime(t *testing.T) {

	gt := GTime{
		Create: gtime.New(time.Now()),
		Test:   "test",
	}
	fmt.Println(filter.Select("timeTest", &gt))
}
