package filter

import (
	"reflect"
	"time"
)

var (
	timeTypes  = reflect.TypeOf(time.Now())
	byteTypes  = reflect.TypeOf([]byte{})
	emptySlice = make([]int, 0, 0)
)
