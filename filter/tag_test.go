package filter

import (
	"strings"
	"testing"
)

func GetSelectTag(tag string) []string {
	tags := strings.Split(tag, ",")
	selectTags := make([]string, 0, 5)
	for _, s := range tags {
		if strings.HasPrefix(s, "select(") {
			selectStr := s[7 : len(s)-1]
			scene := strings.Split(selectStr, "|")
			for _, v := range scene {
				selectTags = append(selectTags, v)
			}
		}
	}
	return selectTags
}
func TestTagSelect(t *testing.T) {
	tag := "name,omitempty,select(req|res),omit(chat|profile|article)"
	want := []string{
		"req", "res",
	}
	got := GetSelectTag(tag)

	if len(got) != len(want) {
		t.Errorf("tag 解析不符合预期want:%v got:%v", want, got)
		return
	}

	for i, v := range got {
		if !(v == want[i]) {
			t.Errorf("tag 解析不符合预期want:%v got:%v", want, got)
		}
	}
}
func TestNewSelectTag(t *testing.T) {
	selector := "req"
	name := "name"
	tag := "name,omitempty,select(req|res),omit(chat|profile|article)"
	got := NewSelectTag(tag, "req", "name")
	if got.IsOmitField {
		t.Errorf("IsOmitField 应该为true")
	}
	if !got.IsSelect {
		t.Errorf("IsSelect 应该为true")
	}

	if got.SelectScene != selector {
		t.Errorf("SelectScene 应为%v 实际%v", selector, got.SelectScene)
	}
	if got.FieldName != name {
		t.Errorf("FieldName 应为%v 实际%v", name, got.FieldName)
	}

	//=== RUN   TestNewSelectTag
	//--- PASS: TestNewSelectTag (0.00s)
	//PASS
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
