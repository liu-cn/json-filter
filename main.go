package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/liu-cn/json-filter/filter"
)

var timeType = reflect.TypeOf(time.Now())

type MyTime = time.Time

type Book struct {
	Page  string `json:"page,select(all)"`
	Title string `json:"title,select(all|justBooks|justBooks)"`
}
type User struct {
	Name  string    `json:"name,select(all)"`
	Time  time.Time `json:"time,select(all|time)"`
	MTime MyTime    `json:"m_time,select(all|time)"`
	Books []Book    `json:"books,select(all|justBooks)"`

	Map map[string]interface{} `json:"map,select(all|justMap)"`
}

func main() {
	user := User{
		Name: "boyan",
		Map:  make(map[string]interface{}),
		Time: time.Now(),
	}

	user.Map["book"] = Book{
		Page:  "10",
		Title: "golang",
	}
	user.Map["book-p"] = &Book{
		Page:  "10",
		Title: "golang",
	}

	user.Books = []Book{
		{
			Page:  "10",
			Title: "golang",
		},
		{
			Page:  "10",
			Title: "golang",
		},
		{
			Page:  "10",
			Title: "golang",
		},
	}

	user.Map["age"] = 10

	user.Map["null_map"] = make(map[string]interface{})

	fmt.Println(filter.SelectMarshal("time", user))

	fmt.Println(reflect.TypeOf(user.Time))
	fmt.Println(reflect.TypeOf(user.MTime))
	fmt.Println(reflect.TypeOf(user.Time) == timeType)
	fmt.Println(reflect.TypeOf(user.MTime) == timeType)
	//marshal, _ := json.Marshal(user.Books)
	//fmt.Println(string(marshal))
	//fmt.Println(filter.SelectMarshal("justMap", user))
	//fmt.Println(filter.SelectMarshal("justBooks", user))
	//fmt.Println(filter.SelectMarshal("justMap", user))
	//marshal, err := json.Marshal(user)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(string(marshal))

}
