package main

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

func NewUser() Users {
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
