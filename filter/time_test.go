package filter

import (
	"fmt"
	"testing"
	"time"
)

type UserTime struct {
	//ID           int       `json:"id,select(list)"`
	ID    int    `json:"id,select(list)"`
	Title string `json:"title,select(list)"`

	BirthTime2   *Time     `json:"birth_time2,select(list)"`
	NilBirthTime Time      `json:"nil_birth_time,select(list)"`
	Timer        time.Time `json:"timer,select($any)"`
	BirthTime    Time      `json:"birth_time,select(list)"`
}

func TestTime(t *testing.T) {
	now := time.Now()
	user := UserTime{
		ID:         111,
		BirthTime:  Time(now),
		BirthTime2: (*Time)(&now),
		Timer:      time.Now(),
	}
	//marshal, err := json.Marshal(&user)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(string(marshal))
	fmt.Println(SelectMarshal("list", user).MustJSON())
}
