# json-filter
golang的json过滤器，随意选择字段，随意输出指定结构体的字段

0. [1分钟入门](#1分钟入门)
1. [过滤方式](#过滤方式)

​		[select() 选择器，选择过滤](#select() 选择器，选择过滤)

   	 [omit() 选择器， 排除过滤](omit() 选择器， 排除过滤)

2. [$any标识符 任意场景解析](#$any标识符 任意场景解析)

3. [过滤后的Filter结构体的方法](#过滤后的Filter结构体的方法)

​	     [Filter.Interface()](#Filter.Interface())  

   	  [Filter.MustJSON()](#Filter.MustJSON())

4. [高级用法](#高级用法)

​		  [对于map的过滤](#对于map的过滤)

​		  [对于切片/数组的直接过滤](#对于切片/数组的直接过滤)

​		  [匿名结构体的过滤](#匿名结构体的过滤)

​			    [展开结构体](#展开结构体)

​			    [不想展开结构体](#不想展开结构体)

​	    [极其复杂的场景，深层嵌套 结构体嵌套结构体，嵌套切片，嵌套map，嵌套指针，等等](#极其复杂的场景，深层嵌套 结构体嵌套结构体，嵌套切片，嵌套map，嵌套指针，等等)

5. [建议使用的姿势](#建议使用的姿势)
6. [在gin中返回json的demo](#在gin中返回json的demo)
7. [不想直接被解析为json字符串？](#不想直接被解析为json字符串？)









支持直接过滤以下数据结构

1. 结构体（包括匿名结构体）
2. map
3. 数组/切片
4. 复杂的结构体，嵌套指针嵌套数组嵌套map 嵌套匿名结构体等。



#### 1分钟入门













```go
//首先你需要引入下面的包：
github.com/liu-cn/json-filter/filter
```



```go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/liu-cn/json-filter/filter"
)

//同一个结构体，你可能想要在article 接口下只返回UID Avatar Nickname这三个字段就够了，其他字段不想要暴露
//另外你还想在profile 接口下返回 Nickname Sex VipEndTime Price 这四个字段，其他字段不想暴露
//这样的情况有很多，想要复用一个结构体来随心所欲的构造自己想要的json数据结构，可以看一个简单的demo

type User struct {
	UID    uint   `json:"uid,select(article)"`    //select中表示选中的场景(该字段将会使用到的场景)
	Avatar string `json:"avatar,select(article)"` //和上面一样此字段在article接口时才会解析该字段

	Nickname string `json:"nickname,select(article|profile)"` //"｜"表示有多个场景都需要这个字段 article接口需要 profile接口也需要

	Sex        int       `json:"sex,select(profile)"`          //该字段是仅仅profile才使用
	VipEndTime time.Time `json:"vip_end_time,select(profile)"` //同上
	Price      string    `json:"price,select(profile)"`        //同上
}

func main() {

	user := User{
		UID:        1,
		Nickname:   "boyan",
		Avatar:     "avatar",
		Sex:        1,
		VipEndTime: time.Now().Add(time.Hour * 24 * 365),
		Price:      "999.9",
	}

	marshal, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal)) //以下是官方的json解析输出结果：可以看到所有的字段都被解析了出来
	//{"uid":1,"nickname":"boyan","avatar":"avatar","sex":1,"vip_end_time":"2023-03-06T23:11:22.622693+08:00","price":"999.9"}

	fmt.Println(filter.SelectMarshal("article", user).MustJSON()) //以下是通过json-filter 过滤后的json，此输出是article接口下的json
	//{"avatar":"avatar","nickname":"boyan","uid":1}

	fmt.Println(filter.SelectMarshal("profile", user).MustJSON()) //profile接口下
	//{"nickname":"boyan","price":"999.9","sex":1,"vip_end_time":"2023-03-06T23:31:28.636529+08:00"}
}

```



##### 做这个的原因

不想暴露出其他不需要的字段，也不想重新建一个struct一个一个字段的赋值（懒）不过创建的model过多意味着维护的成本也会更高，
返回更多字段在不安全的同时意味着需要传输更多的数据，这就意味着会浪费带宽资源，编解码也更加耗时，
有时候我看到过有人为了偷懒把一个结构体序列化后返回，上面带着很多无用的字段，可能只有4-5个字段有用，
其他大多数字段都没有用，不仅影响阅读还浪费带宽，所以或许可以尝试用json-filter的过滤器来过滤你想要的字段吧，
不仅简单，更重要的是很强大，很复杂的结构体也可以过滤出你想要的字段。



#### 过滤方式

##### select() 选择器，选择过滤

select选择器快速入门已经说了，用法应该知道了select选择器标记的字段会被选择，omit则反之

```go
type User struct {
	UID    uint   `json:"uid,select(article)"`
	Avatar string `json:"avatar,select(article)"`
	Nickname string `json:"nickname,select(article|profile),omit(chat)"`
}
```

##### omit() 选择器， 排除过滤

omit则反之，标记的字段会被排除。

因为有时候一个场景，一个结构体很多字段就需要排斥2-3个字段，大部分字段都是需要的，这时候通过select选择就太过于啰嗦了。

比如上面你想要在私信接口(chat)下要UID Avatar 这两个字段，你就可以直接在 Nickname这个字段中把chat接口的场景排除掉就好了

这时候你需要在Nickname 的结构体tag里添加 omit(排除的场景|场景0|场景1|场景n)

这时候你需要调用

```go
f:=filter.OmitMarshal("chat",el) //这时Nickname字段就被排除掉了。
```

#### $any标识符 任意场景解析



你可能想要某个字段在任意场景下都被选中/或者任意场景都排除该字段，但是又不想在任何场景都写一遍。这样可以用$any 标识符干净整洁的完成你的需求。

```go
type User struct {
	UID    uint   `json:"uid,select($any)"`
  Password string `json:"password,omit($any)"`
}

//这样无论是任何场景
SelectMarshal("无论这里选择任何场景", user)//无论何种场景都会输出UID的字段。

OmitMarshal("无论这里选择任何场景",user)//无论何种场景都会排除password 字段。


```



#### 过滤后的Filter结构体的方法

##### Filter.Interface()  

过滤后的json数据结构，（可以被json编码）

```go
使用方法

f:=filter.SelectMarshal("场景",el) //el可以是 要过滤的 /结构体/map/切片/数组

会返回一个过滤后的Filter结构体，可以根据需要调用指定的方法来获取指定数据。

f.Interface()---->返回一个还没被解析的数据结构，此时可以直接使用官方的json.Marshal()来序列化成过滤后的json字符串。
ps：
  f:=filter.SelectMarshal("场景",el) 
  json.Marshal(f.Interface()) //等同于json.Marshal(filter.SelectMarshal("场景",el).Interface()) 
```



##### Filter.MustJSON()

过滤后直接编码成的json字符串

Must前缀的方法，不会返回err，如果使用过程中遇到err直接panic，测试使用无所谓，项目里使用一定要保证结构体正确无误。

```go
f:=filter.SelectMarshal("场景",要过滤的结构体/map/切片/数组) 

fmt.Println(f.MustJSON()) //---> 就不需要上面那样去json.Marshal 了，因为这样直接就是返回的json字符串

//如果想安全的使用这个方法请使用f.JSON() 会返回一个json字符串和err
	j, err := f.JSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(j)
```



#### 高级用法

##### 对于map的过滤

```go
m := map[string]interface{}{
		"name": "哈哈",
		"struct": User{     //此结构体还是用的1分钟入门里的User结构体
			UID:        1,
			Nickname:   "boyan",
			Avatar:     "avatar",
			Sex:        1,
			VipEndTime: time.Now().Add(time.Hour * 24 * 365),
			Price:      "999.9",
		},
	}

	fmt.Println(filter.SelectMarshal("article", m).MustJSON())
//{"name":"哈哈","struct":{"avatar":"avatar","nickname":"boyan","uid":1}}
//可以看到map也是可以直接过滤的。

```



##### 对于切片/数组的直接过滤

是完全支持对数组和切片的直接过滤的。

```go
func main() {
	type Tag struct {
		ID   uint   `json:"id,select(all)"`
		Name string `json:"name,select(justName|all)"`
		Icon string `json:"icon,select(chat|profile|all)"`
	}

	tags := []Tag{   //切片和数组都支持
		{
			ID:   1,
			Name: "c",
			Icon: "icon-c",
		},
		{
			ID:   1,
			Name: "c++",
			Icon: "icon-c++",
		},
		{
			ID:   1,
			Name: "go",
			Icon: "icon-go",
		},
	}

	fmt.Println(filter.SelectMarshal("justName", tags))
	//--->输出结果： [{"name":"c"},{"name":"c++"},{"name":"go"}]

	fmt.Println(filter.SelectMarshal("all", tags))
	//--->输出结果： [{"icon":"icon-c","id":1,"name":"c"},{"icon":"icon-c++","id":1,"name":"c++"},{"icon":"icon-go","id":1,"name":"go"}]

	fmt.Println(filter.SelectMarshal("chat", tags))
	//--->输出结果： [{"icon":"icon-c"},{"icon":"icon-c++"},{"icon":"icon-go"}]

}
```







##### 匿名结构体的过滤

对嵌入的匿名结构体也是完全支持的，并且支持的非常棒。

###### 展开结构体

```go
type Page struct {
	PageInfo int `json:"pageInfo,select($any)"`
	PageNum  int `json:"pageNum,select($any)"`
}

type Article struct {
	Title  string `json:"title,select(article)"`
	Page   `json:",select(article)"`     // 这种tag字段名为空的方式会直接把该结构体展开，当作匿名结构体处理  
  //Page `json:"page,select(article)"` // 注意这里tag里标注了匿名结构体的字段名，所以解析时会解析成对象，不会展开 
	Author string `json:"author,select(admin)"`
}

func main() {

	article := Article{
		Title: "c++从研发到脱发",
		Page: Page{
			PageInfo: 999,
			PageNum:  1,
		},
	}

	articleJson := filter.SelectMarshal("article", article)
	fmt.Println(articleJson)
  //输出结果--->  {"pageInfo":999,"pageNum":1,"title":"c++从研发到脱发"}
}

```

###### 不想展开结构体

```go
//不想把Page结构体展开也是可以的，很简单，把匿名结构体Page的结构体标签名加上
//接下来把
Page `json:",select(article)"`换成
Page   `json:"page,select(article)"`

//接下来看一下输出效果，可以看到字段没有被展开
{"page":{"pageInfo":999,"pageNum":1},"title":"c++从研发到脱发"}
```



##### 极其复杂的场景，深层嵌套 结构体嵌套结构体，嵌套切片，嵌套map，嵌套指针，等等

事实上json-filter可以做的事情绝对不止这么简单，嵌套各种复杂的数据类型，都是可以正确无误的解析的。

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
	return filter.SelectMarshal("article",u).Interface()
}

func (u User) ProfileResp() interface{} {
	//这样当你后面想要优化性能时可以在这里进行优化，
	return filter.SelectMarshal("profile",u).Interface()
}

func (u User) ChatResp() interface{} {
	//假如性能出现瓶颈，想要优化
	chat:= struct {
		UID    uint        `json:"uid"`
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



#### 在gin中返回json的demo

我看到很多人很喜欢用gin，那我就出一个简单的demo

```go
type User struct {
	UID    uint   `json:"uid,select(article|profile)"` //这个在select()里添加了article场景，
	Sex    int    `json:"sex,select(profile)"`         //不会被article场景解析，在article场景过滤时会直接忽略该字段。
	Avatar string `json:"avatar,select(article)"`      //会被article场景解析
}

func (u User) FilterProfile() interface{} {
	return filter.SelectMarshal("profile", u).Interface()
}

func main() {
	r := gin.New()
	r.GET("/user", GetUser)
	r.GET("/user/filter", GetUserFilter)
	log.Fatal(r.Run(":8080"))
}

func GetUserFilter(c *gin.Context) {
	user := User{
		UID:    1,
		Sex:    1,
		Avatar: "avatar",
	}
	c.JSON(200, user.FilterProfile())
	//{
	//  "sex": 1,
	//  "uid": 1
	//}
}
func GetUser(c *gin.Context) {
	user := User{
		UID:    1,
		Sex:    1,
		Avatar: "avatar",
	}

	c.JSON(200, user)
	//{
	//	"uid": 1,
	//	"sex": 1,
	//	"avatar": "avatar"
	//}
}

```

#### 不想直接被解析为json字符串？

你可能不希望直接解析成字符串，希望过滤后再挂在到其他结构体被解析，如果解析成字符串被挂载上去会被当成字符串解析，所以也是支持的。

```go
func OkWithData(data interface{}, c *gin.Context) {
	c.JSON(200, Response{
		Code: 0,
		Msg:  "ok",
		Data: data, //这个data应该是一个结构体或者map不应该是已经解析好的json字符串
	})
}

func UserRes(c *gin.Context) {
  user := User{
		UID:    1,
		Sex:    1,
		Avatar: "avatar",
	}
  
	OkWithData(filter.SelectMarshal("profile", user).Interface(), c)
}
```

