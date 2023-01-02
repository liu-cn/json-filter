package main

type IntAll struct {
	Int    int    `json:"int,select(all|intAll|Int),omit(Int)"`
	Int8   int8   `json:"int8,select(all|intAll|Int8),omit(Int8)"`
	Int16  int16  `json:"int16,select(all|intAll|Int16),omit(Int16)"`
	Int32  int32  `json:"int32,select(all|intAll|Int32),omit(Int32)"`
	Int64  int64  `json:"int64,select(all|intAll|Int64),omit(Int64)"`
	IntP   *int   `json:"int_p,select(all|intAll|IntP),omit(IntP)"`
	Int8P  *int8  `json:"int8_p,select(all|intAll|Int8P),omit(Int8P)"`
	Int16P *int16 `json:"int16_p,select(all|intAll|Int16P),omit(Int16P)"`
	Int32P *int32 `json:"int32_p,select(all|intAll|Int32P),omit(Int32P)"`
	Int64P *int64 `json:"int64_p,select(all|intAll|Int64P),omit(Int64P)"`
}
type SliceIntAll struct {
	SliceInt    []int    `json:"slice_int,select(all|sliceAll|SliceInt),omit(SliceInt)"`
	SliceInt8   []int8   `json:"slice_int_8,select(all|sliceAll|SliceInt8),omit(SliceInt8)"`
	SliceInt16  []int16  `json:"slice_int_16,select(all|sliceAll|SliceInt16),omit(SliceInt16)"`
	SliceInt32  []int32  `json:"slice_int_32,select(all|sliceAll|SliceInt32),omit(SliceInt32)"`
	SliceInt64  []int64  `json:"slice_int_64,select(all|sliceAll|SliceInt64),omit(SliceInt64)"`
	SliceIntP   []*int   `json:"slice_int_p,select(all|sliceAll|SliceIntP),omit(SliceIntP)"`
	SliceInt8P  []*int8  `json:"slice_int_8_p,select(all|sliceAll|SliceInt8P),omit(SliceInt8P)"`
	SliceInt16P []*int16 `json:"slice_int_16_p,select(all|sliceAll|SliceInt16P),omit(SliceInt16P)"`
	SliceInt32I []*int32 `json:"slice_int_32_i,select(all|sliceAll|SliceInt32I),omit(SliceInt32I)"`
	SliceInt64I []*int64 `json:"slice_int_64_i,select(all|sliceAll|SliceInt64I),omit(SliceInt64I)"`
}
type UintAll struct {
	UInt    uint    `json:"u_int,select(all|UInt),omit(UInt)"`
	UInt8   uint8   `json:"u_int8,select(all|UInt8),omit(UInt8)"`
	UInt16  uint16  `json:"u_int16,select(all|UInt16),omit(UInt16)"`
	UInt32  uint32  `json:"u_int32,select(all|UInt32),omit(UInt32)"`
	UInt64  uint64  `json:"u_int64,select(all|UInt64),omit(UInt64)"`
	UIntP   *uint   `json:"u_intP,select(all|UIntP),omit(UIntP)"`
	UInt8P  *uint8  `json:"u_int_8_p,select(all|UInt8P),omit(UInt8P)"`
	UInt16P *uint16 `json:"u_int_16_p,select(all|UInt16P),omit(UInt16P)"`
	UInt32P *uint32 `json:"u_int_32_p,select(all|UInt32P),omit(UInt32P)"`
	UInt64P *uint64 `json:"u_int_64_p,select(all|UInt64P),omit(UInt64P)"`
}
type SliceUintAll struct {
	SliceUint    []uint    `json:"slice_uint,select(all|sliceAll|SliceUint),omit(SliceUint)"`
	SliceUint8   []uint8   `json:"slice_uint_8,select(all|sliceAll|SliceUint8),omit(SliceUint8)"`
	SliceUint16  []uint16  `json:"slice_uint_16,select(all|sliceAll|SliceUint16),omit(SliceUint16)"`
	SliceUint32  []uint32  `json:"slice_uint_32,select(all|sliceAll|SliceUint32),omit(SliceUint32)"`
	SliceUint64  []uint64  `json:"slice_uint_64,select(all|sliceAll|SliceUint64),omit(SliceUint64)"`
	SliceUintP   []*uint   `json:"slice_uint_p,select(all|sliceAll|SliceUintP),omit(SliceUintP)"`
	SliceUint8P  []*uint8  `json:"slice_uint_8_p,select(all|sliceAll|SliceUint8P),omit(SliceUint8P)"`
	SliceUint16P []*uint16 `json:"slice_uint_16_p,select(all|sliceAll|SliceUint16P),omit(SliceUint16P)"`
	SliceUint32P []*uint32 `json:"slice_uint_32_p,select(all|sliceAll|SliceUint32P),omit(SliceUint32P)"`
	SliceUint64P []*uint64 `json:"slice_uint_64_p,select(all|sliceAll|SliceUint64P),omit(SliceUint64P)"`
}
type FloatAll struct {
	Float64  float64  `json:"float64,select(all|Float64),omit(Float64)"`
	Float64P *float64 `json:"float_64_p,select(all|Float64P),omit(Float64P)"`
	Float32  float32  `json:"float32,select(all|Float32),omit(Float32)"`
	Float32P *float32 `json:"float_32_p,select(all|Float32P),omit(Float32P)"`
}
type SliceFloatAll struct {
	SliceFloat64  []float64  `json:"slice_float_64,select(all|sliceAll|SliceFloat64),omit(SliceFloat64)"`
	SliceFloat64P []*float64 `json:"slice_float_64_p,select(all|sliceAll|SliceFloat64P),omit(SliceFloat64P)"`
	SliceFloat32  []float32  `json:"slice_float_32,select(all|sliceAll|SliceFloat32),omit(SliceFloat32)"`
	SliceFloat32P []*float32 `json:"slice_float_32_p,select(all|sliceAll|SliceFloat32P),omit(SliceFloat32P)"`
}
type BoolAll struct {
	Bool  bool  `json:"bool,select(all|Bool),omit(Bool)"`
	BoolP *bool `json:"bool_p,select(all|BoolP),omit(BoolP)"`
}
type SliceBoolAll struct {
	SliceBool  []bool  `json:"slice_bool,select(all|sliceAll|SliceBool),omit(SliceBool)"`
	SliceBoolP []*bool `json:"slice_bool_p,select(all|sliceAll|SliceBoolP),omit(SliceBoolP)"`
}
type StringAll struct {
	String  string  `json:"string,select(all|String),omit(String)"`
	StringP *string `json:"string_p,select(all|StringP),omit(StringP)"`
}
type SliceStringAll struct {
	SliceString  []string  `json:"slice_string,select(all|sliceAll|SliceString),omit(SliceString)"`
	SliceStringS []*string `json:"slice_string_s,select(all|sliceAll|SliceStringS),omit(SliceStringS)"`
}
type InterfaceAll struct {
	Interface  interface{} `json:"interface,select(all|Interface),omit(Interface)"`
	InterfaceP interface{} `json:"interface_p,select(all|InterfaceP),omit(InterfaceP)"`
}
type SliceInterfaceAll struct {
	SliceInterface  []interface{} `json:"slice_interface,select(all|sliceAll|SliceInterface),omit(SliceInterface)"`
	SliceInterfaceP []interface{} `json:"slice_interface_p,select(all|sliceAll|SliceInterfaceP),omit(SliceInterfaceP)"`
}

type All struct {
	IntAll            `json:"int_all,select($any)"`
	UintAll           `json:"uint_all,select($any)"`
	FloatAll          `json:"float_all,select($any)"`
	BoolAll           `json:"bool_all,select($any)"`
	StringAll         `json:"string_all,select($any)"`
	SliceIntAll       `json:"slice_int_all,select($any)"`
	SliceUintAll      `json:"slice_uint_all,select($any)"`
	SliceFloatAll     `json:"slice_float_all,select($any)"`
	SliceBoolAll      `json:"slice_bool_all,select($any)"`
	SliceStringAll    `json:"slice_string_all,select($any)"`
	InterfaceAll      `json:"interface_all,select($any)"`
	SliceInterfaceAll `json:"slice_interface_all,select($any)"`
}
