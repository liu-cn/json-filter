package main

import (
	"github.com/liu-cn/json-filter/filter"
	"time"
)

type Times struct {
	NilBirthTime filter.Time  `json:"nil_birth_time,select(Times)"`
	BirthTime2   *filter.Time `json:"birth_time2,select(Times)"`
	BirthTime    filter.Time  `json:"birth_time,select(Times),omit(Times)"`
	Timer        time.Time    `json:"timer,select($any)"`
}

func NewTimes() Times {
	//layout := "2006-01-02 15:04:05"
	t := time.Time{}

	date := time.Date(2016, 1, 2, 15, 4, 5, 0, time.UTC)
	BirthTime2 := filter.Time(date)
	times := Times{
		NilBirthTime: filter.Time(t),
		BirthTime2:   &BirthTime2,
		BirthTime:    filter.Time(date),
		Timer:        date,
	}
	return times
}
