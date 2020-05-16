package dice

import (
	"fmt"
	"math/rand"
)

type Roller struct {
	rand *rand.Rand
}

func New(seed int64) *Roller {
	s := rand.NewSource(seed)
	return &Roller{
		rand: rand.New(s),
	}
}

func (r *Roller) rollOne() int {
	return r.rand.Intn(5) + 1
}

func (r *Roller) Roll(verb int, push int, hurt int) string {
	dice := verb + push
	// Roll the positive dice.
	results := make([]int, 0, dice)
	for i := 0; i < dice; i++ {
		result := r.rollOne()
		results = append(results, result)
	}

	// Roll the hurt dice and remove any matches.
	for i := 0; i < hurt; i++ {
		remove := r.rollOne()
		for i, result := range results {
			if result == remove {
				results[i] = 0
			}
		}
	}

	// Find the highest result.
	highest := 0
	count := 0
	for _, result := range results {
		if result > highest {
			highest = result
			count = 1
		} else if result == highest {
			count++
		}
	}

	if count > 1 {
		return fmt.Sprintf("%d.1", highest)
	} else {
		return fmt.Sprintf("%d", highest)
	}
}
