package ccl

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const (
	manuelTemplate = `{{.Description}}

Usage:
{{padRight "" 0}}{{.Usage}}
{{if .Commands}}
Commands:
{{range .Commands}}
{{padRight "" 0}}{{.Name}}{{padName .Name 0}}{{.Description}}
{{- end}}
{{- end}}
{{- if .Args}}
Args:
{{range .Args}}
{{padRight "" 0}}{{.Name}}{{padName .Name 0}}{{.Description}}
{{- end}}
{{- end}}
{{if .Flags}}
Flags:
{{range .Flags}}
{{padRight .Abbreviation 2}}, {{.Name}}{{padName .Name 0}}{{.Description}}
{{end}}
{{end}}`
)

// @Todo:
// inject to have -h by default on every command
// have an implicitly help command injected with ccl
// @note: use go text/template for this
type Command struct {
	Name        string
	Usage       string
	Description string

	Args     []Arg
	Flags    []Flag
	Commands []*Command

	Validator func(ctx *ExecContext) error
	Action    func(ctx *ExecContext) error
}

func (cmd *Command) findFlag(name string) Flag {
	for _, f := range cmd.Flags {
		if f.GetName() == name || f.GetAbbreviation() == name {
			return f
		}
	}
	return nil
}

func (cmd *Command) findArg(name string) Arg {
	for _, a := range cmd.Args {
		if a.GetName() == name {
			return a
		}
	}
	return nil
}

func (cmd *Command) findSubCommand(name string) *Command {
	for _, scmd := range cmd.Commands {
		if scmd.Name == name {
			return scmd
		}
	}
	return nil
}

func (cmd *Command) bind(nodes []Node) (*Invocation, error) {
	inv := &Invocation{
		Command: cmd,
		Args:    []Arg{},
		Flags:   make(map[string]Flag, len(cmd.Flags)),
	}

	for iter, node := range nodes {
		switch n := node.(type) {
		case *OptionNode:
			flag := cmd.findFlag(n.Name)
			if flag == nil {
				return nil, &CliError{Kind: ErrorUnknownOption, Name: n.Name}
			}

			var value any

			if len(n.Values) == 0 {
				// booleans can be represented by flag itself, that means it is a true value
				// also the programmer can specify the value explicitly too

				_, ok := flag.(*BoolFlag)
				if ok {
					// existence of the flag means ok the value will be true
					flag.SetValue(true)
					inv.Flags[flag.GetName()] = flag
					continue
				} else {
					value = flag.GetValue()
				}
			}

			// make sure to use default value in case there is no provided value(s)

			raw := n.Values[0]

			value, err := flag.ParseValue(raw)
			if err != nil {
				return nil, &CliError{
					Name:       flag.GetName(),
					Kind:       ErrorFlagValueIncorrectType,
					Value:      raw,
					ActualType: flag.Type(),
					ValueType:  fmt.Sprintf("%T", raw),
				}
			}

			// handle the case to set the value
			flag.SetValue(value)
			inv.Flags[flag.GetName()] = flag

		case *ArgumentNode:
			// don't search for the arg name cause it ain't provided, just the name is provided
			// so the better way to get the positional arg here, is to go with the slice order and get the arg by the idx
			if len(cmd.Args) == 0 {
				continue
			}

			arg := cmd.Args[iter]
			if arg == nil {
				return nil, &CliError{
					Name: n.Name,
					Kind: ErrorUnknownArg,
				}
			}

			var value any

			if len(n.Values) == 0 {
				// booleans can be represented by flag itself, that means it is a true value
				// also the programmer can specify the value explicitly too

				_, ok := arg.(*BoolArg)
				if ok {
					// existence of the flag means ok the value will be true
					arg.SetValue(true)
					inv.Args = append(inv.Args, arg)
					continue
				} else {
					value = arg.GetValue()
				}
			}

			// make sure to use default value in case there is no provided value(s)
			// handle the case where the type is a slice
			raw := n.Name

			value, err := arg.ParseValue(raw)
			if err != nil {
				return nil, &CliError{
					Name:       arg.GetName(),
					Kind:       ErrorArgValueIncorrectType,
					Value:      raw,
					ActualType: arg.Type(),
					ValueType:  fmt.Sprintf("%T", raw),
				}
			}

			// handle the case to set the value
			arg.SetValue(value)
			inv.Args = append(inv.Args, arg)
		}
	}

	for _, flag := range cmd.Flags {
		if _, exists := inv.Flags[flag.GetName()]; exists {
			continue
		}

		flag.SetValue(flag.GetDefaultValue())
		inv.Flags[flag.GetName()] = flag
	}

outer:
	for _, arg := range cmd.Args {
		// search inside the inv.Args
		for _, a := range inv.Args {
			if a.GetName() == arg.GetName() {
				break outer
			}
		}

		arg.SetValue(arg.GetDefaultValue())
		inv.Args = append(inv.Args, arg)
	}

	return inv, nil
}

func (cmd *Command) uniqueArgName() error {
	seen := cmd.Args[0]

	for iter, arg := range cmd.Args[1:] {
		for _, narg := range cmd.Args[iter:] {
			if narg.GetName() == seen.GetName() {
				return &CliError{
					Name: seen.GetName(),
					Kind: ErrorArgNameNotUnique,
				}
			}
		}

		seen = arg
	}

	return nil
}

func (cmd *Command) uniqueFlagName() error {
	seen := cmd.Flags[0]

	for iter, flag := range cmd.Flags[1:] {
		for _, nflag := range cmd.Flags[iter:] {
			if nflag.GetName() == seen.GetName() {
				return &CliError{
					Name: seen.GetName(),
					Kind: ErrorFlagNameNotUnique,
				}
			}
		}

		seen = flag
	}

	return nil

}

