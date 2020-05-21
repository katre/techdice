package parser

import (
	"testing"

	"github.com/katre/techdice/dice"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	check := func(expectedDice []int, expectedScore string, input string) {
		// A static seed makes the rolls deterministic.
		// Sequence: 2 3 3 5 2
		parser := New(dice.New(1), "roll")
		require.NotNil(t, parser)
		result, err := parser.Roll(input)
		if assert.Nil(t, err) {
			assert.Equal(t, expectedDice, result.RemainingDice)
			assert.Equal(t, expectedScore, result.Score)
		}
	}

	check([]int{2}, "2", "roll 1")
	check([]int{2, 3}, "3", "roll 1 push 1")
	check([]int{2, 3, 3}, "3.1", "roll 2 push 1")
	check([]int{2, 3, 3, 5, 2}, "5", "roll 3 push 2")
	check([]int{2}, "2", "roll 2 hurt 1")
	check([]int{2}, "2", "roll 1 push 1 hurt 1")
}
