package filter

import (
	"encoding/json"
	"fmt"
)

func ExampleOmit() {

	//简单案例
	type Tag struct {
		Icon string `json:"icon,omit(article)"`
		Name string `json:"name,omit(profile)"` //profile场景下会被排除
	}
	profile := Omit("profile", &Tag{Icon: "foo", Name: "bar"}) //表示排除omit中带profile的字段
	profileJSON, _ := json.Marshal(profile)
	fmt.Println(string(profileJSON)) //可以直接fmt.Println(profileJSON)来打印
	//Output:
	//{"icon":"foo"}

	//复杂案例
	type (
		Articles struct {
			Password int   `json:"password,omit($any)"` //$any表示任何场景都排除该字段
			Tags     []Tag `json:"tags"`
			Hot      int   `json:"hot,select(img),func(GetHot)"` //热度 过滤时会调用GetHot方法获取该字段的值，用法详解请看: https://github.com/liu-cn/json-filter/blob/main/example/func_test.go
			Like     int   `json:"-"`
			Collect  int   `json:"-"`
		}
	)

	//func (a Articles) GetHot() int { //这个方法里可以对字段进行处理，处理后可以返回一个任意值
	//	return a.Like + a.Collect
	//}

	articles := Articles{Like: 100, Collect: 20, Tags: []Tag{{"icon", "foo"}, {"icon", "bar"}}}
	article := Omit("article", &articles) //尽量传指针，不传指针func选择器中的用指针接收的方法无法被调用
	profile = Omit("profile", &articles)
	articleJSON, _ := json.Marshal(article)
	fmt.Println(string(articleJSON))
	fmt.Println(profile) //可以直接打印，打印会直接输出过滤后的json

	//Output:
	//{"hot":120,"tags":[{"name":"foo"},{"name":"bar"}]}
	//{"hot":120,"tags":[{"icon":"icon"},{"icon":"icon"}]}
}

func ExampleSelect() {

	//简单案例
	type Tag struct {
		Icon string `json:"icon,select(article)"`
		Name string `json:"name,select(profile)"` //profile场景下会展示
	}
	profile := Select("profile", &Tag{Icon: "foo", Name: "bar"})
	fmt.Println(profile)
	//Output:
	//{"name":"bar"}

	//复杂案例
	type User struct {
		Age  int     `json:"age,select(article|profile)"`
		ID   int     `json:"id,select($any)"`                        //$any表示任何场景都选中该字段
		Name *string `json:"name,omitempty,select(article|profile)"` //为nil忽略
		Tags []Tag   `json:"tags,select(article|profile)"`
	}

	name := "小北"
	user := User{ID: 1, Name: &name, Age: 21, Tags: []Tag{{"icon", "foo"}, {"icon", "bar"}}}
	article := Select("article", &user) //尽量传指针
	null := Select("null", &user)
	//user.Name = nil
	profile = Select("profile", &user)
	articleJSON, _ := json.Marshal(article)
	fmt.Println(string(articleJSON))
	fmt.Println(profile) //可以直接打印，打印会直接输出过滤后的json
	fmt.Println(null)

	//Output:
	//{"age":21,"id":1,"name":"小北","tags":[{"icon":"icon"},{"icon":"icon"}]}
	//{"age":21,"id":1,"name":"小北","tags":[{"name":"foo"},{"name":"bar"}]}
	//{"id":1}
}
