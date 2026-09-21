package ccl

import (
	"reflect"
	"strconv"
)

type Flag interface {
	GetName() string
	GetAbbreviation() string
	ParseValue(v string) (any, error)
	GetValue() any
	GetDefaultValue() any
	SetValue(v any)
	Type() string
}

// @Note: probably add a required field for required flags and force them to exist
type FlagBase[T any] struct {
	Name         string // e.g --output
	Abbreviation string // e.g -o
	Description  string
	DefaultValue T // default value if provided
	value        T // value of the flag
}

func (f *FlagBase[T]) GetName() string {
	return f.Name
}

func (f *FlagBase[T]) GetAbbreviation() string {
	return f.Abbreviation
}

func (a *FlagBase[T]) Type() string {
	return reflect.TypeFor[T]().String()
}

func (f *FlagBase[T]) GetValue() any {
	return f.value
}

func (f *FlagBase[T]) GetDefaultValue() any {
	return f.DefaultValue
}

func (f *FlagBase[T]) SetValue(v any) {
	f.value = v.(T)
}

func (f *FlagBase[T]) ParseValue(v string) (any, error) {
	return f.ParseValue(v)
}

// this allows us to create flag helpers for this cases, to insure the type of the value that we gonna receive when done at the parsing level, to build the command context and then send it to the runner that will exec action based on that

// @Note: probably add support cases such as slices of diff types
type (
	// int base
	IntFlag   struct{ FlagBase[int] }
	Int8Flag  struct{ FlagBase[int8] }
	Int16Flag struct{ FlagBase[int16] }
	Int32Flag struct{ FlagBase[int32] }
	Int64Flag struct{ FlagBase[int64] }

	// uint base
	UIntFlag   struct{ FlagBase[uint] }
	UInt8Flag  struct{ FlagBase[uint8] }
	UInt16Flag struct{ FlagBase[uint16] }
	UInt32Flag struct{ FlagBase[uint32] }
	UInt64Flag struct{ FlagBase[uint64] }

	// floats
	Float32Flag struct{ FlagBase[float32] }
	Float64Flag struct{ FlagBase[float64] }

	// string
	StringFlag struct{ FlagBase[string] }

	// bool
	BoolFlag struct{ FlagBase[bool] }
)

func (f *StringFlag) ParseValue(v string) (any, error) {
	return v, nil
}

func (f *IntFlag) ParseValue(v string) (any, error) {
	return strconv.Atoi(v)
}

func (f *Int8Flag) ParseValue(v string) (any, error) {
	return strconv.ParseInt(v, 10, 8)
}

func (f *Int16Flag) ParseValue(v string) (any, error) {
	return strconv.ParseInt(v, 10, 16)
}

func (f *Int32Flag) ParseValue(v string) (any, error) {
	return strconv.ParseInt(v, 10, 32)
}

func (f *Int64Flag) ParseValue(v string) (any, error) {
	return strconv.ParseInt(v, 10, 64)
}

func (f *UIntFlag) ParseValue(v string) (any, error) {
	a, err := strconv.ParseUint(v, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(a), err
}

func (f *UInt8Flag) ParseValue(v string) (any, error) {
	return strconv.ParseUint(v, 10, 8)
}

func (f *UInt16Flag) ParseValue(v string) (any, error) {
	return strconv.ParseUint(v, 10, 16)
}

func (f *UInt32Flag) ParseValue(v string) (any, error) {
	return strconv.ParseUint(v, 10, 32)
}

func (f *UInt64Flag) ParseValue(v string) (any, error) {
	return strconv.ParseUint(v, 10, 64)
}

func (f *Float32Flag) ParseValue(v string) (any, error) {
	return strconv.ParseFloat(v, 32)
}

func (f *Float64Flag) ParseValue(v string) (any, error) {
	return strconv.ParseFloat(v, 64)
}

func (f *BoolFlag) ParseValue(v string) (any, error) {
	return strconv.ParseBool(v)
}
