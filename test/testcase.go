package main

import "time"

type MyTime = time.Time
type YouTime = MyTime

type TestCases struct {
	Article    `json:"article,select(all|Anonymous)"`
	*Anonymous `json:",select(all|Anonymous)"`
	Time       time.Time   `json:"time,select(all|Time),omit(Time)"`
	TimeP      *time.Time  `json:"time_p,select(all|TimeP),omit(TimeP)"`
	MTime      MyTime      `json:"m_time,select(all|MTime),omit(MTime)"`
	MTimeP     *MyTime     `json:"m_time_p,select(all|MTimeP),omit(MTimeP)"`
	YTime      YouTime     `json:"y_time,select(all|YTime),omit(YTime)"`
	YTimeP     *YouTime    `json:"y_time_p,select(all|YTimeP),omit(YTimeP)"`
	Int        int         `json:"int,select(all|intAll|Int),omit(Int)"`
	Int8       int8        `json:"int8,select(all|intAll|Int8),omit(Int8)"`
	Int16      int16       `json:"int16,select(all|intAll|Int16),omit(Int16)"`
	Int32      int32       `json:"int32,select(all|intAll|Int32),omit(Int32)"`
	Int64      int64       `json:"int64,select(all|intAll|Int64),omit(Int64)"`
	IntP       *int        `json:"int_p,select(all|intAll|IntP),omit(IntP)"`
	Int8P      *int8       `json:"int8_p,select(all|intAll|Int8P),omit(Int8P)"`
	Int16P     *int16      `json:"int16_p,select(all|intAll|Int16P),omit(Int16P)"`
	Int32P     *int32      `json:"int32_p,select(all|intAll|Int32P),omit(Int32P)"`
	Int64P     *int64      `json:"int64_p,select(all|intAll|Int64P),omit(Int64P)"`
	UInt       uint        `json:"u_int,select(all|UInt),omit(UInt)"`
	UInt8      uint8       `json:"u_int8,select(all|UInt8),omit(UInt8)"`
	UInt16     uint16      `json:"u_int16,select(all|UInt16),omit(UInt16)"`
	UInt32     uint32      `json:"u_int32,select(all|UInt32),omit(UInt32)"`
	UInt64     uint64      `json:"u_int64,select(all|UInt64),omit(UInt64)"`
	UIntP      *uint       `json:"u_intP,select(all|UIntP),omit(UIntP)"`
	UInt8P     *uint8      `json:"u_int_8_p,select(all|UInt8P),omit(UInt8P)"`
	UInt16P    *uint16     `json:"u_int_16_p,select(all|UInt16P),omit(UInt16P)"`
	UInt32P    *uint32     `json:"u_int_32_p,select(all|UInt32P),omit(UInt32P)"`
	UInt64P    *uint64     `json:"u_int_64_p,select(all|UInt64P),omit(UInt64P)"`
	Float64    float64     `json:"float64,select(all|Float64),omit(Float64)"`
	Float64P   *float64    `json:"float_64_p,select(all|Float64P),omit(Float64P)"`
	Float32    float32     `json:"float32,select(all|Float32),omit(Float32)"`
	Float32P   *float32    `json:"float_32_p,select(all|Float32P),omit(Float32P)"`
	Bool       bool        `json:"bool,select(all|Bool),omit(Bool)"`
	BoolP      *bool       `json:"bool_p,select(all|BoolP),omit(BoolP)"`
	Byte       byte        `json:"byte,select(all|byteAll|Byte),omit(Byte)"`
	ByteP      *byte       `json:"byte_p,select(all|byteAll|ByteP),omit(ByteP)"`
	String     string      `json:"string,select(all|String),omit(String)"`
	StringP    *string     `json:"string_p,select(all|StringP),omit(StringP)"`
	Interface  interface{} `json:"interface,select(all|Interface),omit(Interface)"`
	InterfaceP interface{} `json:"interface_p,select(all|InterfaceP),omit(InterfaceP)"`
	Struct     struct{}    `json:"struct,select(all|struct|Struct),omit(Struct)"`
	StructEl   struct {
		Name string `json:"name,select(all|struct)"`
	} `json:"struct_el,select(all|struct)"`
	StructP     *struct{}               `json:"struct_p,select(all|struct)"`
	Structs     UsersCase               `json:"structs,select(all|struct)"`
	StructsP    *UserP                  `json:"structs_p,select(all|struct)"`
	Map         map[string]interface{}  `json:"map,select(all|mapAll)"`
	MapP        *map[string]interface{} `json:"map_p,select(all|mapAll)"`
	SliceInt    []int                   `json:"slice_int,select(all|sliceAll|SliceInt),omit(SliceInt)"`
	SliceInt8   []int8                  `json:"slice_int_8,select(all|sliceAll|SliceInt8),omit(SliceInt8)"`
	SliceInt16  []int16                 `json:"slice_int_16,select(all|sliceAll|SliceInt16),omit(SliceInt16)"`
	SliceInt32  []int32                 `json:"slice_int_32,select(all|sliceAll|SliceInt32),omit(SliceInt32)"`
	SliceInt64  []int64                 `json:"slice_int_64,select(all|sliceAll|SliceInt64),omit(SliceInt64)"`
	SliceIntP   []*int                  `json:"slice_int_p,select(all|sliceAll|SliceIntP),omit(SliceIntP)"`
	SliceInt8P  []*int8                 `json:"slice_int_8_p,select(all|sliceAll|SliceInt8P),omit(SliceInt8P)"`
	SliceInt16P []*int16                `json:"slice_int_16_p,select(all|sliceAll|SliceInt16P),omit(SliceInt16P)"`
	SliceInt32I []*int32                `json:"slice_int_32_i,select(all|sliceAll|SliceInt32I),omit(SliceInt32I)"`
	SliceInt64I []*int64                `json:"slice_int_64_i,select(all|sliceAll|SliceInt64I),omit(SliceInt64I)"`
	SliceUint   []uint                  `json:"slice_uint,select(all|sliceAll|SliceUint),omit(SliceUint)"`
	SliceUint8  []uint8                 `json:"slice_uint_8,select(all|sliceAll|SliceUint8),omit(SliceUint8)"`
	SliceUint16 []uint16                `json:"slice_uint_16,select(all|sliceAll|SliceUint16),omit(SliceUint16)"`
	SliceUint32 []uint32                `json:"slice_uint_32,select(all|sliceAll|SliceUint32),omit(SliceUint32)"`
	SliceUint64 []uint64                `json:"slice_uint_64,select(all|sliceAll|SliceUint64),omit(SliceUint64)"`

	SliceUintP          []*uint                   `json:"slice_uint_p,select(all|sliceAll|SliceUintP),omit(SliceUintP)"`
	SliceUint8P         []*uint8                  `json:"slice_uint_8_p,select(all|sliceAll|SliceUint8P),omit(SliceUint8P)"`
	SliceUint16P        []*uint16                 `json:"slice_uint_16_p,select(all|sliceAll|SliceUint16P),omit(SliceUint16P)"`
	SliceUint32P        []*uint32                 `json:"slice_uint_32_p,select(all|sliceAll|SliceUint32P),omit(SliceUint32P)"`
	SliceUint64P        []*uint64                 `json:"slice_uint_64_p,select(all|sliceAll|SliceUint64P),omit(SliceUint64P)"`
	SliceFloat64        []float64                 `json:"slice_float_64,select(all|sliceAll|SliceFloat64),omit(SliceFloat64)"`
	SliceFloat64P       []*float64                `json:"slice_float_64_p,select(all|sliceAll|SliceFloat64P),omit(SliceFloat64P)"`
	SliceFloat32        []float32                 `json:"slice_float_32,select(all|sliceAll|SliceFloat32),omit(SliceFloat32)"`
	SliceFloat32P       []*float32                `json:"slice_float_32_p,select(all|sliceAll|SliceFloat32P),omit(SliceFloat32P)"`
	SliceBool           []bool                    `json:"slice_bool,select(all|sliceAll|SliceBool),omit(SliceBool)"`
	SliceBoolP          []*bool                   `json:"slice_bool_p,select(all|sliceAll|SliceBoolP),omit(SliceBoolP)"`
	SliceByte           []byte                    `json:"slice_byte,select(all|sliceAll|SliceByte),omit(SliceByte)"`
	SliceByteP          []*byte                   `json:"slice_byte_p,select(all|sliceAll|SliceByteP),omit(SliceByteP)"`
	SliceString         []string                  `json:"slice_string,select(all|sliceAll|SliceString),omit(SliceString)"`
	SliceStringS        []*string                 `json:"slice_string_s,select(all|sliceAll|SliceStringS),omit(SliceStringS)"`
	SliceInterface      []interface{}             `json:"slice_interface,select(all|sliceAll|SliceInterface),omit(SliceInterface)"`
	SliceInterfaceP     []interface{}             `json:"slice_interface_p,select(all|sliceAll|SliceInterfaceP),omit(SliceInterfaceP)"`
	SliceStruct         []struct{}                `json:"slice_struct,select(all|sliceAll|SliceStruct),omit(SliceStruct)"`
	SliceStructP        []*struct{}               `json:"slice_struct_p,select(all|sliceAll|SliceStructP),omit(SliceStructP)"`
	SliceTime           []time.Time               `json:"slice_time,select(all|sliceAll|SliceTime),omit(SliceTime)"`
	SliceUsersCase      []UsersCase               `json:"slice_users_case,select(all|sliceAll|SliceUsersCase),omit(SliceUsersCase)"`
	SliceUserP          []*UserP                  `json:"slice_user_p,select(all|sliceAll|SliceUserP),omit(SliceUserP)"`
	SliceMap            []map[string]interface{}  `json:"slice_map,select(all|sliceAll)"`
	SliceMapP           []*map[string]interface{} `json:"slice_map_p,select(all|sliceAll)"`
	SliceSliceInt       [][]int                   `json:"slice_slice_int,select(all|sliceAll)"`
	SliceSliceUsersCase [][]UsersCase             `json:"slice_slice_users_case,select(all|sliceAll)"`
}

