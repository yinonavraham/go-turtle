package parser

import (
	"fmt"
	"github.com/yinonavraham/go-turtle/pkg/lexer"
	"github.com/yinonavraham/go-turtle/pkg/parser/ast"
	"io"
	"strconv"
	"strings"
)

// Parse the text from the provided reader into a LOGO script AST.
//
// Returns an error in case parsing (or lexing) failed.
func Parse(r io.Reader) (ast.Script, error) {
	p := parser{}
	return p.Parse(r)
}

var builtinCommandAliases = map[string]string{
	"FD": ast.CommandForward,
	"BK": ast.CommandBack,
	"RT": ast.CommandRight,
	"LT": ast.CommandLeft,
	//"": ast.CommandHome,
	//"": ast.CommandRepeat,
	"PU": ast.CommandPenUp,
	"PD": ast.CommandPenDown,
	"ST": ast.CommandShowTurtle,
	"HT": ast.CommandHideTurtle,
	"CL": ast.CommandClean,
	"CS": ast.CommandClearScreen,
}

type parser struct {
	tokens     []lexer.Token
	current    int
	cmdParsers map[string]cmdParseFunc
}

type cmdParseFunc func(cmds *[]ast.Command) error

func (p *parser) Parse(r io.Reader) (ast.Script, error) {
	tokens, err := lexer.Lex(r)
	if err != nil {
		return ast.Script{}, err
	}
	script := ast.Script{}
	p.init(tokens)
	if err := p.parseCommands(&script.Commands); err != nil {
		return ast.Script{}, err
	}
	return script, nil
}

func (p *parser) init(tokens []lexer.Token) {
	p.tokens = tokens
	p.current = 0
	p.cmdParsers = map[string]cmdParseFunc{
		ast.CommandForward: p.parseForward,
		ast.CommandBack:    p.parseBack,
	}
}

func (p *parser) hasNext() bool {
	return p.current < len(p.tokens)
}

func (p *parser) parseCommands(cmds *[]ast.Command) error {
	for p.hasNext() {
		if err := p.parseCommand(cmds); err != nil {
			return err
		}
	}
	return nil
}

func (p *parser) parseCommand(cmds *[]ast.Command) error {
	token := p.tokens[p.current]
	cmdName := strings.ToUpper(token.Value)
	if aliasedCmd, ok := builtinCommandAliases[cmdName]; ok {
		cmdName = aliasedCmd
	}
	parse, ok := p.cmdParsers[cmdName]
	if !ok {
		return fmt.Errorf("unknown command: %s (line %d, column %d)", token.Value, token.Position.Line, token.Position.Column)
	}
	if err := parse(cmds); err != nil {
		return err
	}
	return nil
}

func (p *parser) parseForward(cmds *[]ast.Command) error {
	token := p.tokens[p.current]
	p.current++
	expr, err := p.parseExpression()
	if err != nil {
		return err
	}
	cmd := ast.Forward{
		Steps: expr,
		ItemBase: ast.ItemBase{
			StartPosition: token.Position,
			EndPosition:   toEndPosition(p.tokens[p.current-1]),
		},
	}
	*cmds = append(*cmds, cmd)
	return nil
}

func (p *parser) parseBack(cmds *[]ast.Command) error {
	token := p.tokens[p.current]
	p.current++
	expr, err := p.parseExpression()
	if err != nil {
		return err
	}
	cmd := ast.Back{
		Steps: expr,
		ItemBase: ast.ItemBase{
			StartPosition: token.Position,
			EndPosition:   toEndPosition(p.tokens[p.current-1]),
		},
	}
	*cmds = append(*cmds, cmd)
	return nil
}

func toEndPosition(token lexer.Token) lexer.Position {
	return lexer.Position{Line: token.Position.Line, Column: token.Position.Column + len(token.Value)}
}

func (p *parser) parseExpression() (ast.Expr, error) {
	token := p.tokens[p.current]
	p.current++
	var expr ast.Expr
	switch {
	case strings.HasPrefix(token.Value, ":"):
		expr = ast.VariableExpr{
			Name: token.Value,
			ItemBase: ast.ItemBase{
				StartPosition: token.Position,
				EndPosition:   toEndPosition(token),
			},
		}
	default:
		intVal, err := strconv.Atoi(token.Value)
		if err != nil {
			return nil, fmt.Errorf("integer value expected: %s (line %d, column %d): %s", token.Value, token.Position.Line, token.Position.Column, err)
		}
		expr = ast.IntegerExpr{
			Value: intVal,
			ItemBase: ast.ItemBase{
				StartPosition: token.Position,
				EndPosition:   toEndPosition(token),
			},
		}
	}
	return expr, nil
}
