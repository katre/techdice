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
	// A static seed makes the rolls deterministic.
	// Sequence: 2 3 3 5 2
	roller := New(1)
	assert.Equal(t, "2", roller.Roll(1, 0, 0))

	roller = New(1)
	assert.Equal(t, "3", roller.Roll(1, 1, 0))

	roller = New(1)
	assert.Equal(t, "3.1", roller.Roll(3, 0, 0))

	roller = New(1)
	assert.Equal(t, "5", roller.Roll(3, 2, 0))

	roller = New(1)
	assert.Equal(t, "2", roller.Roll(2, 0, 1))
}
