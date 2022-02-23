package main

import (
	"encoding/json"
	"fmt"

	"github.com/liu-cn/json-filter/filter"
)

type Book struct {
	Page  string `json:"page,select(all)"`
	Title string `json:"title,select(all|justBooks|justBooks)"`
}
type User struct {
	Name string `json:"name,select(all)"`

	Books []Book `json:"books,select(all|justBooks)"`

	Map map[string]interface{} `json:"map,select(all|justMap)"`
}

func main() {
	user := User{
		Name: "boyan",
		Map:  make(map[string]interface{}),
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

	fmt.Println(filter.SelectMarshal("justBooks", user.Books))
	marshal, _ := json.Marshal(user.Books)
	fmt.Println(string(marshal))
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
