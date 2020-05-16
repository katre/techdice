package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRollOnce(t *testing.T) {
	for i := 0; i < 100; i++ {
		value := rollOne()
		assert.Greater(t, value, 0)
		assert.Less(t, value, 7)
	}
}
