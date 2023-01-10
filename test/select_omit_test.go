package main

import (
	"testing"
)

func TestOmit(t *testing.T) {

	cases := NewTestCases()
	for _, want := range wants {
		r := eq(want, cases, false)
		if !r.eq {
			t.Errorf("cacheJson:%v noCache:%v s:%v", r.cacheJson, r.noCache, r.s)
		}
	}
	//=== RUN   TestOmit
	//--- PASS: TestOmit (0.12s)
	//PASS
}

func TestSelect(t *testing.T) {

	cases := NewTestCases()
	for _, want := range wants {
		r := eq(want, cases, true)
		if !r.eq {
			t.Errorf("cacheJson:%v noCache:%v s:%v", r.cacheJson, r.noCache, r.s)
		}
	}
	//=== RUN   TestSelect
	//--- PASS: TestSelect (0.04s)
	//PASS
}

var wants = []string{
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
