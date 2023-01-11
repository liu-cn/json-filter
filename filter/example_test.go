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
	article := Select("article", &user) //尽量传指针
	null := Select("null", &user)
	//user.Name = nil
	profile := Select("profile", &user)
	articleJSON, _ := json.Marshal(article)
	fmt.Println(string(articleJSON))
	fmt.Println(profile) //可以直接打印，打印会直接输出过滤后的json
	fmt.Println(null)

	//Output:
	//{"age":21,"id":1,"name":"小北","tags":[{"icon":"icon"},{"icon":"icon"}]}
	//{"age":21,"id":1,"name":"小北","tags":[{"name":"foo"},{"name":"bar"}]}
	//{"id":1}
}

func (a *Article) GetHot() {

}

type (
	ExampleOmitTag struct {
		Icon string `json:"icon,omit(article)"`
		Name string `json:"name,omit(profile)"`
	}
	ExampleOmitArticles struct {
		Password int              `json:"password,omit($any)"` //$any表示任何场景都排除该字段
		Tags     []ExampleOmitTag `json:"tags"`
		Hot      int              `json:"hot,select(img),func(GetHot)"` //热度 过滤时会调用GetHot方法获取该字段的值
		Like     int              `json:"-"`
		Collect  int              `json:"-"`
	}
)

func (a ExampleOmitArticles) GetHot() int { //这个方法里可以对字段进行处理，处理后可以返回一个任意值
	return a.Like + a.Collect
}

func ExampleOmit() {

	articles := ExampleOmitArticles{Like: 100, Collect: 20, Tags: []ExampleOmitTag{{"icon", "foo"}, {"icon", "bar"}}}
	article := Omit("article", &articles) //尽量传指针，不传指针func选择器中的用指针接收的方法无法被调用
	profile := Omit("profile", &articles)
	articleJSON, _ := json.Marshal(article)
	fmt.Println(string(articleJSON))
	fmt.Println(profile) //可以直接打印，打印会直接输出过滤后的json

	//Output:
	//{"hot":120,"tags":[{"name":"foo"},{"name":"bar"}]}
	//{"hot":120,"tags":[{"icon":"icon"},{"icon":"icon"}]}
}

type InterfaceTest struct {
	Data interface{} `json:"data"`
}

func ExampleFilter_Interface() {

	fmt.Println(Omit("", InterfaceTest{Data: map[string]interface{}{
		"1": 1,
	}}))
	fmt.Println(Omit("", InterfaceTest{Data: map[string]interface{}{
		"1": 1,
		"2": map[string]interface{}{
			"3": 3,
		},
	}}))
	//Output:
	//{"data":{"1":1}}
	//{"data":{"1":1,"2":{"3":3}}}

}
