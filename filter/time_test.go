package filter

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

type UserTime struct {
	ID         int       `json:"id,select(list)"`
	BirthTime  Time      `json:"birth_time,select(list)"`
	BirthTime2 *Time     `json:"birth_time2,select(list)"`
	Timer      time.Time `json:"timer,select($any)"`
}

func TestTime(t *testing.T) {
	now := time.Now()
	user := UserTime{
		ID:         111,
		BirthTime:  Time(now),
		BirthTime2: (*Time)(&now),
		Timer:      time.Now(),
	}
	v := reflect.ValueOf(Time{})
	val, ok := v.Interface().(json.Marshaler)
	fmt.Println(val, ok)

	fmt.Println(SelectMarshal("list", user).MustJSON())

	u2 := UserTime{}
	marshal, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(marshal, &u2)
	if err != nil {
		panic(err)
	}

	fmt.Println(u2)
	fmt.Println(u2.BirthTime)
}
