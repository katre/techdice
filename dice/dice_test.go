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
		assert.Equal(t, expectedDice, result.RemainingDice)
		assert.Equal(t, expectedScore, result.Score)
	}

	check([]int{2}, "2", 1, 0, 0)
	check([]int{2, 3}, "3", 1, 1, 0)
	check([]int{2, 3, 3}, "3.1", 3, 0, 0)
	check([]int{2, 3, 3, 5, 2}, "5", 3, 2, 0)
	check([]int{2}, "2", 2, 0, 1)
}

func TestDescribeResult(t *testing.T) {
	check := func(expected string, verb, push, hurt []int) {
		r := NewResult(verb, push, hurt)
		assert.Equal(t, expected, r.Describe())
	}

	check("2 [2]", []int{2}, []int{}, []int{})
	check("3 [2] push: [3]", []int{2}, []int{3}, []int{})
	check("3.1 [2 3] push: [3]", []int{2, 3}, []int{3}, []int{})
	check("2 [2 3] hurt: [3]", []int{2, 3}, []int{}, []int{3})
}
