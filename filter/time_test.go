package filter

import (
	"fmt"
	"testing"
	"time"
)

type UserTime struct {
	ID           int       `json:"id,select(list)"`
	BirthTime    Time      `json:"birth_time,select(list)"`
	BirthTime2   *Time     `json:"birth_time2,select(list)"`
	NilBirthTime Time      `json:"nil_birth_time,omitempty,select(list)"`
	Timer        time.Time `json:"timer,select($any)"`
}

func TestTime(t *testing.T) {
	now := time.Now()
	user := UserTime{
		ID:         111,
		BirthTime:  Time(now),
		BirthTime2: (*Time)(&now),
		Timer:      time.Now(),
	}

	fmt.Println(SelectMarshal("list", user).MustJSON())
}
