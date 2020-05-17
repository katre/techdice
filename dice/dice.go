package dice

import (
	"math/rand"
)

type Result struct {
	VerbDice      []int
	PushDice      []int
	HurtDice      []int
	RemainingDice []int
	Score         string
}

func NewResult(verbDice, pushDice, hurtDice []int) Result {
	// Find the highest result.
	/*
		highest := 0
		count := 0
		for _, result := range dice {
			if result > highest {
				highest = result
				count = 1
			} else if result == highest {
				count++
			}
		}

		var score string
		if count > 1 {
			score = fmt.Sprintf("%d.1", highest)
		} else {
			score = fmt.Sprintf("%d", highest)
		}
	*/
	score := "0"

	return Result{
		VerbDice:      verbDice,
		PushDice:      pushDice,
		HurtDice:      hurtDice,
		RemainingDice: []int{},
		Score:         score,
	}
}

func (r Result) String() string {
	return r.Score
}

func (r Result) Describe() string {
	return ""
}

// The actual roller.

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

func (r *Roller) rollSeveral(n int) []int {
	results := make([]int, 0, n)
	for i := 0; i < n; i++ {
		result := r.rollOne()
		results = append(results, result)
	}
	return results
}

func (r *Roller) Roll(verb, push, hurt int) Result {
	verbDice := r.rollSeveral(verb)
	pushDice := r.rollSeveral(push)
	hurtDice := r.rollSeveral(hurt)

	return NewResult(verbDice, pushDice, hurtDice)
}
