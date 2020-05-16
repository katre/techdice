package dice

import (
	"math/rand"
)

func rollOne() int {
	return rand.Intn(5) + 1
}