func (cmd *Command) uniqueNaming() error {
	if len(cmd.Args) > 0 {
		return cmd.uniqueArgName()
	}

	if len(cmd.Flags) > 0 {
		return cmd.uniqueFlagName()
	}

	return nil
}

// this validators are regarding checking and making sure that the programmer is using the api and providing what is needed is the correct manner
func (cmd *Command) customValidators() error {
	// require name of the command to exist
	if len(cmd.Name) == 0 {
		return &CliError{
			Kind: ErrorNoNameCommand,
		}
	}
	// require the name to be unique
	if err := cmd.uniqueNaming(); err != nil {
		return err
	}

	return nil
}

func (cmd *Command) PrintManuel() error {
	manuelTemplateFuncs := template.FuncMap{
		"padRight": func(abbv string, shift int) string {
			str := strings.Builder{}
			const rightPadding = 8
			shift = rightPadding - shift
			hasAbbv := len(abbv) > 0

			if hasAbbv {
				abbvStart := 5 // default shift
				if len(cmd.Flags) == 1 {
					abbvStart = len(cmd.Flags[0].GetAbbreviation())
				} else if len(cmd.Flags) > 1 {
					abbvStart = len(cmd.Flags[len(cmd.Flags)-2].GetAbbreviation())
				}

				shift -= abbvStart
			}

			str.WriteString(strings.Repeat(" ", shift))
			str.WriteString(abbv)

			if hasAbbv {
				// the -2 is cause of the
				// one being the , (comma)
				// the other one being " " (space)
				str.WriteString(strings.Repeat(" ", rightPadding-shift-len(abbv)-2))
			}

			return str.String()
		},
		"padName": func(name string, shift int) string {
			// space to insert between name of arg/flag/command and description
			// @note: maybe we need to add command later on in case there is a long one
			max := 5
			if len(cmd.Args) == 1 {
				max = len(cmd.Args[0].GetName())
			} else if len(cmd.Args) > 1 {
				max = len(cmd.Args[len(cmd.Args)-2].GetName())
			}

			sizeOfLastFlag := max
			if len(cmd.Flags) == 1 {
				sizeOfLastFlag = len(cmd.Flags[0].GetName())
			} else if len(cmd.Flags) > 1 {
				sizeOfLastFlag = len(cmd.Flags[len(cmd.Flags)-2].GetName())
			}

			if sizeOfLastFlag > max {
				max = sizeOfLastFlag
			}

			// the +2 is because of adding 2 other char
			// one being the , (comma)
			// the other one being " " (space)
			return strings.Repeat(" ", max-len(name)+2+shift)
		},
	}

	t, err := template.New("global manuel").Funcs(manuelTemplateFuncs).Parse(manuelTemplate)
	if err != nil {
		return err
	}

	if err := t.Execute(os.Stdout, cmd); err != nil {
		return err
	}

	return nil

}

func (cmd *Command) Run(osArgs []string) error {
	// @Todo: customValidation,
	if err := cmd.customValidators(); err != nil {
		return err
	}

	isParentCommand := len(cmd.Commands) > 0
	if isParentCommand {
		// append it on the parent command if it is not already there
		cmd.Flags = append(cmd.Flags, &BoolFlag{
			Name:         "--help",
			Abbreviation: "-h",
			Description:  "prints the manuel of the command",
		})

		// implicit insertion of the help command
		cmd.Commands = append(cmd.Commands, &Command{
			Name:        "help",
			Usage:       fmt.Sprintf("%s help [argument]", cmd.Name),
			Description: "prints the command manuel",
			Args: []Arg{
				&StringArg{
					Name:        "command",
					Description: "command to print the manuel of",
				},
			},
			Action: func(ctx *ExecContext) error {
				command := ctx.GetStringArg("command")
				if len(command) == 0 {
					// global manuel
					return cmd.PrintManuel()
				}

				// print the help of the provided command if it is registered
				subCmd := cmd.findSubCommand(command)
				if subCmd == nil {
					return &CliError{
						Name: command,
						Kind: ErrorUnknownCmd,
					}
				}
				return subCmd.PrintManuel()
			},
		})
	}

	// insert it to all the command that exist in the current parent command
	for _, cmd := range cmd.Commands {
		cmd.Flags = append(cmd.Flags, &BoolFlag{
			Name:         "--help",
			Abbreviation: "-h",
			Description:  "prints the manuel of the command",
		})
	}

	l := NewCmdLexer(osArgs[1:])
	p := NewCmdParser(l)

	nodes := p.Parse()

	// checks if the current arg is a sub command
	if len(cmd.Commands) > 0 && len(osArgs) > 1 && len(nodes) > 0 {
		if _, ok := nodes[0].(*ArgumentNode); ok {
			subCmd := cmd.findSubCommand(osArgs[1])
			if subCmd == nil {
				return &CliError{
					Name: osArgs[1],
					Kind: ErrorUnknownCmd,
				}
			}
			return subCmd.Run(osArgs[1:])
		}
	}

	inv, err := cmd.bind(nodes)
	if err != nil {
		return err
	}

	execCtx := &ExecContext{
		Command: inv.Command,
		Args:    inv.Args,
		Flags:   inv.Flags,
	}

	// validation is done by the programmer, we give them exec context and they can do certain stuff if they want to
	if cmd.Validator != nil {
		if err := cmd.Validator(execCtx); err != nil {
			return err
		}
	}

	if cmd.Action == nil {
		return &CliError{
			Name: cmd.Name,
			Kind: ErrorNoActionOnCommand,
		}
	}

	if helpFlag := execCtx.GetBoolFlag("--help"); helpFlag {
		return cmd.PrintManuel()
	}

	return cmd.Action(execCtx)
}
