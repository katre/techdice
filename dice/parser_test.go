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

	assert.Equal(t, "2", parser.Roll("tech 1"))
}
