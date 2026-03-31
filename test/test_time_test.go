package main

import (
	"fmt"
	"github.com/liu-cn/json-filter/filter"
	"testing"
	"time"
)

type JSONTime struct {
	time.Time
}

func (t *JSONTime) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}
	return []byte(`"` + t.UTC().Format(time.RFC3339) + `"`), nil
}

type GTime struct {
	Create *JSONTime `json:"create,select(timeTest)"`
	Test   string    `json:"test,select(timeTest)"`
}

func TestGTime(t *testing.T) {
	now := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)

	gt := GTime{
		Create: &JSONTime{Time: now},
		Test:   "test",
	}
	got := mustJson(filter.Select("timeTest", &gt))
	want := `{"create":"2024-01-02T03:04:05Z","test":"test"}`
	if got != want {
		t.Fatalf("unexpected json: %s", got)
	}
	fmt.Println(got)
}
