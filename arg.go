package ccl

import (
	"reflect"
	"strconv"
)

type Arg interface {
	GetName() string
	GetValue() any
	GetDefaultValue() any
	ParseValue(v string) (any, error)
	SetValue(v any)
	Type() string
}

type ArgBase[T any] struct {
	Name         string
	Description  string
	DefaultValue T
	value        T // value of the flag
}

func (a *ArgBase[T]) GetName() string {
	return a.Name
}

func (a *ArgBase[T]) Type() string {
	return reflect.TypeFor[T]().String()
}

func (a *ArgBase[T]) GetValue() any {
	return a.value
}

func (a *ArgBase[T]) GetDefaultValue() any {
	return a.DefaultValue
}

func (f *ArgBase[T]) SetValue(v any) {
	f.value = v.(T)
}

func (f *ArgBase[T]) ParseValue(v string) (any, error) {
	return f.ParseValue(v)
}

// this even allows us to create flag helpers for this cases, to insure the type of the value that we gonna receive when done at the parsing level,to  build the command context and then send it to the runner that will exec action based on that

// @Note: probably add the rest of the cases, to support multi flags or different types
// even support cases such as uint & int (with all bitSizes), same for floats ..etc
type (
	// int base
	IntArg   struct{ ArgBase[int] }
	Int8Arg  struct{ ArgBase[int8] }
	Int16Arg struct{ ArgBase[int16] }
	Int32Arg struct{ ArgBase[int32] }
	Int64Arg struct{ ArgBase[int64] }

	// uint base
	UIntArg   struct{ ArgBase[uint] }
	UInt8Arg  struct{ ArgBase[uint8] }
	UInt16Arg struct{ ArgBase[uint16] }
	UInt32Arg struct{ ArgBase[uint32] }
	UInt64Arg struct{ ArgBase[uint64] }

	// floats
	Float32Arg struct{ ArgBase[float32] }
	Float64Arg struct{ ArgBase[float64] }

	// string
	StringArg struct{ ArgBase[string] }

	// bool
	BoolArg struct{ ArgBase[bool] }
)

func (f *StringArg) ParseValue(v string) (any, error) {
	return v, nil
}

func (f *IntArg) ParseValue(v string) (any, error) {
	return strconv.Atoi(v)
}

func (f *Int8Arg) ParseValue(v string) (any, error) {
	return strconv.ParseInt(v, 10, 8)
}

func (f *Int16Arg) ParseValue(v string) (any, error) {
	return strconv.ParseInt(v, 10, 16)
}

func (f *Int32Arg) ParseValue(v string) (any, error) {
	return strconv.ParseInt(v, 10, 32)
}

func (f *Int64Arg) ParseValue(v string) (any, error) {
	return strconv.ParseInt(v, 10, 64)
}

func (f *UIntArg) ParseValue(v string) (any, error) {
	a, err := strconv.ParseUint(v, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(a), err
}

func (f *UInt8Arg) ParseValue(v string) (any, error) {
	return strconv.ParseUint(v, 10, 8)
}

func (f *UInt16Arg) ParseValue(v string) (any, error) {
	return strconv.ParseUint(v, 10, 16)
}

func (f *UInt32Arg) ParseValue(v string) (any, error) {
	return strconv.ParseUint(v, 10, 32)
}

func (f *UInt64Arg) ParseValue(v string) (any, error) {
	return strconv.ParseUint(v, 10, 64)
}

func (f *Float32Arg) ParseValue(v string) (any, error) {
	return strconv.ParseFloat(v, 32)
}

func (f *Float64Arg) ParseValue(v string) (any, error) {
	return strconv.ParseFloat(v, 64)
}

func (f *BoolArg) ParseValue(v string) (any, error) {
	return strconv.ParseBool(v)
}
