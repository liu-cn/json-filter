package filter

import "github.com/gogf/gf/v2/os/gtime"

type GTime interface {
	Clone() *gtime.Time
}
