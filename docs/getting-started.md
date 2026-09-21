# Getting Started

This page shows the shortest path from zero to a working CCL command.

## 1. Import The Library

If you are using this repository locally, import it as `ccl`.

```go
import "ccl"
```

If you publish the module elsewhere, change the import path to match that location.

## 2. Create A Root Command

Every CCL app starts with a `Command`.

```go
app := &ccl.Command{
    Name:        "tool",
    Usage:       "tool [command]",
    Description: "A sample CLI",
    Action: func(ctx *ccl.ExecContext) error {
        return ctx.Command.PrintManuel()
    },
}
```

This is your root command. It can print its own manual, run a default action, or host subcommands.

## 3. Add A Subcommand

Most real CLIs become useful once they have at least one subcommand.

```go
app := &ccl.Command{
    Name:        "demo",
    Usage:       "demo [command]",
    Description: "Demo application",
    Commands: []*ccl.Command{
        {
            Name:        "version",
            Usage:       "demo version",
            Description: "Print the application version",
            Action: func(ctx *ccl.ExecContext) error {
                fmt.Println("demo v1.0.0")
                return nil
            },
        },
    },
    Action: func(ctx *ccl.ExecContext) error {
        return ctx.Command.PrintManuel()
    },
}
```

When a command has child commands, CCL also adds built-in help support for the parent and each child.

## 4. Add Flags And Positional Arguments

Use positional arguments for required values and flags for optional options.

```go
runCmd := &ccl.Command{
    Name:        "run",
    Usage:       "tool run [argument] [options]",
    Description: "Run a target file",
    Args: []ccl.Arg{
        &ccl.StringArg{Name: "path", Description: "file to run"},
    },
    Flags: []ccl.Flag{
        &ccl.StringFlag{Name: "--dir", Abbreviation: "-d", Description: "working directory"},
        &ccl.UIntFlag{Name: "--max-errors", Abbreviation: "-m", Description: "maximum errors to show", DefaultValue: 10},
    },
}
```

## 5. Read Parsed Values In Your Action

Inside `Action`, use the typed getters on `ExecContext`.

```go
Action: func(ctx *ccl.ExecContext) error {
    path := ctx.GetStringArg("path")
    dir := ctx.GetStringFlag("--dir")
    maxErrors := ctx.GetUIntFlag("--max-errors")

    fmt.Println(path, dir, maxErrors)
    return nil
}
```

The getter you call should match the type you declared.

## 6. Handle Errors In `main`

Pass `os.Args` to `Run`, then print `CliError` values separately from unexpected errors.

```go
if err := app.Run(os.Args); err != nil {
    var cliErr *ccl.CliError
    if errors.As(err, &cliErr) {
        fmt.Fprintln(os.Stderr, cliErr.Error())
        os.Exit(2)
    }

    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}
```

## What `Run` Does

When you call `Run`, CCL:

1. Reads the command-line arguments.
2. Finds the matching command or subcommand.
3. Binds flags and positional arguments.
4. Fills in default values where needed.
5. Runs your validator, if present.
6. Runs your action.

## A Good First Rule

If a command is hard to understand from its `Name`, `Usage`, and `Description`, improve those three fields before you add more code. Those are the first lines your users will see.
