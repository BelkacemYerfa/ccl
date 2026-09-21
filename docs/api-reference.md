# API Reference

This page is a practical guide to the types you will actually use when building a CLI with CCL.

## Command

Use `ccl.Command` to describe one command or subcommand.

### When To Use It

Use a `Command` whenever you want to:

- define the root command for an app
- add a subcommand such as `run`, `build`, or `version`
- attach flags, positional arguments, or both
- run custom validation before the action starts

### Command Fields

- `Name string`: the command name shown and matched by the parser.
- `Usage string`: the usage line shown in help/manual output.
- `Description string`: the short text that explains what the command does.
- `Args []Arg`: positional arguments, in the order the user types them.
- `Flags []Flag`: named options for the command.
- `Commands []*Command`: child commands.
- `Validator func(ctx *ExecContext) error`: optional pre-run validation.
- `Action func(ctx *ExecContext) error`: the function that runs when the command is selected.

### Methods

- `Run(osArgs []string) error`: parse the command line and execute the correct command.
- `PrintManuel() error`: print the generated manual to standard output. The method name is spelled `Manuel` in the current API.

### Child Command Behavior

If a command has child commands, `Run` also adds:

- a `help` subcommand
- a `--help` flag with the `-h` abbreviation on the parent and each child command

That means users can ask for help without you writing a separate help command by hand.

## Args

Use an `Arg` when the user should type a value without a flag name.

### Available Argument Types

- `IntArg`
- `Int8Arg`
- `Int16Arg`
- `Int32Arg`
- `Int64Arg`
- `UIntArg`
- `UInt8Arg`
- `UInt16Arg`
- `UInt32Arg`
- `UInt64Arg`
- `Float32Arg`
- `Float64Arg`
- `StringArg`
- `BoolArg`

### Argument Base Fields

All argument helpers use `ArgBase[T]` under the hood.

- `Name string`
- `Description string`
- `DefaultValue T`

### Argument Rule

Use positional arguments for required input that is naturally ordered, such as a file path or a target name.

## Flags

Use a `Flag` when the user should provide a named option.

### Available Flag Types

- `IntFlag`
- `Int8Flag`
- `Int16Flag`
- `Int32Flag`
- `Int64Flag`
- `UIntFlag`
- `UInt8Flag`
- `UInt16Flag`
- `UInt32Flag`
- `UInt64Flag`
- `Float32Flag`
- `Float64Flag`
- `StringFlag`
- `BoolFlag`

### Flag Base Fields

All flag helpers use `FlagBase[T]` under the hood.

- `Name string`: canonical flag name such as `--output`.
- `Abbreviation string`: short form such as `-o`.
- `Description string`: help text shown to the user.
- `DefaultValue T`: fallback value when the user does not set the flag.

### Flag Rule

Use flags for optional settings, toggles, and configuration values that should not be positional.

## ExecContext

`ExecContext` is the object you receive in validators and actions.

### ExecContext Fields

- `Command *Command`: the command being executed.
- `Args []Arg`: the parsed positional arguments.
- `Flags map[string]Flag`: the parsed flags.

### Typed Getters

Use the getter that matches the type you declared.

- `GetIntFlag`, `GetInt8Flag`, `GetInt16Flag`, `GetInt32Flag`, `GetInt64Flag`
- `GetUIntFlag`, `GetUInt8Flag`, `GetUInt16Flag`, `GetUInt32Flag`, `GetUInt64Flag`
- `GetFloat32Flag`, `GetFloat64Flag`
- `GetStringFlag`, `GetBoolFlag`
- `GetIntArg`, `GetInt8Arg`, `GetInt16Arg`, `GetInt32Arg`, `GetInt64Arg`
- `GetUIntArg`, `GetUInt8Arg`, `GetUInt16Arg`, `GetUInt32Arg`, `GetUInt6Arg`
- `GetFloat32Arg`, `GetFloat64Arg`
- `GetStringArg`, `GetBoolArg`

The current API exposes `GetUInt6Arg`, which returns `uint64`. Keep that spelling in mind when writing code against the library.

### Accessor Behavior

The getters return the zero value of the requested type when the name is missing or the stored value does not match the requested type.

## Errors

Use `CliError` when you want to understand why a command failed.

Related types:

- `ErrorKind`
- `ValueCtx`

Read [Error Handling](errors.md) for a plain-language explanation of the error kinds and how to handle them in `main`.
