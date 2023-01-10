package filter

import (
	"encoding/json"
	"fmt"
)

func ExampleSelect() {
	type (
		Tag struct {
			Icon string `json:"icon,select(article)"`
			Name string `json:"name,select(profile)"`
		}
		User struct {
			Age  int     `json:"age,select(article|profile)"`
			ID   int     `json:"id,select($any)"`                        //$any表示任何场景都选中该字段
			Name *string `json:"name,omitempty,select(article|profile)"` //为nil忽略
			Tags []Tag   `json:"tags,select(article|profile)"`
		}
	)
	name := "小北"
	user := User{ID: 1, Name: &name, Age: 21, Tags: []Tag{{"icon", "foo"}, {"icon", "bar"}}}
	article := Select("article", &user) //传指针或值均可
	null := Select("null", user)
	user.Name = nil
	profile := Select("profile", user)
	articleJSON, _ := json.Marshal(article)
	fmt.Println(string(articleJSON))
	fmt.Println(profile) //可以直接打印，打印会直接输出过滤后的json
	fmt.Println(null)

	//Output:
	//{"id":1,"name":"小北","tags":[{"icon":"icon"},{"icon":"icon"}]}
	//{"age":21,"id":1,"name":"小北","tags":[{"name":"foo"},{"name":"bar"}]}
	//{"id":1}
}

func ExampleOmit() {

	type (
		Tag struct {
			Icon string `json:"icon,omit(article)"`
			Name string `json:"name,omit(profile)"`
		}
		User struct {
			Age      int   `json:"age"`
			Password int   `json:"password,omit($any)"` //$any表示任何场景都排除该字段
			Tags     []Tag `json:"tags"`
		}
	)
	user := User{Age: 21, Tags: []Tag{{"icon", "foo"}, {"icon", "bar"}}}
	article := Omit("article", &user) //传指针或值均可
	profile := Omit("profile", user)
	articleJSON, _ := json.Marshal(article)
	fmt.Println(string(articleJSON))
	fmt.Println(profile) //可以直接打印，打印会直接输出过滤后的json

	//Output:
	//{"age":21,"tags":[{"name":"foo"},{"name":"bar"}]}
	//{"age":21,"tags":[{"icon":"icon"},{"icon":"icon"}]}
}
