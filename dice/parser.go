package parser

import (
	"fmt"
	"strings"

	"aoc/util/lexer"
	"github.com/katre/techdice/dice"
)

type Parser struct {
	roller *dice.Roller
}

func New(roller *dice.Roller) *Parser {
	return &Parser{
		roller: roller,
	}
}

func (p *Parser) Roll(input string) (dice.Result, error) {
	s := lexer.NewScanner(strings.NewReader(input))

	// Expected format: NUM (push NUM)? (hurt NUM)?
	verb, err := s.ScanNumber("verb")
	if err != nil {
		return dice.Result{}, err
	}

	// scan until end of input
	var push, hurt int
	for {
		token := s.Scan()
		if token.TokenType == lexer.EOF {
			break
		} else if token.TokenType == lexer.IDENT && token.Literal == "push" {
			value, err := s.ScanNumber("push")
			if err != nil {
				return dice.Result{}, err
			}
			push += value
		} else if token.TokenType == lexer.IDENT && token.Literal == "hurt" {
			value, err := s.ScanNumber("hurt")
			if err != nil {
				return dice.Result{}, err
			}
			hurt += value
		} else {
			// Unknown token.
			return dice.Result{}, fmt.Errorf("Unknown token %q", token)
		}
	}

	return p.roller.Roll(verb, push, hurt), nil
}
