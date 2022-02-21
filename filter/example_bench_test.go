package filter

import (
	"encoding/json"
	"testing"
)

func JsonMarshal() string {
	marshal, err := json.Marshal(NewUsers())
	if err != nil {
		panic(err)
	}
	return string(marshal)
}

func BenchmarkUserExample(b *testing.B) {

	b.Run("justName", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SelectMarshal("justName", NewUsers())
		}
	})

	b.Run("lookup", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SelectMarshal("lookup", NewUsers())
		}
	})

	b.Run("profile", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SelectMarshal("profile", NewUsers())
		}
	})

	b.Run("chat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SelectMarshal("chat", NewUsers())
		}
	})

	b.Run("lang", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SelectMarshal("lang", NewUsers())
		}
	})

	b.Run("json(官方原生json解析)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			JsonMarshal()
		}
	})

	//goos: darwin
	//goarch: amd64
	//pkg: github.com/liu-cn/json-filter/filter
	//cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
	//BenchmarkUserExample
	//BenchmarkUserExample/justName
	//BenchmarkUserExample/justName-16         	  250756	      4635 ns/op
	//BenchmarkUserExample/lookup
	//BenchmarkUserExample/lookup-16           	   44769	     26982 ns/op
	//BenchmarkUserExample/profile
	//BenchmarkUserExample/profile-16          	   42028	     24887 ns/op
	//BenchmarkUserExample/chat
	//BenchmarkUserExample/chat-16             	  218448	      5309 ns/op
	//BenchmarkUserExample/lang
	//BenchmarkUserExample/lang-16             	  132003	      9233 ns/op
	//BenchmarkUserExample/json(官方原生json解析)
	//BenchmarkUserExample/json(官方原生json解析)-16                     	  264424	      3855 ns/op

	//	可以看到在选择字段较少的情况下是接近原生json解析的，选择的字段越多越消耗性能（如果需要全字段解析一定要用官方的json库解析），
	//	因为json-filter在过滤的时候是把结构体所有字段重新编码了一遍，所以不可避免的需要有额外的开销。
}
