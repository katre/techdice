package dice

import (
	"bytes"
	"fmt"
	"math/rand"
)

type Result struct {
	VerbDice      []int
	PushDice      []int
	ManaDice      []int
	HurtDice      []int
	RemainingDice []int
	Score         string
}

func addDice(dice []int, invalid map[int]bool, added []int) []int {
	for _, val := range added {
		if !invalid[val] {
			dice = append(dice, val)
		}
	}
	return dice
}

func NewResult(verbDice, pushDice, manaDice, hurtDice []int) Result {
	// Calculate the remaining live dice.
	invalid := make(map[int]bool)
	for _, val := range hurtDice {
		invalid[val] = true
	}

	remaining := make([]int, 0, len(verbDice)+len(pushDice)+len(manaDice))
	remaining = addDice(remaining, invalid, verbDice)
	remaining = addDice(remaining, invalid, pushDice)
	remaining = addDice(remaining, invalid, manaDice)

	// Find the highest result.
	highest := 0
	count := 0
	for _, result := range remaining {
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

	return Result{
		VerbDice:      verbDice,
		PushDice:      pushDice,
		ManaDice:      manaDice,
		HurtDice:      hurtDice,
		RemainingDice: remaining,
		Score:         score,
	}
}

func (r Result) String() string {
	return r.Score
}

func (r Result) Describe() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s %v", r.Score, r.VerbDice)
	if len(r.PushDice) > 0 {
		fmt.Fprintf(&buf, " push: %v", r.PushDice)
	}
	if len(r.ManaDice) > 0 {
		fmt.Fprintf(&buf, " mana: %v", r.ManaDice)
	}
	if len(r.HurtDice) > 0 {
		fmt.Fprintf(&buf, " hurt: %v", r.HurtDice)
	}
	return buf.String()
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

func (r *Roller) Roll(verb, push, mana, hurt int) Result {
	verbDice := r.rollSeveral(verb)
	pushDice := r.rollSeveral(push)
	manaDice := r.rollSeveral(mana)
	hurtDice := r.rollSeveral(hurt)

	return NewResult(verbDice, pushDice, manaDice, hurtDice)
}
