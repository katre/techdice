package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A static seed makes the rolls deterministic.
// Seed: 5 (chosen to get early duplicates)
// Sequence: 1 5 2 5 4

func TestRollOnce(t *testing.T) {
	roller := New(5)
	require.NotNil(t, roller)
	seen := make(map[int]bool)
	for i := 0; i < 100; i++ {
		value := roller.rollOne()
		assert.Greater(t, value, 0)
		assert.Less(t, value, 7)
		seen[value] = true
	}
	for i := 1; i <= 6; i++ {
		assert.True(t, seen[i])
	}
	assert.Equal(t, 6, len(seen))
}

func TestRoll(t *testing.T) {
	check := func(expectedDice []int, expectedScore string, verb, push, mana, hurt int) {
		roller := New(5)
		result := roller.Roll(verb, push, mana, hurt)
		assert.Equal(t, expectedDice, result.RemainingDice)
		assert.Equal(t, expectedScore, result.Score)
	}

	check([]int{1}, "1", 1, 0, 0, 0)
	check([]int{1, 5}, "5", 1, 1, 0, 0)
	check([]int{1, 5}, "5", 1, 0, 1, 0)
	// get a duplicate
	check([]int{1, 5, 2, 5}, "5.1", 4, 0, 0, 0)
	check([]int{1, 5, 2, 5, 4}, "5.1", 3, 2, 0, 0)
	check([]int{1}, "1", 2, 0, 0, 2)
}

func TestDescribeResult(t *testing.T) {
	check := func(expected string, verb, push, mana, hurt []int) {
		r := NewResult(verb, push, mana, hurt)
		assert.Equal(t, expected, r.Describe())
	}

	check("2 [2]", []int{2}, []int{}, []int{}, []int{})
	check("3 [2] push: [3]", []int{2}, []int{3}, []int{}, []int{})
	check("3 [2] mana: [3]", []int{2}, []int{}, []int{3}, []int{})
	check("3.1 [2 3] push: [3]", []int{2, 3}, []int{3}, []int{}, []int{})
	check("3.1 [2 3] push: [3] mana: [2]", []int{2, 3}, []int{3}, []int{2}, []int{})
	check("2 [2 3] hurt: [3]", []int{2, 3}, []int{}, []int{}, []int{3})
	check("2.1 [2 3] push: [2] mana: [1] hurt: [3]", []int{2, 3}, []int{2}, []int{1}, []int{3})
}
