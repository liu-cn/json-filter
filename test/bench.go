package main

import (
	"fmt"
	"time"
)

func TimeAndRes(fn func() interface{}, runNum int) Info {
	now := time.Now()
	var res interface{}
	for i := 0; i < runNum; i++ {
		v := fn()
		if i == runNum-1 {
			res = v
		}
	}
	return Info{
		start:   now,
		end:     time.Now(),
		LastRes: res,
	}
}

type Info struct {
	start   time.Time
	end     time.Time
	LastRes interface{}
}

func (i *Info) Print(str ...interface{}) {
	s := ""
	for _, v := range str {
		s += fmt.Sprintf("%v", v)
	}

	fmt.Println(s, i.end.Sub(i.start))
}
