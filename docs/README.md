# CCL Documentation

This folder is written for people who want to build and use a CLI, not just for people who want to read the source code.

The docs focus on two things: how to ship a working command quickly, and how CCL behaves when a user types input at the terminal.

## Best Reading Order

1. [Getting Started](getting-started.md) for the first working command.
2. [Examples](examples.md) for a fuller app-shaped sample.
3. [API Reference](api-reference.md) for the public types and methods.
4. [Error Handling](errors.md) to handle failures cleanly.
5. [Architecture](architecture.md) for the mental model of the library.

## What You Get

- Command trees with subcommands
- Typed flags and positional arguments
- A validation hook before execution
- Built-in help/manual output
- Typed getters on `ExecContext`

## If Something Feels Unclear

Read [Getting Started](getting-started.md) first. It explains the library in the same order you will usually build a command: define the command, add input, read values, then handle errors.
