package parser

import (
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

func (p *Parser) Roll(input string) string {
	s := lexer.NewScanner(strings.NewReader(input))
	if s == nil {
		// panic
	}
	return ""
}
