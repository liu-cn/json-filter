package filter

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestTagReg(t *testing.T) {
	//re := "select(art,article,chat)"

	type Model struct {
		Name string `json:"name,omitempty,select(req|res),omit(chat|profile|article)"`
	}

	tag := reflect.TypeOf(&Model{}).Elem().Field(0).Tag.Get("json")
	fmt.Println(tag)

	selectStr := ""
	omitStr := ""

	tags := strings.Split(tag, ",")

	fieldName := tags[0]
	for _, s := range tags {
		//fmt.Println(s)
		if strings.HasPrefix(s, "select(") {
			selectStr = s[7 : len(s)-1]
		}

		if strings.HasPrefix(s, "omit(") {
			omitStr = s[5 : len(s)-1]
		}
	}
	fmt.Println(selectStr)
	fmt.Println(omitStr)
	fmt.Println(fieldName)

}

func OmitTest() {
	_ = NewOmitTag("name,omitempty,select(req|res),omit(chat|profile|article)", "article", "IsOmitField:true")
}

func BenchmarkTags(b *testing.B) {

	b.Run("select", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewSelectTag("name,omitempty,select(req|res|user),omit(chat|profile|article)", "user", "Name")
		}
	})
	b.Run("select-f", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewSelectTag("name,omitempty,select(req|res|user),omit(chat|profile|article)", "req", "Name")
		}
	})
	b.Run("omit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewOmitTag("name,omitempty,select(req|res|user),omit(chat|profile|article)", "article", "Name")
		}
	})
	b.Run("omit-f", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewOmitTag("name,omitempty,select(req|res|user),omit(chat|profile|article)", "chat", "Name")
		}
	})

	//goos: windows
	//goarch: amd64
	//pkg: filter
	//cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
	//	BenchmarkTags
	//	BenchmarkTags/select
	//	BenchmarkTags/select-4           4147592               295.9 ns/op
	//	BenchmarkTags/select-f
	//	BenchmarkTags/select-f-4         4163122               287.8 ns/op
	//	BenchmarkTags/omit
	//	BenchmarkTags/omit-4             3761622               306.6 ns/op
	//	BenchmarkTags/omit-f
	//	BenchmarkTags/omit-f-4           3889572               301.8 ns/op
	//	PASS
}
func TestTagRe(t *testing.T) {
	//SelectTest()
	OmitTest()
}
