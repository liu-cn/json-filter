package filter

import (
	"encoding/json"
	"fmt"
)

type Users struct {
	UID          uint   `json:"uid,select($any)"`
	Name         string `json:"name,select(comment|chat|profile|justName)"`
	Age          int    `json:"age,select(comment|chat|profile)"`
	Avatar       string `json:"avatar,select(comment|chat|profile)"`
	Birthday     int    `json:"birthday,select(profile)"`
	Password     string `json:"password"`
	PasswordSlat string `json:"password_slat"`
	LangAge      []Lang `json:"langAge,select(profile|lookup|lang)"`
}

type Lang struct {
	Name string `json:"name,select(profile|lang)"`
	Arts []*Art `json:"arts,select(profile|lookup)"`
}

type Art struct {
	Name    string                 `json:"name,select(profile)"`
	Profile map[string]interface{} `json:"profile,select(profile|lookup)"`
	Values  []string               `json:"values,select(profile|lookup)"`
}

func (u Users) ArticleResp() interface{} {
	//这样当你后面想要优化性能时可以在这里进行优化，
	return SelectMarshal("article", u)
}

func (u Users) ProfileResp() interface{} {
	//这样当你后面想要优化性能时可以在这里进行优化，
	return SelectMarshal("profile", u)
}

func (u Users) ChatResp() interface{} {
	//假如性能出现瓶颈，想要优化
	chat := struct {
		UID  uint   `json:"uid"` //标记了$any无论选择任何场景都会解析该参数
		Name string `json:"name"`
	}{
		UID:  u.UID,
		Name: u.Name,
	}
	jsonStr, err := json.Marshal(chat)
	if err != nil {
		panic(err)
	}

	return string(jsonStr)
}

func newUsers() Users {
	return Users{
		UID:          1,
		Name:         "boyan",
		Age:          20,
		Avatar:       "https://www.avatar.com",
		Birthday:     2001,
		PasswordSlat: "slat",
		Password:     "123",
		LangAge: []Lang{
			{
				Name: "c",
				Arts: []*Art{
					{
						Name: "cc",
						Profile: map[string]interface{}{
							"c": "clang",
						},
						Values: []string{"1", "2"},
					},
				},
			},
			{
				Name: "c++",
				Arts: []*Art{
					{
						Name: "c++",
						Profile: map[string]interface{}{
							"c++": "cpp",
						},
						Values: []string{"cpp1", "cpp2"},
					},
				},
			},
			{
				Name: "Go",
				Arts: []*Art{
					{
						Name: "Golang",
						Profile: map[string]interface{}{
							"Golang": "go",
						},
						Values: []string{"Golang", "Golang1"},
					},
				},
			},
		},
	}
}

func filterUser() {
	users := newUsers()
	jsonStr, _ := json.Marshal(users)
	fmt.Println(string(jsonStr))
	//{"uid":1,"name":"boyan","age":20,"avatar":"https://www.avatar.com","birthday":2001,"password":"123","password_slat":"slat","langAge":[{"name":"c","arts":[{"name":"cc","profile":{"c":"clang"},"values":["1","2"]}]},{"name":"c++","arts":[{"name":"c++","profile":{"c++":"cpp"},"values":["cpp1","cpp2"]}]},{"name":"Go","arts":[{"name":"Golang","profile":{"Golang":"go"},"values":["Golang","Golang1"]}]}]}

	profile := SelectMarshal("profile", users)
	fmt.Println(profile) //|| fmt.Println(users.ProfileResp())

	//{"age":20,"avatar":"https://www.avatar.com","birthday":2001,"langAge":[{"arts":[{"name":"cc","profile":{"c":"clang"},"values":["1","2"]}],"name":"c"},{"arts":[{"name":"c++","profile":{"c++":"cpp"},"values":["cpp1","cpp2"]}],"name":"c++"},{"arts":[{"name":"Golang","profile":{"Golang":"go"},"values":["Golang","Golang1"]}],"name":"Go"}],"name":"boyan","uid":1}

	chat := SelectMarshal("chat", users)
	fmt.Println(chat) //||或者直接:fmt.Println(users.ChatResp())
	//{"age":20,"avatar":"https://www.avatar.com","name":"boyan","uid":1}

	unknown := SelectMarshal("unknown", users)
	fmt.Println(unknown)
	//{"uid":1}

	justName := SelectMarshal("justName", users)
	fmt.Println(justName)
	//{"name":"boyan","uid":1}

	lang := SelectMarshal("lang", users)
	fmt.Println(lang)
	//{"langAge":[{"name":"c"},{"name":"c++"},{"name":"Go"}],"uid":1}

	null := SelectMarshal("", users)
	fmt.Println(null)
	//{"uid":1}

	lookup := SelectMarshal("lookup", users)
	fmt.Println(lookup)
	//{"langAge":[{"arts":[{"profile":{"c":"clang"},"values":["1","2"]}]},{"arts":[{"profile":{"c++":"cpp"},"values":["cpp1","cpp2"]}]},{"arts":[{"profile":{"Golang":"go"},"values":["Golang","Golang1"]}]}],"uid":1}
}
