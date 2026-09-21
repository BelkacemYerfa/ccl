package ccl

type ExecContext struct {
	Command *Command
	// although this fields mainly used internally, we provide them to the programmer, for probable cases where they need them
	Args  []Arg
	Flags map[string]Flag
}

func (execCtx *ExecContext) arg[T any](name string) T {
	var v any
	for _, arg := range execCtx.Args {
		if arg.GetName() == name {
			v = arg.GetValue()
		}
	}

	if a, ok := v.(T); ok {
		return a
	}

	var zeroed T
	return zeroed
}

func (execCtx *ExecContext) flag[T any](name string) T {
	var v any

	if flag, exist := execCtx.Flags[name]; exist {
		v = flag.GetValue()
	} else {
		// if it is an abbreviation
		flag := execCtx.Command.findFlag(name)
		if flag != nil {
			v = execCtx.Flags[flag.GetName()].GetValue()
		}
	}

	// handle case of flag wasn't found

	if a, ok := v.(T); ok {
		return a
	}

	var zeroed T
	return zeroed
}

type Invocation struct {
	Command *Command

	Args  []Arg
	Flags map[string]Flag
}

// ctx methods to get a flag or an arg value safely with type guaranty
func (ectx *ExecContext) GetIntFlag(name string) int {
	return ectx.flag[int](name)
}

func (ectx *ExecContext) GetInt8Flag(name string) int8 {
	return ectx.flag[int8](name)
}

func (ectx *ExecContext) GetInt16Flag(name string) int16 {
	return ectx.flag[int16](name)
}

func (ectx *ExecContext) GetInt32Flag(name string) int32 {
	return ectx.flag[int32](name)
}

func (ectx *ExecContext) GetInt64Flag(name string) int64 {
	return ectx.flag[int64](name)
}

func (ectx *ExecContext) GetUIntFlag(name string) uint {
	return ectx.flag[uint](name)
}

func (ectx *ExecContext) GetUInt8Flag(name string) uint8 {
	return ectx.flag[uint8](name)
}

func (ectx *ExecContext) GetUInt16Flag(name string) uint16 {
	return ectx.flag[uint16](name)
}

func (ectx *ExecContext) GetUInt32Flag(name string) uint32 {
	return ectx.flag[uint32](name)
}

func (ectx *ExecContext) GetUInt64Flag(name string) uint64 {
	return ectx.flag[uint64](name)
}

func (ectx *ExecContext) GetFloat32Flag(name string) float32 {
	return ectx.flag[float32](name)
}

func (ectx *ExecContext) GetFloat64Flag(name string) float64 {
	return ectx.flag[float64](name)
}

func (ectx *ExecContext) GetStringFlag(name string) string {
	return ectx.flag[string](name)
}

func (ectx *ExecContext) GetBoolFlag(name string) bool {
	return ectx.flag[bool](name)
}

func (ectx *ExecContext) GetIntArg(name string) int {
	return ectx.arg[int](name)
}

func (ectx *ExecContext) GetInt8Arg(name string) int8 {
	return ectx.arg[int8](name)
}

func (ectx *ExecContext) GetInt16Arg(name string) int16 {
	return ectx.arg[int16](name)
}

func (ectx *ExecContext) GetInt32Arg(name string) int32 {
	return ectx.arg[int32](name)
}

func (ectx *ExecContext) GetInt64Arg(name string) int64 {
	return ectx.arg[int64](name)
}

func (ectx *ExecContext) GetUIntArg(name string) uint {
	return ectx.arg[uint](name)
}

func (ectx *ExecContext) GetUInt8Arg(name string) uint8 {
	return ectx.arg[uint8](name)
}

func (ectx *ExecContext) GetUInt16Arg(name string) uint16 {
	return ectx.arg[uint16](name)
}

func (ectx *ExecContext) GetUInt32Arg(name string) uint32 {
	return ectx.arg[uint32](name)
}

func (ectx *ExecContext) GetUInt6Arg(name string) uint64 {
	return ectx.arg[uint64](name)
}

func (ectx *ExecContext) GetFloat32Arg(name string) float32 {
	return ectx.arg[float32](name)
}

func (ectx *ExecContext) GetFloat64Arg(name string) float64 {
	return ectx.arg[float64](name)
}

func (ectx *ExecContext) GetStringArg(name string) string {
	return ectx.arg[string](name)
}

func (ectx *ExecContext) GetBoolArg(name string) bool {
	return ectx.arg[bool](name)
}
