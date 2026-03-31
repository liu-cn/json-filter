# json-filter

基于场景的 Go JSON 字段过滤器。它让同一个 struct 通过 `json` tag 复用到多个接口响应里，而不是为每个接口维护一份单独的 DTO。

Scene-based JSON field filtering for Go. Reuse one model across multiple API responses with `json` tags while staying compatible with `encoding/json`.

- [中文](#中文)
- [English](#english)

## 中文

### 这库解决什么问题

很多项目里，同一个模型会在不同接口里返回不同字段：

- 文章列表只需要 `uid`、`nickname`、`avatar`
- 个人中心还要返回 `sex`、`vip_end_time`、`price`
- 管理后台可能又是一套字段

传统做法通常是：

- 新建很多响应结构体
- 手动拷贝字段
- 或者直接把原始 struct 全量返回

`json-filter` 的做法是把“哪些字段在哪些场景下返回”直接写在 `json` tag 里，然后用 `Select` / `Omit` 在运行时生成最终 JSON。

### 特性

- 支持 `struct`、`map`、`slice`、`array`、指针、`interface{}`
- 支持深层嵌套组合
- 支持 `select(...)`、`omit(...)`、`omitempty`、`func(...)`
- 支持 `$any`
- 支持匿名字段展开
- 兼容 `encoding/json`
- 支持自定义 `json.Marshaler` / `encoding.TextMarshaler`
- Go 版本要求：`1.17+`

### 安装

```bash
go get github.com/liu-cn/json-filter
```

### 1 分钟上手

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/liu-cn/json-filter/filter"
)

type User struct {
	UID      uint   `json:"uid,select(article)"`
	Avatar   string `json:"avatar,select(article)"`
	Nickname string `json:"nickname,select(article|profile)"`

	Sex        int       `json:"sex,select(profile)"`
	VipEndTime time.Time `json:"vip_end_time,select(profile)"`
	Price      string    `json:"price,select(profile)"`
}

func main() {
	user := User{
		UID:        1,
		Avatar:     "avatar",
		Nickname:   "boyan",
		Sex:        1,
		VipEndTime: time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
		Price:      "999.9",
	}

	article, _ := json.Marshal(filter.Select("article", user))
	fmt.Println(string(article))
	// {"avatar":"avatar","nickname":"boyan","uid":1}

	fmt.Println(filter.Select("profile", user))
	// {"nickname":"boyan","price":"999.9","sex":1,"vip_end_time":"2026-04-01T00:00:00Z"}
}
```

### 推荐怎么用

有两个入口层级：

- `Select(scene, value)` / `Omit(scene, value)`
  直接返回一个可以交给 `json.Marshal`、`gin.Context.JSON` 的值。适合接口直接返回。
- `SelectFilter(scene, value)` / `OmitFilter(scene, value)`
  返回 typed 的 `Filter`，适合你还想继续取 `JSON`、`Bytes`、`Map`、`Slice`、`Interface`。

推荐约定：

- 直接响应 HTTP：优先用 `Select` / `Omit`
- 需要继续处理过滤结果：优先用 `SelectFilter` / `OmitFilter`

### 核心 API

```go
filter.Select(scene, value)
filter.Omit(scene, value)

filter.SelectFilter(scene, value)
filter.OmitFilter(scene, value)
```

`Filter` 提供这些方法：

- `JSON() (string, error)`
- `MustJSON() string`
- `Bytes() ([]byte, error)`
- `MustBytes() []byte`
- `Interface() interface{}`
- `Map() map[string]interface{}`
- `Slice() []interface{}`

说明：

- `Map()` 只适合顶层结果是对象时使用
- `Slice()` 只适合顶层结果是数组时使用
- 如果顶层结果可能是标量或 `null`，优先用 `Interface()`、`JSON()`、`Bytes()`

### Tag 规则

#### `select(...)`

只在指定场景下输出该字段。

```go
Name string `json:"name,select(article|profile)"`
```

#### `omit(...)`

在指定场景下排除该字段。

```go
Password string `json:"password,omit(profile|admin)"`
```

#### `omitempty`

零值时忽略该字段。

```go
Nickname string  `json:"nickname,omitempty,select(profile)"`
Avatar   *string `json:"avatar,omitempty,select(profile)"`
Age      int     `json:"age,omitempty,select(profile)"`
```

#### `func(...)`

在过滤时调用当前 struct 上的零参数方法，用方法返回值替代字段值。

```go
Avatar string `json:"avatar,select(profile),func(BuildAvatar)"`
```

注意：

- 方法是定义在“当前字段所属的 struct”上的
- 如果方法是指针接收器，请给 `Select` / `Omit` 传指针

#### `$any`

表示任意场景都匹配。

```go
UID      uint   `json:"uid,select($any)"`
Password string `json:"password,omit($any)"`
```

#### 匿名字段展开

字段名留空时，会把匿名嵌入结构体展开到当前层级。

```go
type Page struct {
	PageInfo int `json:"page_info,select(article)"`
}

type Article struct {
	Title string `json:"title,select(article)"`
	Page  `json:",select(article)"`
}
```

输出会类似：

```json
{"page_info":1,"title":"hello"}
```

如果你写成：

```go
Page Page `json:"page,select(article)"`
```

那它会作为普通嵌套对象输出，而不会展开。

### 行为说明

这几条语义建议先了解清楚：

- `select` 模式下，只会保留带 `json` tag 且命中 `select(...)` 的字段
- `omit` 模式下，字段默认保留；没有 `json` tag 的字段也会保留，并使用 struct 字段名
- `json:"-"` 始终忽略
- `nil` 指针、`nil` 元素会按 `null` 处理
- 顶层结果可以是对象、数组、标量或 `null`
- `map[bool]...`、数字 key map、字符串 key map 都支持
- 自定义 `json.Marshaler` / `encoding.TextMarshaler` 会按叶子值处理

### 示例

#### `omit` 示例

```go
type User struct {
	Name     string `json:"name"`
	Password string `json:"password,omit($any)"`
	Phone    string `json:"phone,omit(public)"`
}

user := User{
	Name:     "boyan",
	Password: "123456",
	Phone:    "18800000000",
}

fmt.Println(filter.Omit("public", user))
// {"name":"boyan"}
```

#### `func(...)` 示例

```go
type Image struct {
	URL     string `json:"url,select(api),func(BuildURL)"`
	Name    string
	Ext     string
}

func (i Image) BuildURL() string {
	return i.Name + i.Ext
}

fmt.Println(filter.Select("api", Image{
	Name: "avatar",
	Ext:  ".png",
}))
// {"url":"avatar.png"}
```

#### 继续处理过滤结果

```go
f := filter.SelectFilter("profile", user)

jsonStr, err := f.JSON()
if err != nil {
	panic(err)
}
fmt.Println(jsonStr)

bs, err := f.Bytes()
if err != nil {
	panic(err)
}
fmt.Println(string(bs))

payload := f.Interface()
_ = payload
```

#### 在 Gin 中使用

```go
func GetUser(c *gin.Context) {
	user := User{
		UID:      1,
		Avatar:   "avatar",
		Nickname: "boyan",
		Sex:      1,
	}

	c.JSON(200, filter.Select("profile", user))
}
```

### 什么时候适合用它

适合：

- 中小型后端项目
- 同一模型复用到多个响应场景
- 想快速减少 DTO 重复定义

不太适合：

- 极度追求零反射开销的链路
- 想把返回结构显式写死在类型系统里
- 团队不希望把接口规则写进 tag

### 兼容性说明

- 当前模块 `go.mod` 为 `go 1.17`
- 与标准库 `encoding/json` 配合使用
- 老 API 仍保留兼容：
  - `SelectMarshal` / `OmitMarshal`
  - `MustMarshalJSON`
  - `MastMarshalJSON`

新代码建议优先使用：

- `SelectFilter` / `OmitFilter`
- `Bytes` / `MustBytes`

### 更多示例

可以直接看仓库里的示例和测试：

- `example/`
- `filter/example_test.go`
- `filter/filter_test.go`

## English

### What It Does

`json-filter` is a scene-based JSON field filter for Go. It lets one struct serve
multiple API responses by encoding scene rules directly in `json` tags.

### Install

```bash
go get github.com/liu-cn/json-filter
```

Go version: `1.17+`

### Quick Example

```go
type User struct {
	UID      uint   `json:"uid,select(article)"`
	Avatar   string `json:"avatar,select(article)"`
	Nickname string `json:"nickname,select(article|profile)"`
	Sex      int    `json:"sex,select(profile)"`
}

fmt.Println(filter.Select("article", user))
// {"avatar":"avatar","nickname":"boyan","uid":1}
```

### Recommended API

- `Select(scene, value)` / `Omit(scene, value)`
  Use these when you want to pass the result directly to `json.Marshal` or an
  HTTP framework response helper.
- `SelectFilter(scene, value)` / `OmitFilter(scene, value)`
  Use these when you want the typed `Filter` helpers such as `JSON`, `Bytes`,
  `Map`, `Slice`, or `Interface`.

### Tag Syntax

- `select(article|profile)` include in these scenes
- `omit(admin|internal)` exclude in these scenes
- `omitempty` drop zero values
- `func(BuildValue)` call a zero-arg method on the containing struct
- `$any` match every scene

### Notes

- In `select` mode, only fields with matching `select(...)` tags are included.
- In `omit` mode, fields are included by default unless excluded.
- `json:"-"` is always ignored.
- Structs, maps, slices, arrays, pointers, interfaces, and nested combinations are supported.
- Custom `json.Marshaler` and `encoding.TextMarshaler` leaf values are supported.

### License

[MIT](LICENSE)
