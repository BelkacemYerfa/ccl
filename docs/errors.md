# Error Handling

Most CCL errors fall into one of two groups:

- setup mistakes while you are defining commands
- user input mistakes while someone is running the CLI

CCL reports both through `*ccl.CliError`.

## Common Error Types

| Kind                          | What it usually means                                                    | Typical fix                                             |
| ----------------------------- | ------------------------------------------------------------------------ | ------------------------------------------------------- |
| `ErrorUnknownOption`          | The user typed a flag your command does not define.                      | Check the flag name and abbreviation.                   |
| `ErrorUnknownArg`             | The user typed a positional value that does not match the command shape. | Check the command usage and argument order.             |
| `ErrorArgValueIncorrectType`  | A positional value could not be parsed into the type you declared.       | Make sure the value matches the expected type.          |
| `ErrorFlagValueIncorrectType` | A flag value could not be parsed into the type you declared.             | Make sure the flag value matches the expected type.     |
| `ErrorUnknownCmd`             | The user typed a command that does not exist.                            | Check the subcommand name.                              |
| `ErrorArgNameNotUnique`       | Two arguments on the same command use the same name.                     | Rename one of the arguments.                            |
| `ErrorFlagNameNotUnique`      | Two flags on the same command use the same name.                         | Rename one of the flags.                                |
| `ErrorNoNameCommand`          | A command was created without a name.                                    | Set `Command.Name`.                                     |
| `ErrorNoActionOnCommand`      | A command was run without an action.                                     | Add an `Action` or make the command a help-only parent. |

## What The Error Methods Give You

`CliError.ErrorTitle()` returns a short category such as `Flag unknown` or `No action found`.

`CliError.Error()` returns the full message, including the failing name and, when relevant, the expected and received types.

Examples:

- `flag: --dir isn't known`
- `argument name is of type string, but received int`

## How To Handle Errors In `main`

Use `errors.As` so you can treat CCL errors differently from unexpected failures.

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

## A Good Habit

When a command fails in a way the user can fix, print the error message exactly as returned by CCL. When the failure is unexpected, keep that message separate so you can debug it later.
