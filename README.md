# json-filter

切换语言：

​        [简体中文](#简体中文)

​        [English](#English)

<img title="" src="https://github.com/liu-cn/json-filter/blob/main/logo.png" alt="" data-align="center">

## English

# future

    1.重写缓存策略（已实现但是还没完全测试，性能有很大提升，在filter_field_cache 分支下）
    2.内置json编码功能，消除对官方json编码工具的依赖（同时还可以提升性能），在过滤的同时编码。
    3.提升反射性能

Golang's JSON field filter can select fields at will, output fields of specified structures at will, and reuse structures.  **It fully supports generics** and is Compatible with all versions of go,The official json can support and be compatible with all json-filter

list：

[Learn it in 1 minute](#Learn it in 1 minute)

[Select filter](#Select filter)

[Omit selector exclusion filter](#Omit selector exclusion filter)

[$any will select this field in any case](#$any will select this field in any case)

[Method of filtering filter structure](#Method of filtering filter structure)

[Advanced Usage](#Advanced Usage)

[Extremely complex scene deep nesting - structure nesting - structure nesting - slice nesting - map nesting - pointer nesting - and so on](#Extremely complex scene deep nesting - structure nesting - structure nesting - slice nesting - map nesting - pointer nesting - and so on)

[Recommended posture](#Recommended posture)

[Return the demo of JSON in gin](#Return the demo of JSON in gin)

[Don't want to be parsed directly into JSON strings?](#Don't want to be parsed directly into JSON strings?)

Support direct filtering of the following data structures

1. struct（Include anonymous structures）

2. map

3. array/slice

4. pointer nesting

5. Omitempty zero value ignored

#### Learn it in 1 minute

```go
//First, you need to introduce the following package：
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

//For the same structure, you may want to return only uid and avatar nickname fields under the article interface. Other fields do not want to be exposed

//In addition, you also want to return the four fields of nickname sex vipendtime price under the profile interface. Other fields do not want to be exposed

//There are many such situations. If you want to reuse a structure to construct the JSON data structure you want, you can see a simple demo

type User struct {
    UID    uint   `json:"uid,select(article)"`    //elect indicates the selected scene (the case that this field will use)
    Avatar string `json:"avatar,select(article)"` //As above, this field will only be resolved when the article interface

    Nickname string `json:"nickname,select(article|profile)"` //"｜"It means that this field is required for multiple cases. The article interface and the profile interface are also required

    Sex        int       `json:"sex,select(profile)"`          //This field is only used by profile
    VipEndTime time.Time `json:"vip_end_time,select(profile)"` //ditto
    Price      string    `json:"price,select(profile)"`        //ditto
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
    fmt.Println(string(marshal)) //The following is the official JSON parsing output: you can see that all fields have been parsed
    //{"uid":1,"nickname":"boyan","avatar":"avatar","sex":1,"vip_end_time":"2023-03-06T23:11:22.622693+08:00","price":"999.9"}


  //usage：filter.Select("select case",This can be：slice/array/struct/pointer/map)
    fmt.Println(filter.Select("article", user)) //The following is the JSON filtered by JSON filter. This output is the JSON under the article interface
    //{"avatar":"avatar","nickname":"boyan","uid":1}

    fmt.Println(filter.Select("profile", user)) //profile result
    //{"nickname":"boyan","price":"999.9","sex":1,"vip_end_time":"2023-03-06T23:31:28.636529+08:00"}
}
```

##### The reason for doing this

I don't want to expose other unnecessary fields, and I don't want to rebuild a struct and assign values to fields one by one (lazy). However, too many models mean that the maintenance cost will be higher,
Returning more fields is unsafe and means that more data needs to be transmitted, which means that bandwidth resources will be wasted and encoding and decoding will be more time-consuming,
Sometimes I've seen someone serialize a structure and return it in order to be lazy, with many useless fields on it. Maybe only 4-5 fields are useful,
Most other fields are useless, which not only affects reading but also wastes bandwidth, so maybe you can try to filter the fields you want with JSON filter,
Not only simple, but more importantly, very powerful and complex structures can also filter out the fields you want.

#### Filtering mode

##### Select filter

As mentioned in the quick start of select selector, you should know the usage. The fields marked by select selector will be selected, and omit will be the opposite

```go
type User struct {
    UID    uint   `json:"uid,select(article)"`
    Avatar string `json:"avatar,select(article)"`
    Nickname string `json:"nickname,select(article|profile),omit(chat)"`
}
```

##### Omit selector exclusion filter

Omit, on the contrary, the marked fields will be excluded.
Because sometimes in a scene, many fields in a structure need to exclude 2-3 fields, and most fields are required. At this time, selecting through select is too verbose.
For example, if you want to use the fields uid and avatar under the private message interface (ChAT), you can directly exclude the scene of the chat interface from the field nickel
At this time, you need to add omit (excluded scene | scene 0 | Scene 1 | scene n) to the structure tag of nickname
At this time, you need to call

```go
f:=filter.Omit("chat",el) //The nickname field is then excluded.
```

##### Omitempty zero value ignored

Zero value ignoring is supported, and the omitempty must be written after the label name of the structure field

```go
Nickname *string `json:"nickname,omitempty,select(article|profile)"`   //nil Omitempty
Nickname string `json:"nickname,omitempty,select(article|profile)"`   //“” Omitempty
Age int `json:"age,omitempty,select(article|profile)"` //0 Omitempty

//Empty structures can also be ignored
```

#### $any will select this field in any case

You may want a field to be selected in any scenario and / or excluded from any scenario, but you don't want to write it again in any scenario. In this way, you can use the $any identifier to complete your requirements cleanly and neatly.

```go
type User struct {
    UID    uint   `json:"uid,select($any)"`
  Password string `json:"password,omit($any)"`
}

//In this way, no matter any case
Select("Whatever case you choose here", user)//No matter what kind of case, the field of uid will be output.

Omit("Whatever case you choose here",user)//The password field is excluded in any scenario.
```

#### Method of filtering filter structure

fmt.Println(f) //---> You don't need to go to JSON like that Marshall, because this is the returned JSON string directly

//If you want to use this method safely, use F. JSON () to return a JSON string and err
    j, err := f.JSON()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(j)

```
#### Advanced Usage

##### Filtering for map

```go
m := map[string]interface{}{
        "name": "哈哈",
        "struct": User{     //This structure is also the first user structure used
            UID:        1,
            Nickname:   "boyan",
            Avatar:     "avatar",
            Sex:        1,
            VipEndTime: time.Now().Add(time.Hour * 24 * 365),
            Price:      "999.9",
        },
    }

    fmt.Println(filter.Select("article", m))
//{"name":"哈哈","struct":{"avatar":"avatar","nickname":"boyan","uid":1}}
//You can see that the map can also be filtered directly.
```

##### Direct filtering for slices / arrays

It fully supports the direct filtering of arrays and slices.

```go
func main() {
    type Tag struct {
        ID   uint   `json:"id,select(all)"`
        Name string `json:"name,select(justName|all)"`
        Icon string `json:"icon,select(chat|profile|all)"`
    }

    tags := []Tag{   //Both slices and arrays are supported
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

    fmt.Println(filter.Select("justName", tags))
    //--->output： [{"name":"c"},{"name":"c++"},{"name":"go"}]

    fmt.Println(filter.Select("all", tags))
    //--->output： [{"icon":"icon-c","id":1,"name":"c"},{"icon":"icon-c++","id":1,"name":"c++"},{"icon":"icon-go","id":1,"name":"go"}]

    fmt.Println(filter.Select("chat", tags))
    //--->output： [{"icon":"icon-c"},{"icon":"icon-c++"},{"icon":"icon-go"}]

}
```

##### Filtering anonymous structures

The embedded anonymous structure is also fully supported, and the support is very good.
Expand structure

```go
type Page struct {
    PageInfo int `json:"pageInfo,select($any)"`
    PageNum  int `json:"pageNum,select($any)"`
}

type Article struct {
    Title  string `json:"title,select(article)"`
    Page   `json:",select(article)"`     // In this way, if the tag field name is empty, the structure will be expanded directly and treated as an anonymous structure
  //Page `json:"page,select(article)"` // Note that the tag here is marked with the field name of the anonymous structure, so it will be resolved into an object during parsing and will not be expanded
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

    articleJson := filter.Select("article", article)
    fmt.Println(articleJson)
  //output--->  {"pageInfo":999,"pageNum":1,"title":"c++从研发到脱发"}
}
```

###### Do not want to expand the structure

```go
//It's also possible not to expand the page structure. It's very simple to add the structure label signature of the anonymous Structure page
//Next put
Page `json:",select(article)"`change into
Page   `json:"page,select(article)"`

//Next, take a look at the output effect. You can see that the field is not expanded
{"page":{"pageInfo":999,"pageNum":1},"title":"c++从研发到脱发"}
```

##### Extremely complex scene deep nesting - structure nesting - structure nesting - slice nesting - map nesting - pointer nesting - and so on

In fact, JSON filter can do more than that. Nested various complex data types can be parsed correctly.
Let's look at a relatively complex structure.

```go
This is a rather complex structure,
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


  //Let's take a look at the data parsed by the native JSON and parse all the fields.
  jsonStr, _ := json.Marshal(user)
     fmt.Println(string(jsonStr))//
//{"uid":1,"name":"boyan","age":20,"avatar":"https://www.avatar.com","birthday":2001,"password":"123","password_slat":"slat","langAge":[{"name":"c","arts":[{"name":"cc","profile":{"c":"clang"},"values":["1","2"]}]},{"name":"c++","arts":[{"name":"c++","profile":{"c++":"cpp"},"values":["cpp1","cpp2"]}]},{"name":"Go","arts":[{"name":"Golang","profile":{"Golang":"go"},"values":["Golang","Golang1"]}]}]}


 // If I only want to add some user information related to the programming language
  lang := filter.Select("lang", user)
    fmt.Println(lang)
    //{"langAge":[{"name":"c"},{"name":"c++"},{"name":"Go"}],"uid":1}

  //format
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

  //If I just want to get some field information of uid and all arts under langage, you can do this
 lookup := filter.Select("lookup", user)
    fmt.Println(lookup)
    //{"langAge":[{"arts":[{"profile":{"c":"clang"},"values":["1","2"]}]},{"arts":[{"profile":{"c++":"cpp"},"values":["cpp1","cpp2"]}]},{"arts":[{"profile":{"Golang":"go"},"values":["Golang","Golang1"]}]}],"uid":1}


//After formatting, you can see that the name under arts is not displayed,
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
//Deep nested data structures can also be parsed correctly, but it is not recommended that the data structure is too complex, which will consume too much performance. In general, there is no problem with the parsing performance in scenarios. Unless you have high performance requirements, you may need to manually create a struct and assign values to fields one by one.
```

#### Recommended posture

```go
type User struct {
    UID    uint        `json:"uid,select($any)"` //Marked with $any, this parameter will be resolved no matter which case is selected
    Name   string      `json:"name,select(article|profile|chat)"`
    Avatar interface{} `json:"data,select(profile|chat)"`
}

func (u User) ArticleResp() interface{} {
    //In this way, when you want to optimize the performance later, you can optimize it here,
    return filter.Select("article",u)
}

func (u User) ProfileResp() interface{} {
    //In this way, when you want to optimize the performance later, you can optimize it here,
    return filter.Select("profile",u)
}

func (u User) ChatResp() interface{} {
    //If there is a performance bottleneck, you want to optimize it
    chat:= struct {
        UID    uint        `json:"uid"`
        Name   string  `json:"name"`
    }{
        UID: u.UID,
        Name: u.Name,
    }
    return chat
}
```

#### Return the demo of JSON in gin

I see many people like to use gin, so I'll make a simple demo

```go
type User struct {
    UID    uint   `json:"uid,select(article|profile)"` //This adds an article case to select()
    Sex    int    `json:"sex,select(profile)"`         //It will not be resolved by the article case, and this field will be ignored directly when filtering the article case.
    Avatar string `json:"avatar,select(article)"`      //Will be resolved by the article case
}

func (u User) FilterProfile() interface{} {
    return filter.Select("profile", u)
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
    //    "uid": 1,
    //    "sex": 1,
    //    "avatar": "avatar"
    //}
}
```

#### Don't want to be parsed directly into JSON strings?

You may not want to directly parse into a string. You want to filter it and then hang it to other structures for parsing. If it is parsed into a string and mounted, it will be parsed as a string, so it is also supported.

```go
func OkWithData(data interface{}, c *gin.Context) {
    c.JSON(200, Response{
        Code: 0,
        Msg:  "ok",
        Data: data, //The data should be a structure or the map should not be a JSON string that has been parsed
    })
}

func UserRes(c *gin.Context) {
  user := User{
        UID:    1,
        Sex:    1,
        Avatar: "avatar",
    }

    OkWithData(filter.Select("profile", user), c)
}
```

### 简体中文

golang的json字段过滤器，随意选择字段，随意输出指定结构体的字段，复用结构体，**全面支持泛型**，**对于go所有版本均完美兼容**，官方json库能做的json-filter全部兼容和支持。

视频教程快速入门：https://www.bilibili.com/video/BV1s14y1G72v/

**代码有很多值得优化和改进的地方，非常欢迎大家一起参与贡献代码，优化代码，贡献文档，提意见，后面会持续优化，希望能变得更强，开源项目全靠用爱发电。**

大纲：

[1分钟入门](#1分钟入门)

[select选择器选择过滤](#select选择器选择过滤)

[omit选择器排除过滤](#omit选择器排除过滤)

[func选择器自定义方法进行字段处理](#func选择器自定义方法进行字段处理)

[$any标识符任意场景解析](#$any标识符任意场景解析)

[过滤后的Filter结构体的方法](#过滤后的Filter结构体的方法)

[高级用法-对于map的过滤-切片/数组的直接过滤-匿名结构体的过滤](#高级用法)

[极其复杂的场景深层嵌套-结构体嵌套结构体-嵌套切片-嵌套map-嵌套指针-等等](#极其复杂的场景深层嵌套-结构体嵌套结构体-嵌套切片-嵌套map-嵌套指针-等等)

[建议使用的姿势](#建议使用的姿势)

[在gin中返回json的demo](#在gin中返回json的demo)

[不想直接被解析为json字符串？](#不想直接被解析为json字符串？)

支持直接过滤以下数据结构

1. 结构体（包括匿名结构体）
2. map
3. 数组/切片
4. 复杂的结构体，嵌套指针嵌套数组嵌套map 嵌套匿名结构体等。
5. 支持为空忽略，

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

    //用法：filter.Select("select里的一个场景",这里可以是slice/array/struct/pointer/map)
    article:=filter.Select("article", user)
    articleBytes, _ := json.Marshal(article)
    fmt.Println(string(articleBytes)) //以下是通过json-filter 过滤后的json，此输出是article接口下的json
    //{"avatar":"avatar","nickname":"boyan","uid":1}

  //filter.Select fmt打印的时候会自动打印过滤后的json字符串
    fmt.Println(filter.Select("article", user)) //以下是通过json-filter 过滤后的json，此输出是article接口下的json
    //{"avatar":"avatar","nickname":"boyan","uid":1}

    fmt.Println(filter.Select("profile", user)) //profile接口下
    //{"nickname":"boyan","price":"999.9","sex":1,"vip_end_time":"2023-03-06T23:31:28.636529+08:00"}
}
```

**注意！！下面还有更高级的过滤方式，建议把后面的文档看完。**

##### 做这个的原因

不想暴露出其他不需要的字段，也不想重新建一个struct一个一个字段的赋值（懒）不过创建的model过多意味着维护的成本也会更高，
返回更多字段在不安全的同时意味着需要传输更多的数据，这就意味着会浪费带宽资源，编解码也更加耗时，
有时候我看到过有人为了偷懒把一个结构体序列化后返回，上面带着很多无用的字段，可能只有4-5个字段有用，
其他大多数字段都没有用，不仅影响阅读还浪费带宽，所以或许可以尝试用json-filter的过滤器来过滤你想要的字段吧，
不仅简单，更重要的是很强大，很复杂的结构体也可以过滤出你想要的字段。

#### 过滤方式

##### select选择器选择过滤

select选择器快速入门已经说了，用法应该知道了select选择器标记的字段会被选择，omit则反之

```go
type User struct {
    UID    uint   `json:"uid,select(article)"`
    Avatar string `json:"avatar,select(article)"`
    Nickname string `json:"nickname,select(article|profile),omit(chat)"`
}
```

##### omit选择器排除过滤

omit则反之，标记的字段会被排除。

因为有时候一个场景，一个结构体很多字段就需要排斥2-3个字段，大部分字段都是需要的，这时候通过select选择就太过于啰嗦了。

比如上面你想要在私信接口(chat)下要UID Avatar 这两个字段，你就可以直接在 Nickname这个字段中把chat接口的场景排除掉就好了

这时候你需要在Nickname 的结构体tag里添加 omit(排除的场景|场景0|场景1|场景n)

这时候你需要调用

```go
f:=filter.Omit("chat",el) //这时Nickname字段就被排除掉了。
```

##### omitempty零值忽略

支持零值忽略，omitempty必须写在结构体字段标签名字后面

```go
Nickname *string `json:"nickname,omitempty,select(article|profile)"`   //为nil忽略
Nickname string `json:"nickname,omitempty,select(article|profile)"`   //为“”忽略
Age int `json:"age,omitempty,select(article|profile)"` //为0忽略

//空结构体也可以忽略
```

### func选择器自定义方法进行字段处理

```go
type Image struct {
    Url     []byte `json:"url,select(img),func(GetUrl)"`
    Path    string `json:"path,select(img),func(GetImagePath)"`
    Name    string `json:"name"`
    Hot     int    `json:"hot,select(img),func(GetHot)"` //热度
    Like    int
    Collect int
    Forward int
}

func (i Image) GetUrl() string {
    return string(i.Url) + ".jpg"
}

// 指针接收器的方法只有在过滤时候传送指针才可以保证此方法被正常调用
func (i *Image) GetImagePath() string {
    return i.Path + i.Name + ".png"
}

// 计算热度
func (i Image) GetHot() int {
    return i.Like * i.Forward * i.Collect
}

func TestFunc(t *testing.T) {
    img := Image{
        Url:     []byte("url"),
        Path:    "path",
        Name:    "_golang",
        Collect: 10,
        Like:    100,
        Forward: 50,
    }
    fmt.Println(filter.Select("img", img))
    //{"hot":50000,"path":"path","url":"url.jpg"}

    fmt.Println(filter.Select("img", &img)) //只有传入指针才可以调用绑定指针接收器方法
    //{"hot":50000,"path":"path_golang.png","url":"url.jpg"}
}
```

#### $any标识符任意场景解析

你可能想要某个字段在任意场景下都被选中/或者任意场景都排除该字段，但是又不想在任何场景都写一遍。这样可以用$any 标识符干净整洁的完成你的需求。

```go
type User struct {
    UID    uint   `json:"uid,select($any)"`
  Password string `json:"password,omit($any)"`
}

//这样无论是任何场景
Select("无论这里选择任何场景", user)//无论何种场景都会输出UID的字段。

Omit("无论这里选择任何场景",user)//无论何种场景都会排除password 字段。


f:=filter.Select("场景",要过滤的结构体/map/切片/数组) 

fmt.Println(f) //---> 就不需要上面那样去json.Marshal 了，因为这样直接就是返回的json字符串

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

    fmt.Println(filter.Select("article", m))
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

    fmt.Println(filter.Select("justName", tags))
    //--->输出结果： [{"name":"c"},{"name":"c++"},{"name":"go"}]

    fmt.Println(filter.Select("all", tags))
    //--->输出结果： [{"icon":"icon-c","id":1,"name":"c"},{"icon":"icon-c++","id":1,"name":"c++"},{"icon":"icon-go","id":1,"name":"go"}]

    fmt.Println(filter.Select("chat", tags))
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

    articleJson := filter.Select("article", article)
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

##### 极其复杂的场景深层嵌套-结构体嵌套结构体-嵌套切片-嵌套map-嵌套指针-等等

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
  lang := filter.Select("lang", user)
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
 lookup := filter.Select("lookup", user)
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
    return filter.Select("article",u)
}

func (u User) ProfileResp() interface{} {
    //这样当你后面想要优化性能时可以在这里进行优化，
    return filter.Select("profile",u)
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
    return chat
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
    return filter.Select("profile", u)
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
    //    "uid": 1,
    //    "sex": 1,
    //    "avatar": "avatar"
    //}
}
```
