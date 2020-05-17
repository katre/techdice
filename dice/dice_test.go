package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRollOnce(t *testing.T) {
	roller := New(1)
	require.NotNil(t, roller)
	for i := 0; i < 100; i++ {
		value := roller.rollOne()
		assert.Greater(t, value, 0)
		assert.Less(t, value, 7)
	}
}

func TestRoll(t *testing.T) {
	check := func(expectedDice []int, expectedScore string, verb, push, hurt int) {
		// A static seed makes the rolls deterministic.
		// Sequence: 2 3 3 5 2
		roller := New(1)
		result := roller.Roll(verb, push, hurt)
		assert.Equal(t, expectedDice, result.Dice)
		assert.Equal(t, expectedScore, result.Score)
	}

	check([]int{2}, "2", 1, 0, 0)
	check([]int{2, 3}, "3", 1, 1, 0)
	check([]int{2, 3, 3}, "3.1", 3, 0, 0)
	check([]int{2, 3, 3, 5, 2}, "5", 3, 2, 0)
	check([]int{2}, "2", 2, 0, 1)
}