type Child struct {
	CName string `json:"c_name,select(all|2|struct)"`
	CAge  int    `json:"c_age,select(all|struct)"`
}
type UsersCase struct {
	Name   string `json:"name,select(all|1),omit(1)"`
	Age    int    `json:"age,select(all|2),omit(1)"`
	Struct Child  `json:"struct,select(all|struct)"`
}

type ChildP struct {
	CName *string `json:"c_name,select(all|2)"`
	CAge  *int    `json:"c_age,select(all|struct)"`
}
type UserP struct {
	Name   *string `json:"name,select(all|1)"`
	Age    *int    `json:"age,select(all|2)"`
	Struct *ChildP `json:"struct,select(all|struct)"`
}

type BaseInfo struct {
	Name     string `json:"name,select(all|Anonymous)"`
	Title    string `json:"title,select(all|Anonymous)"`
	PageInfo `json:",select(all|Anonymous)"`
}

type PageInfo struct {
	PageNum  int `json:"page_num,select(all|Anonymous)"`
	PageSize int `json:"page_size,select(all|Anonymous)"`
}

type Article struct {
	Price    string `json:"price,select(all|Anonymous)"`
	BaseInfo `json:",select(all|Anonymous)"`
}

type AnonymousValue struct {
	AnonymousValueName string `json:"anonymous_value_name,select(all|Anonymous)"`
}
type Anonymous struct {
	AnonymousValue `json:",select(all|Anonymous)"`
}

