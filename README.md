# json-filter
golang的json过滤器，随意选择字段，随意输出指定结构体的字段，目前SelectMarshal已经可以使用。



很多时候你想要在不同场景复用这一个结构体，同时还想要在json做序列化的时候只输出你想要的字段，这样是可以做到的，你只需要在json的tag里添加select()标记，标签用法：select(场景1|场景2)，这样在json-filter解析并序列化的时候就会只序列化你想要的字段，

```go
//首先你需要引入下面的包：
github.com/liu-cn/json-filter/filter
```

filter.SelectMarshal("场景",要过滤的结构体)

#### 一个简单的demo

##### select()选择器

```go

//假如在文章(article)场景下，你只需要返回age，name，avatar，uid这四个字段就够了，
//不想暴露出其他不需要的字段，也不想重新建一个struct一个一个字段的赋值，
//返回更多字段在不安全的同时意味着需要传输更多的数据，这就意味着会浪费带宽资源，编解码也更加耗时，
//有时候我看到过有人为了偷懒把一个毫无意义的结构体序列化后返回，上面带着很多无用的字段，可能只有4-5个字段有用，
//其他大多数字段都没有用，不仅影响阅读还浪费带宽，所以或许可以尝试用json-filter的过滤器来过滤你想要的字段吧，
//不仅简单，更重要的是很强大，很复杂的结构体也可以过滤出你想要的字段。

import (
	"fmt"

	"github.com/liu-cn/json-filter/filter"
)

type User struct {
	UID      uint     `json:"uid,select(article|profile)"`  //这个在select()里添加了article场景，
	Name     string   `json:"name,select(article|profile)"` //会被article场景解析
	Age      int      `json:"age,select(article|profile)"`  //会被article场景解析
	Sex      int      `json:"sex,select(profile)"`          //不会被article场景解析，在article场景过滤时会直接忽略该字段。
	Avatar   string   `json:"avatar,select(article)"`       //会被article场景解析
	Password string   `json:"password"`
	Slat     string   `json:"-"` //任何场景都会被忽略
	Lang     []string `json:"lang,select(lang)"`
}

func main() {
	user := User{
		Name:     "boyan",
		UID:      1,
		Age:      20,
		Sex:      1,
		Avatar:   "https://www.avatar.com",
		Password: "pwd",
		Slat:     "slat",
		Lang:     []string{"c", "c++", "Go"},
	}

	articleJson := filter.SelectMarshal("article", user)
	fmt.Println(articleJson) //输出以下json：
	//{"age":20,"avatar":"https://www.avatar.com","name":"boyan","uid":1}

	//需求来了有一个需要展示编程语言的场景，要求只展示你的编程语言的字段，其他字段都不要展示，那怎么办呢？
	//很简单
	langJson := filter.SelectMarshal("lang", &user) //user传递指针和值都无所谓。
	fmt.Println(langJson)                           //输出以下json：
	//{"lang":["c","c++","Go"]}
}

```



#### $any，解析任意场景

```go

//假如有个字段无论任何场景都是需要用到，但是每个场景都写一遍又太臃肿，所以在select()里添加$any场景，表示任何场景都会携带该字段。
  //我想在任何场景的解析下都解析uid字段，需要这么写
 type User struct {
    UID    uint        `json:"uid,select($any)"` //标记了$any无论选择任何场景都会解析该参数
    Name   string      `json:"name,select(article|profile|chat)"`
    Avatar interface{} `json:"data,select(profile|chat)"`
	}
  
user := User{
		Name:   "boyan",
		Avatar: "avatar",
		UID:    1,
	}

articleJson := filter.SelectMarshal("article", user)
fmt.Println(articleJson)
{"name":"boyan","uid":1}
```



#### 复杂的场景，结构体嵌套结构体，嵌套切片，嵌套map，嵌套指针，等等。

事实上json-filter可以做的事情绝对不止这么简单，嵌套各种复杂的数据类型，都是可以正确的解析的。

