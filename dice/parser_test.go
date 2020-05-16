package parser

import (
	"testing"

	"github.com/katre/techdice/dice"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	parser := New(dice.New(1))
	require.NotNil(t, parser)
	assert.Equal(t, "2", parser.Roll("1"))

	parser = New(dice.New(1))
	assert.Equal(t, "3", parser.Roll("1 push 1"))

	parser = New(dice.New(1))
	assert.Equal(t, "3.1", parser.Roll("2 push 1"))

	parser = New(dice.New(1))
	assert.Equal(t, "5", parser.Roll("3 push 2"))

	parser = New(dice.New(1))
	assert.Equal(t, "2", parser.Roll("2 hurt 1"))

	parser = New(dice.New(1))
	assert.Equal(t, "2", parser.Roll("1 push 1 hurt 1"))
}
