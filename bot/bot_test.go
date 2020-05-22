package bot

import (
	"testing"

	"github.com/katre/techdice/dice"

	"github.com/stretchr/testify/assert"
)

func TestDescribe(t *testing.T) {
	check := func(expected string, verb, push, mana, hurt []int) {
		result := dice.NewResult(verb, push, mana, hurt)
		s := describe(result)
		assert.Equal(t, expected, s)
	}
	check("Score: 3 [3]", []int{3}, []int{}, []int{}, []int{})
	check("Score: 5 [3 5, push: 2]", []int{3, 5}, []int{2}, []int{}, []int{})
	check("Score: 5 [3 5, mana: 1]", []int{3, 5}, []int{}, []int{1}, []int{})
	check("Score: 5 [3 5, push: 4 2, mana: 1]", []int{3, 5}, []int{4, 2}, []int{1}, []int{})
	check("Score: 5 [~~3~~ 5, hurt: 3]", []int{3, 5}, []int{}, []int{}, []int{3})
	check("Score: 3 [3 ~~5~~, hurt: 5]", []int{3, 5}, []int{}, []int{}, []int{5})
}
