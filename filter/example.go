package filter

type Child struct {
	CName string `json:"c_name,select(all|2|struct)"`
	CAge  int    `json:"c_age,select(all|struct)"`
}
type Users struct {
	Name   string `json:"name,select(all|1)"`
	Age    int    `json:"age,select(all|2)"`
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

type Example struct {
	Int        int         `json:"int,select(all|intAll)"`
	Int8       int8        `json:"int8,select(all|intAll)"`
	Int16      int16       `json:"int16,select(all|intAll)"`
	Int32      int32       `json:"int32,select(all|intAll)"`
	Int64      int64       `json:"int64,select(all|intAll)"`
	IntP       *int        `json:"int_p,select(all|intAll)"`
	Int8P      *int8       `json:"int8_p,select(all|intAll)"`
	Int16P     *int16      `json:"int16_p,select(all|intAll)"`
	Int32P     *int32      `json:"int32_p,select(all|intAll)"`
	Int64P     *int64      `json:"int64_p,select(all|intAll)"`
	UInt       uint        `json:"u_int,select(all)"`
	UInt8      uint8       `json:"u_int8,select(all)"`
	UInt16     uint16      `json:"u_int16,select(all)"`
	UInt32     uint32      `json:"u_int32,select(all)"`
	UInt64     uint64      `json:"u_int64,select(all)"`
	UIntP      *uint       `json:"u_intP,select(all)"`
	UInt8P     *uint8      `json:"u_int_8_p,select(all)"`
	UInt16P    *uint16     `json:"u_int_16_p,select(all)"`
	UInt32P    *uint32     `json:"u_int_32_p,select(all)"`
	UInt64P    *uint64     `json:"u_int_64_p,select(all)"`
	Float64    float64     `json:"float64,select(all)"`
	Float64P   *float64    `json:"float_64_p,select(all)"`
	Float32    float32     `json:"float32,select(all)"`
	Float32P   *float32    `json:"float_32_p,select(all)"`
	Bool       bool        `json:"bool,select(all)"`
	BoolP      *bool       `json:"bool_p,select(all)"`
	Byte       byte        `json:"byte,select(all)"`
	ByteP      *byte       `json:"byte_p,select(all)"`
	String     string      `json:"string,select(all)"`
	StringP    *string     `json:"string_p,select(all)"`
	Interface  interface{} `json:"interface,select(all)"`
	InterfaceP interface{} `json:"interface_p,select(all)"`
	Struct     struct{}    `json:"struct,select(all)"`
	StructP    *struct{}   `json:"struct_p,select(all)"`
	Structs    Users       `json:"structs,select(all|struct)"`

	StructsP *UserP `json:"structs_p,select(all|struct)"`
}
