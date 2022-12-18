package main

import (
	"encoding/json"
	"github.com/liu-cn/json-filter/filter"
	"github.com/liu-cn/pkg/benchmark"
	"testing"
)

func getCacheVal(s string, el interface{}, isSelect bool) string {
	st := ""
	benchmark.TimeAndRes(func() interface{} {
		if isSelect {
			ss := filter.Select(s, el)
			marshal, err := json.Marshal(ss)
			if err != nil {
				panic(err)
			}
			json1 := string(marshal)
			st = json1
			return json1
		} else {
			ss := filter.Omit(s, el)

			marshal, err := json.Marshal(ss)
			if err != nil {
				panic(err)
			}
			json1 := string(marshal)
			st = json1
			return json1
		}
	}, 100)
	return st
}

type result struct {
	s     string
	eq    bool
	json1 string
	json2 string
}

func eq(s string, el interface{}, isSelect bool) result {
	val := getCacheVal(s, el, isSelect)
	var json string
	if isSelect {
		json = filter.SelectMarshal(s, el).MustJSON()
	} else {
		json = filter.OmitMarshal(s, el).MustJSON()
	}
	return result{
		s:     s,
		json1: val,
		json2: json,
		eq:    json == val,
	}

}

func TestOmit(t *testing.T) {

	wants := []string{
		"Time",
		"TimeP",
		"MTime",
		"MTimeP",
		"YTime",
		"YTimeP",
		"Int",
		"Int8",
		"Int16",
		"Int32",
		"Int64",
		"IntP",
		"Int8P",
		"Int16P",
		"Int32P",
		"Int64P",
		"UInt",
		"UInt8",
		"UInt16",
		"UInt32",
		"UInt64",
		"UIntP",
		"UInt8P",
		"UInt16P",
		"UInt32P",
		"UInt64P",
		"Float64",
		"Float64P",
		"Float32",
		"Float32P",
		"Bool",
		"BoolP",
		"Byte",
		"ByteP",
		"String",
		"StringP",
		"SliceInt",
		"SliceInt8",
		"SliceInt16",
		"SliceInt32",
		"SliceInt64",
		"SliceIntP",
		"SliceInt8P",
		"SliceInt16P",
		"SliceInt32I",
		"SliceInt64I",
		"SliceUint",
		"SliceUint8",
		"SliceUint16",
		"SliceUint32",
		"SliceUint64",
		"SliceUintP",
		"SliceUint8P",
		"SliceUint16P",
		"SliceUint32P",
		"SliceUint64P",
		"SliceFloat64",
		"SliceFloat64P",
		"SliceFloat32",
		"SliceFloat32P",
		"SliceBool",
		"SliceBoolP",
		"SliceByte",
		"SliceByteP",
		"SliceString",
		"SliceStringS",
		"SliceInterface",
		"SliceInterfaceP",
		"SliceStruct",
		"SliceStructP",
		"SliceTime",
		"SliceUsersCase",
		"SliceUserP",
	}
	cases := NewTestCases()
	for _, want := range wants {
		r := eq(want, cases, false)
		if !r.eq {
			t.Errorf("json1:%v json2:%v s:%v", r.json1, r.json2, r.s)
		}
	}
	//=== RUN   TestOmit
	//--- PASS: TestOmit (1.91s)
	//PASS

}

func TestSelect(t *testing.T) {

	wants := []string{
		"Time",
		"TimeP",
		"MTime",
		"MTimeP",
		"YTime",
		"YTimeP",
		"Int",
		"Int8",
		"Int16",
		"Int32",
		"Int64",
		"IntP",
		"Int8P",
		"Int16P",
		"Int32P",
		"Int64P",
		"UInt",
		"UInt8",
		"UInt16",
		"UInt32",
		"UInt64",
		"UIntP",
		"UInt8P",
		"UInt16P",
		"UInt32P",
		"UInt64P",
		"Float64",
		"Float64P",
		"Float32",
		"Float32P",
		"Bool",
		"BoolP",
		"Byte",
		"ByteP",
		"String",
		"StringP",
		"SliceInt",
		"SliceInt8",
		"SliceInt16",
		"SliceInt32",
		"SliceInt64",
		"SliceIntP",
		"SliceInt8P",
		"SliceInt16P",
		"SliceInt32I",
		"SliceInt64I",
		"SliceUint",
		"SliceUint8",
		"SliceUint16",
		"SliceUint32",
		"SliceUint64",
		"SliceUintP",
		"SliceUint8P",
		"SliceUint16P",
		"SliceUint32P",
		"SliceUint64P",
		"SliceFloat64",
		"SliceFloat64P",
		"SliceFloat32",
		"SliceFloat32P",
		"SliceBool",
		"SliceBoolP",
		"SliceByte",
		"SliceByteP",
		"SliceString",
		"SliceStringS",
		"SliceInterface",
		"SliceInterfaceP",
		"SliceStruct",
		"SliceStructP",
		"SliceTime",
		"SliceUsersCase",
		"SliceUserP",
	}
	cases := NewTestCases()
	for _, want := range wants {
		r := eq(want, cases, true)
		if !r.eq {
			t.Errorf("json1:%v json2:%v s:%v", r.json1, r.json2, r.s)
		}
	}
	//=== RUN   TestSelect
	//--- PASS: TestSelect (0.37s)
	//PASS
}
