package filter

import (
	"encoding/json"
	"fmt"
	"testing"
)

func NewExample() Example {

	Int := 100
	Int8 := int8(8)
	Int16 := int16(16)
	Int32 := int32(32)
	Int64 := int64(64)

	uInt := uint(100)
	uInt8 := uint8(8)
	uInt16 := uint16(16)
	uInt32 := uint32(32)
	uInt64 := uint64(64)

	Bool := true
	f32p := float32(320.1)
	f64p := 320.1
	Byte := byte(10)
	nameP := "nameP"
	str := "string p"
	interfaces := "interface p"
	ageP := 10
	tests := Example{
		Int:      100,
		Int8:     8,
		Int16:    16,
		Int32:    32,
		Int64:    64,
		IntP:     &Int,
		Int8P:    &Int8,
		Int16P:   &Int16,
		Int32P:   &Int32,
		Int64P:   &Int64,
		UInt:     uint(1000),
		UInt8:    uint8(80),
		UInt16:   uint16(160),
		UInt32:   uint32(320),
		UInt64:   uint64(640),
		UIntP:    &uInt,
		UInt8P:   &uInt8,
		UInt16P:  &uInt16,
		UInt32P:  &uInt32,
		UInt64P:  &uInt64,
		Bool:     true,
		BoolP:    &Bool,
		Float32:  32.1,
		Float32P: &f32p,
		Float64:  64.1,
		Float64P: &f64p,
		Byte:     1,
		ByteP:    &Byte,

		String:     "string",
		StringP:    &str,
		Interface:  "interface",
		InterfaceP: &interfaces,

		Struct:  struct{}{},
		StructP: &struct{}{},
		Structs: Users{
			Name: "name",
			Age:  10,
			Struct: Child{
				CAge:  100,
				CName: "cname",
			},
		},

		StructsP: &UserP{
			Name: &nameP,
			Age:  &ageP,
			Struct: &ChildP{
				CName: &nameP,
				CAge:  &ageP,
			},
		},
	}
	return tests
}

func testSelector(selector string) string {
	return SelectMarshal(selector, NewExample())
}
func testJSON() string {
	j, err := json.Marshal(NewExample())
	if err != nil {
		panic(err)
	}
	return string(j)
}

func TestExample(t *testing.T) {
	t.Run("select-all", func(t *testing.T) {
		fmt.Println(testSelector("all"))
		//{"bool":true,"bool_p":true,"byte":1,"byte_p":10,"float32":32.1,"float64":64.1,"float_32_p":320.1,"float_64_p":320.1,"int":100,"int16":16,"int16_p":16,"int32":32,"int32_p":32,"int64":64,"int64_p":64,"int8":8,"int8_p":8,"int_p":100,"interface":"interface","interface_p":"interface p","string":"string","string_p":"string p","struct":null,"struct_p":null,"structs":{"age":10,"name":"name","struct":{"c_age":100,"c_name":"cname"}},"structs_p":{"age":10,"name":"nameP","struct":{"c_age":10,"c_name":"nameP"}},"u_int":1000,"u_int16":160,"u_int32":320,"u_int64":640,"u_int8":80,"u_intP":100,"u_int_16_p":16,"u_int_32_p":32,"u_int_64_p":64,"u_int_8_p":8}
	})
	t.Run("json", func(t *testing.T) {
		fmt.Println(testJSON())
		//{"int":100,"int8":8,"int16":16,"int32":32,"int64":64,"int_p":100,"int8_p":8,"int16_p":16,"int32_p":32,"int64_p":64,"u_int":1000,"u_int8":80,"u_int16":160,"u_int32":320,"u_int64":640,"u_intP":100,"u_int_8_p":8,"u_int_16_p":16,"u_int_32_p":32,"u_int_64_p":64,"float64":64.1,"float_64_p":320.1,"float32":32.1,"float_32_p":320.1,"bool":true,"bool_p":true,"byte":1,"byte_p":10,"string":"string","string_p":"string p","interface":"interface","interface_p":"interface p","struct":{},"struct_p":{},"structs":{"name":"name","age":10,"struct":{"c_name":"cname","c_age":100}},"structs_p":{"name":"nameP","age":10,"struct":{"c_name":"nameP","c_age":10}}}
	})

	t.Run("select-intAll", func(t *testing.T) {
		fmt.Println(testSelector("intAll"))
		//{"int":100,"int16":16,"int16_p":16,"int32":32,"int32_p":32,"int64":64,"int64_p":64,"int8":8,"int8_p":8,"int_p":100}

	})
	t.Run("select-struct", func(t *testing.T) {
		fmt.Println(testSelector("struct"))
		//{"structs":{"struct":{"c_age":100,"c_name":"cname"}},"structs_p":{"struct":{"c_age":10}}}
	})
}

func BenchmarkExample(b *testing.B) {

	b.Run("select-all", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testSelector("all")
			//BenchmarkExample/select-all-16         	   24118	     42690 ns/op
		}
	})

	b.Run("select-intAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testSelector("intAll")
			//BenchmarkExample/select-intAll-16      	   64729	     16692 ns/op
		}
	})
	b.Run("select-struct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testSelector("struct")
			//BenchmarkExample/select-struct-16      	   51351	     21172 ns/op
		}
	})

	b.Run("json", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testJSON()
			//BenchmarkExample/json-16               	  317278	      3324 ns/op
		}
	})

	//goos: darwin
	//goarch: amd64
	//pkg: filter/filter
	//cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
	//BenchmarkExample
	//BenchmarkExample/select-all
	//BenchmarkExample/select-all-16         	   24118	     42690 ns/op
	//BenchmarkExample/select-intAll
	//BenchmarkExample/select-intAll-16      	   64729	     16692 ns/op
	//BenchmarkExample/select-struct
	//BenchmarkExample/select-struct-16      	   51351	     21172 ns/op
	//BenchmarkExample/json
	//BenchmarkExample/json-16               	  317278	      3324 ns/op
	//PASS
}
