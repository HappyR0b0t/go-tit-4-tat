package game

import "math/rand"

// Move interface
type Strategy interface {
	Move(turnNumber int) int
}

// Random strategy
type RandomStrategy struct{}

func (s RandomStrategy) Move(turnNumber int) int {
	return rand.Intn(2)
}

// Always return 1
type AlwaysOneStrategy struct{}

func (s AlwaysOneStrategy) Move(turnNumber int) int {
	return 1
}

// Always return 0
type AlwaysZeroStrategy struct{}

func (s AlwaysZeroStrategy) Move(turnNumber int) int {
	return 0
}

// Return 1 on even turns, return 0 on uneven turns
type OnesAndZeroesStrategy struct{}

func (s OnesAndZeroesStrategy) Move(turnNumber int) int {
	if turnNumber == 0 {
		return 1
	} else if turnNumber%2 == 0 {
		return 1
	} else {
		return 0
	}
}

// Copy opponents moves
type CopycatStrategy struct{}

func (s CopycatStrategy) Move(turnNumber int, turns [][2]int) int {
	if turnNumber == 0 {
		return 1
	} else {
		return 0
	}
}

// Generic strategy
type GenericStrategy struct {
	turns []int
}

func (s GenericStrategy) Move(turnNumber int) int {
	return s.turns[turnNumber]
}

func BuildCombinations(result *[][]int, currentCombination []int, elements []int, targetLength int) {
	if len(currentCombination) == targetLength {
		combo := make([]int, targetLength)
		copy(combo, currentCombination)
		*result = append(*result, combo)
		return
	}

	for _, el := range elements {
		newCombination := append(currentCombination, el)
		BuildCombinations(result, newCombination, elements, targetLength)
	}
}

func GenerateGenericStrategies(s GenericStrategy, turns []int) GenericStrategy {
	s.turns = turns
	return s
}
