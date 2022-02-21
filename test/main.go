package main

import (
	"encoding/json"
	"fmt"
	"github.com/liu-cn/json-filter/filter"
)

func main() {
	example := filter.NewUsers()
	jsonStr, _ := json.Marshal(example)
	fmt.Println(string(jsonStr))
	//{"uid":1,"name":"boyan","age":20,"avatar":"https://www.avatar.com","birthday":2001,"password":"123","password_slat":"slat","langAge":[{"name":"c","arts":[{"name":"cc","profile":{"c":"clang"},"values":["1","2"]}]},{"name":"c++","arts":[{"name":"c++","profile":{"c++":"cpp"},"values":["cpp1","cpp2"]}]},{"name":"Go","arts":[{"name":"Golang","profile":{"Golang":"go"},"values":["Golang","Golang1"]}]}]}

	profile := filter.SelectMarshal("profile", example)
	fmt.Println(profile)
	//{"age":20,"avatar":"https://www.avatar.com","birthday":2001,"langAge":[{"arts":[{"name":"cc","profile":{"c":"clang"},"values":["1","2"]}],"name":"c"},{"arts":[{"name":"c++","profile":{"c++":"cpp"},"values":["cpp1","cpp2"]}],"name":"c++"},{"arts":[{"name":"Golang","profile":{"Golang":"go"},"values":["Golang","Golang1"]}],"name":"Go"}],"name":"boyan","uid":1}

	chat := filter.SelectMarshal("chat", example)
	fmt.Println(chat)
	//{"age":20,"avatar":"https://www.avatar.com","name":"boyan","uid":1}

	unknown := filter.SelectMarshal("unknown", example)
	fmt.Println(unknown)
	//{"uid":1}

	justName := filter.SelectMarshal("justName", example)
	fmt.Println(justName)
	//{"name":"boyan","uid":1}

	lang := filter.SelectMarshal("lang", example)
	fmt.Println(lang)
	//{"langAge":[{"name":"c"},{"name":"c++"},{"name":"Go"}],"uid":1}

	null := filter.SelectMarshal("", example)
	fmt.Println(null)
	//{"uid":1}

	lookup := filter.SelectMarshal("lookup", example)
	fmt.Println(lookup)
	//{"langAge":[{"arts":[{"profile":{"c":"clang"},"values":["1","2"]}]},{"arts":[{"profile":{"c++":"cpp"},"values":["cpp1","cpp2"]}]},{"arts":[{"profile":{"Golang":"go"},"values":["Golang","Golang1"]}]}],"uid":1}

}
