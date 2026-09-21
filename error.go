package ccl

import (
	"fmt"
)

type ErrorKind int

const (
	ErrorUnknownOption ErrorKind = iota
	ErrorUnknownArg
	// type mismatch of the provided values
	ErrorArgValueIncorrectType
	ErrorFlagValueIncorrectType
	// unknown command
	ErrorUnknownCmd

	// dev errors for the programmer
	// names are not unique
	ErrorArgNameNotUnique
	ErrorFlagNameNotUnique
	// command name is empty
	ErrorNoNameCommand
	// action command not found
	ErrorNoActionOnCommand
)

type ValueCtx struct {
	Value      string
	ActualType string
	ValueType  string
}

type CliError struct {
	Kind  ErrorKind
	Name  string
	Cause string

	ValueCtx
}

func (ce *CliError) ErrorTitle() string {
	switch ce.Kind {
	case ErrorArgNameNotUnique:
		return "Argument name not unique"

	case ErrorFlagNameNotUnique:
		return "Flag name not unique"

	case ErrorUnknownCmd:
		return "Command unknown"

	case ErrorUnknownOption:
		return "Flag unknown"

	case ErrorUnknownArg:
		return "Argument unknown"

	case ErrorArgValueIncorrectType:
		return "Incorrect argument value"

	case ErrorFlagValueIncorrectType:
		return "Incorrect flag value"

	case ErrorNoNameCommand:
		return "No command name found"

	case ErrorNoActionOnCommand:
		return "No action found"

	default:
		return "internal CLI ERROR"
	}
}

func (ce *CliError) Error() string {
	switch ce.Kind {
	case ErrorArgNameNotUnique:
		return fmt.Sprintf("argument named: %v needs to be unique", ce.Name)

	case ErrorFlagNameNotUnique:
		return fmt.Sprintf("flag named: %v needs to be unique", ce.Name)

	case ErrorUnknownCmd:
		return fmt.Sprintf("command: %v isn't known", ce.Name)

	case ErrorUnknownOption:
		return fmt.Sprintf("flag: %v isn't known", ce.Name)

	case ErrorUnknownArg:
		return fmt.Sprintf("arg: %v isn't known", ce.Name)

	case ErrorArgValueIncorrectType:
		return fmt.Sprintf("argument %v is of type %v, but received %v", ce.Name, ce.ActualType, ce.ValueType)

	case ErrorFlagValueIncorrectType:
		return fmt.Sprintf("flag %v is of type %v, but received %v", ce.Name, ce.ActualType, ce.ValueType)

	case ErrorNoNameCommand:
		return "no command name found, name is required"

	case ErrorNoActionOnCommand:
		return fmt.Sprintf("no action found on the %v command", ce.Name)

	default:
		return "internal CLI ERROR"
	}
}
