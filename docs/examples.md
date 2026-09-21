# Examples

The sample app in `examples/cli.go` shows a compiler-style CLI with a root command and two subcommands.

## What The Example Does

- `cli`: the root command that prints the manual
- `run`: accepts a `path` argument and a few options such as `--dir` and `--max-errors`
- `version`: prints the current OS and architecture

## Why This Example Is Useful

It demonstrates the parts most real tools need:

- a root command that owns the app
- a subcommand that reads both arguments and flags
- a simple read-only command with no extra input

## A Good Pattern To Copy

When you build your own CLI, keep each command focused on one job.

```go
root := &ccl.Command{
    Name:        "app",
    Usage:       "app [command]",
    Description: "Example application",
    Commands: []*ccl.Command{
        {
            Name:        "serve",
            Usage:       "app serve [options]",
            Description: "Start the service",
            Flags: []ccl.Flag{
                &ccl.StringFlag{Name: "--port", Abbreviation: "-p", DefaultValue: "8080"},
            },
            Action: func(ctx *ccl.ExecContext) error {
                fmt.Println("starting on port", ctx.GetStringFlag("--port"))
                return nil
            },
        },
    },
    Action: func(ctx *ccl.ExecContext) error {
        return ctx.Command.PrintManuel()
    },
}
```

## Write Command Docs For People, Not For The Parser

If you want your CLI to be easy to use, make sure each command has:

- a short name that makes sense at the terminal
- a `Usage` line that shows the expected shape of the command
- a description that explains the outcome, not just the syntax
- flag and argument names that match the code exactly

## A Simple Checklist

Before you ship a command, check that a new user can answer these questions from the help text alone:

1. What does this command do?
2. What do I type first?
3. Which values are required?
4. Which values are optional?
5. What happens if I make a mistake?