来看一个相对复杂结构体的情况。

```go

这是一个相当复杂的结构体，
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


func main() {
  user:=NewUser()
  
  
  //我们先看一下原生json解析后的数据，把所有的字段都解析了出来。
  jsonStr, _ := json.Marshal(user)
 	fmt.Println(string(jsonStr))//
//{"uid":1,"name":"boyan","age":20,"avatar":"https://www.avatar.com","birthday":2001,"password":"123","password_slat":"slat","langAge":[{"name":"c","arts":[{"name":"cc","profile":{"c":"clang"},"values":["1","2"]}]},{"name":"c++","arts":[{"name":"c++","profile":{"c++":"cpp"},"values":["cpp1","cpp2"]}]},{"name":"Go","arts":[{"name":"Golang","profile":{"Golang":"go"},"values":["Golang","Golang1"]}]}]}

  
 // 如果我只想要编程语言相关加上部分用户信息的话
  lang := filter.SelectMarshal("lang", user)
	fmt.Println(lang)
	//{"langAge":[{"name":"c"},{"name":"c++"},{"name":"Go"}],"uid":1}
  
  //格式化后
  {
    "langAge":[
        {
            "name":"c"
        },
        {
            "name":"c++"
        },
        {
            "name":"Go"
        }
    ],
    "uid":1
	}
  
  //如果我只是想获取uid加上langAge下所有Art的部分字段信息， 你可以这样
 lookup := filter.SelectMarshal("lookup", user)
	fmt.Println(lookup)
	//{"langAge":[{"arts":[{"profile":{"c":"clang"},"values":["1","2"]}]},{"arts":[{"profile":{"c++":"cpp"},"values":["cpp1","cpp2"]}]},{"arts":[{"profile":{"Golang":"go"},"values":["Golang","Golang1"]}]}],"uid":1}
  
  
//格式化后，可以看到arts下name并没有展示，
  {
    "langAge":[
        {
            "arts":[
                {
                    "profile":{
                        "c":"clang"
                    },
                    "values":[
                        "1",
                        "2"
                    ]
                }
            ]
        },
        {
            "arts":[
                {
                    "profile":{
                        "c++":"cpp"
                    },
                    "values":[
                        "cpp1",
                        "cpp2"
                    ]
                }
            ]
        },
        {
            "arts":[
                {
                    "profile":{
                        "Golang":"go"
                    },
                    "values":[
                        "Golang",
                        "Golang1"
                    ]
                }
            ]
        }
    ],
    "uid":1
	}
}
//深层嵌套的数据结构也是可以正确无误的解析的，但是不建议数据结构太过于复杂，太复杂解析时会过于消耗性能，一般场景下解析的性能都是没有问题的，除非对性能要求很高的话可能需要你自己手动创建一个struct一个一个字段的赋值了。

```







#### 建议使用的姿势

```go


type User struct {
	UID    uint        `json:"uid,select($any)"` //标记了$any无论选择任何场景都会解析该参数
	Name   string      `json:"name,select(article|profile|chat)"`
	Avatar interface{} `json:"data,select(profile|chat)"`
}

func (u User) ArticleResp() interface{} {
	//这样当你后面想要优化性能时可以在这里进行优化，
	return filter.SelectMarshal("article",u)
}

func (u User) ProfileResp() interface{} {
	//这样当你后面想要优化性能时可以在这里进行优化，
	return filter.SelectMarshal("profile",u)
}

func (u User) ChatResp() interface{} {
	//假如性能出现瓶颈，想要优化
	chat:= struct {
		UID    uint        `json:"uid"` //标记了$any无论选择任何场景都会解析该参数
		Name   string  `json:"name"`
	}{
		UID: u.UID,
		Name: u.Name,
	}
	jsonStr, err := json.Marshal(chat)  //json-filter解析要比官方的json解析慢。
	if err!=nil {
		panic(err)
	}

	return string(jsonStr)
}

```

