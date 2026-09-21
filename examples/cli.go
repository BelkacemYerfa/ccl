package cmd

import (
	"ccl"
	"errors"
	"fmt"
	"os"
	"runtime"
)

var (
	// main compiler executable
	rootCmd = &ccl.Command{
		Name:        "cli",
		Usage:       "cli [command]",
		Description: ".....",
		// append the other commands, from help to run, version ...etc
		Commands: []*ccl.Command{
			runCmd,
			versionCmd,
		},
		Action: func(ctx *ccl.ExecContext) error {
			// print the default global use in our case
			return ctx.Command.PrintManuel()
		},
	}

	runCmd = &ccl.Command{
		Name:        "run",
		Description: "Takes the filepath of a program, and executes it",
		Usage:       "cli run [argument] [options]",
		Args: []ccl.Arg{
			&ccl.StringArg{
				Name:        "path",
				Description: "target file to compile",
			},
		},
		Flags: []ccl.Flag{
			&ccl.StringFlag{
				Name:         "--dir",
				Abbreviation: "-d",
				Description:  "directory that contains programs",
			},
			&ccl.UIntFlag{
				Name:         "--max-errors",
				Abbreviation: "-mx",
				Description:  "error window to show in the terminal",
				DefaultValue: 10,
			},
			&ccl.StringFlag{
				Name:         "--diagnostic-level",
				Abbreviation: "-dl",
				Description:  "diagnostic level including [auto, full, squash] were auto is the default mode",
				DefaultValue: "auto",
			},
		},
		Action: func(ctx *ccl.ExecContext) error {
			targetFile := ctx.GetStringArg("path")
			diagnosticLevel := ctx.GetStringFlag("--diagnostic-level")
			maxErrorsCount := ctx.GetUIntFlag("--max-errors")
			fmt.Println(targetFile, diagnosticLevel, maxErrorsCount)
			return nil
		},
	}

	versionCmd = &ccl.Command{
		Name:        "version",
		Usage:       "cli version",
		Description: "Prints the current version of the compiler",
		Action: func(ctx *ccl.ExecContext) error {
			os := runtime.GOOS
			arch := runtime.GOARCH
			output := fmt.Sprintf("version %v/%v", os, arch)
			fmt.Println(output)
			return nil
		},
	}
)

func Runner() {
	if err := rootCmd.Run(os.Args); err != nil {
		// custom errors struct that is used internally
		// this cast is required to access more info if needed
		if cliErr, ok := errors.AsType[*ccl.CliError](err); ok {
			fmt.Println(cliErr)
		}

		fmt.Println(err)
	}
}
