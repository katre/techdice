package parser

import (
	"github.com/katre/techdice/dice"
)

type Parser struct {
	roller *dice.Roller
}

func New(roller *dice.Roller) *Parser {
	return nil
}

func (p *Parser) Roll(input string) string {
	return ""
}
