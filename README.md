# CCL

CCL is a small Go library for building command-line tools by describing commands in code.

You define a command tree, add typed flags and positional arguments, then let CCL parse the user input and hand the values to your action. It is designed to stay simple enough to read at a glance while still supporting nested commands and generated help text.

## Start Here

If you are new to the library, read these in order:

1. [Getting Started](docs/getting-started.md)
2. [Examples](docs/examples.md)
3. [API Reference](docs/api-reference.md)
4. [Error Handling](docs/errors.md)

## What CCL Does For You

- Reads `os.Args` and matches them to your commands
- Binds flags and positional arguments into typed values
- Gives you a validator hook before the action runs
- Prints a built-in manual for the current command
- Returns structured CLI errors when something is wrong

## Mental Model

Think of a CCL app as four steps:

1. Describe the command tree.
2. Let CCL parse the user input.
3. Read the parsed values from `ExecContext`.
4. Run your action.

## Small Example

```go
app := &ccl.Command{
    Name:        "demo",
    Usage:       "demo [command]",
    Description: "A sample CLI",
    Action: func(ctx *ccl.ExecContext) error {
        return ctx.Command.PrintManuel()
    },
}
```

This is enough to create a root command that prints its own manual when no subcommand is selected.

## Documentation

- [Getting Started](docs/getting-started.md)
- [Examples](docs/examples.md)
- [API Reference](docs/api-reference.md)
- [Error Handling](docs/errors.md)
- [Architecture](docs/architecture.md)
