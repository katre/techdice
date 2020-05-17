package bot

import (
	"github.com/katre/techdice/dice"
	"github.com/katre/techdice/parser"
)

type Bot struct {
}

func New(seed int64, botId string) *Bot {
	// Create a new roller.
	roller := dice.New(seed)
	//parser := parser.New(roller)
	parser.New(roller)

	return &Bot{}
}