func NewTestCases() TestCases {

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
	tests := TestCases{
		Anonymous: &Anonymous{
			AnonymousValue{
				AnonymousValueName: "anonymous_value_name",
			},
		},
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

		Struct: struct{}{},

		StructEl: struct {
			Name string `json:"name,select(all|struct)"`
		}(struct{ Name string }{Name: "el"}),

		Map: map[string]interface{}{
			"string": "map val",
			"struct": UsersCase{
				Name: "hhhhh",
			},
		},
		MapP: &map[string]interface{}{
			"map_p": "map val",
		},
		StructP: &struct{}{},
		Structs: UsersCase{
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

		SliceBool:  []bool{Bool, Bool},
		SliceBoolP: []*bool{&Bool},
		SliceByte:  []byte{Byte},
		SliceByteP: []*byte{&Byte},

		SliceInt: []int{
			1, 2,
		},
		SliceInt8: []int8{
			1, 2,
		},
		SliceInt16: []int16{
			1, 2,
		},
		SliceInt32: []int32{
			1, 2,
		},
		SliceInt64: []int64{
			1, 2,
		},
		SliceIntP: []*int{
			&Int, &Int,
		},
		SliceInt8P: []*int8{
			&Int8,
		},
		SliceInt16P: []*int16{
			&Int16,
		},
		SliceInt32I: []*int32{
			&Int32,
		},
		SliceInt64I: []*int64{
			&Int64,
		},
		SliceUint:     []uint{1, 3},
		SliceUint8:    []uint8{1, 2},
		SliceUint16:   []uint16{1, 3},
		SliceUint32:   []uint32{1, 2},
		SliceUint64:   []uint64{1, 4},
		SliceUintP:    []*uint{&uInt},
		SliceUint8P:   []*uint8{&uInt8},
		SliceUint16P:  []*uint16{&uInt16},
		SliceUint32P:  []*uint32{&uInt32},
		SliceUint64P:  []*uint64{&uInt64},
		SliceFloat64:  []float64{12.3},
		SliceFloat64P: []*float64{&f64p},
		SliceFloat32:  []float32{12.7},
		SliceFloat32P: []*float32{&f32p},

		SliceString: []string{
			"slice string", "123",
		},
		SliceStringS: []*string{
			&str, &str,
		},
		SliceInterface: []interface{}{
			"12", "13",
		},
		SliceInterfaceP: []interface{}{
			&str, &str,
		},
		SliceStruct:  []struct{}{},
		SliceStructP: []*struct{}{},

		SliceTime: []time.Time{
			time.Now(), time.Now(),
		},
		SliceUsersCase: []UsersCase{
			{Name: nameP},
		},
		SliceUserP: []*UserP{
			{
				Age: &ageP,
			},
		},
		SliceMap: []map[string]interface{}{
			{
				"map1": 1,
			},
			{
				"map2": 2,
			},
		},
		SliceMapP: []*map[string]interface{}{
			{
				"map1": 1,
			},
			{
				"map2": 2,
			},
		},

		SliceSliceInt: [][]int{
			{1, 23},
			{2, 3},
		},
		SliceSliceUsersCase: [][]UsersCase{
			{
				UsersCase{
					Age: 1,
				},
			},
			{
				UsersCase{
					Age: 2,
				},
			},
		},
	}
	return tests
}
