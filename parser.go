package ccl

import (
	"strconv"
	"strings"
)

type CmdLexer struct {
	Cur  int
	Args []string
}

type CmdTokenKind string

type CmdToken struct {
	Token CmdTokenKind
	Text  string
}

const (
	CmdTokenIdentifier CmdTokenKind = "identifier"
	CmdTokenLongFlag   CmdTokenKind = "long_flag"
	CmdTokenShotFlag   CmdTokenKind = "short_flag"
	CmdTokenEOA        CmdTokenKind = "EOA"
)

func NewCmdLexer(args []string) *CmdLexer {
	return &CmdLexer{
		Cur:  0,
		Args: args,
	}
}

func (cl *CmdLexer) currentArg() string {
	if cl.Cur >= len(cl.Args) {
		// report an error or something
		return string(CmdTokenEOA)
	}

	return cl.Args[cl.Cur]
}

func (cl *CmdLexer) advanceArg() {
	if cl.Cur >= len(cl.Args) {
		// report an error or something
		return
	}

	cl.Cur++
}

func (cl *CmdLexer) NextCmdToken() CmdToken {
	cmdToken := CmdToken{}

	if cl.Cur >= len(cl.Args) {
		cmdToken.Token = CmdTokenEOA
		cmdToken.Text = "EOA"
		return cmdToken
	}

	arg := cl.currentArg()

	cmdToken.Text = arg

	switch {
	case strings.HasPrefix(arg, "-"):
		// check if after it a number or not
		tokWithoutMinus := strings.TrimPrefix(arg, "-")
		if _, err := strconv.ParseFloat(tokWithoutMinus, 64); err != nil {
			// means it is not a valid number
			cmdToken.Token = CmdTokenShotFlag
		} else {
			cmdToken.Token = CmdTokenIdentifier
		}

	case strings.HasPrefix(arg, "--"):
		cmdToken.Token = CmdTokenLongFlag
	default:
		cmdToken.Token = CmdTokenIdentifier
	}

	// advance
	cl.advanceArg()

	return cmdToken
}

type Node interface{}

type ArgumentNode struct {
	// here we need the rest of registered metadata hers that is registered
	Name string

	// command arguments
	Values []string
}

type OptionNode struct {
	Name string

	// flag values
	Values []string
}

type CmdParser struct {
	cl *CmdLexer

	// inner pointer for the state
	currentTok *CmdToken
	peekTok    *CmdToken
}

func NewCmdParser(cl *CmdLexer) *CmdParser {
	cmdParser := CmdParser{
		cl: cl,
	}

	// set the internal tokens
	cmdParser.nextTok()
	cmdParser.nextTok()

	return &cmdParser
}

func (cp *CmdParser) nextTok() {
	cp.currentTok = cp.peekTok
	pt := cp.cl.NextCmdToken()
	cp.peekTok = &pt
}

func (cp *CmdParser) curTokenKindIs(kind CmdTokenKind) bool {
	return cp.currentTok.Token == kind
}

func (cp *CmdParser) peekTokenKindIs(kind CmdTokenKind) bool {
	return cp.peekTok.Token == kind
}

func (cp *CmdParser) Parse() []Node {
	nodes := []Node{}

	for !cp.curTokenKindIs(CmdTokenEOA) {

		switch cp.currentTok.Token {
		case CmdTokenLongFlag, CmdTokenShotFlag:
			nodes = append(nodes, cp.parseOption())

		default:
			// we assume that this is a cmd
			nodes = append(nodes, cp.parseCommand())
		}
	}

	// means no node was parsed here
	return nodes
}

// keep parsing until the next new token
// so wall of the identifiers are going to be consumed here
func (cp *CmdParser) parseOption() *OptionNode {
	opn := OptionNode{
		Name: cp.currentTok.Text,
	}

	cp.nextTok()

	for cp.curTokenKindIs(CmdTokenIdentifier) {
		// continue collecting all of the identifier here as part of the options
		// and validation of them will happen later on in the next phases
		opn.Values = append(opn.Values, cp.currentTok.Text)
		cp.nextTok()
	}

	return &opn
}

// so wall of the identifiers are going to be consumed here
func (cp *CmdParser) parseCommand() *ArgumentNode {
	cmn := ArgumentNode{
		Name: cp.currentTok.Text,
	}

	cp.nextTok()

	for cp.curTokenKindIs(CmdTokenIdentifier) {
		// continue collecting all of the identifier here as part of the options
		// and validation of them will happen later on in the next phases
		cmn.Values = append(cmn.Values, cp.currentTok.Text)
		cp.nextTok()
	}

	return &cmn
}
