package main

// Define possible moves
const (
	Cooperate = "C"
	Defect    = "D"
)

// Generate all possible combinations of moves for memory lengths from 0 to the given memory length
func generateCombinations(memoryLength int) []string {
	var allCombinations []string

	// Loop through all memory lengths from 0 to memoryLength
	for length := 0; length <= memoryLength; length++ {
		allCombinations = append(allCombinations, generateCombinationsForLength(length)...)
	}

	return allCombinations
}

// Generate all possible combinations of moves for a specific memory length
func generateCombinationsForLength(length int) []string {
	if length == 0 {
		return []string{""} // Include "" for the initial state
	}

	smallerCombinations := generateCombinationsForLength(length - 1)
	var combinations []string

	for _, combo := range smallerCombinations {
		combinations = append(combinations, combo+"C", combo+"D")
	}

	return combinations
}

// Generate the Always Cooperate strategy
func generateAlwaysCooperate(memoryLength int) map[string]float64 {
	combinations := generateCombinations(memoryLength)
	strategy := make(map[string]float64)

	for _, combo := range combinations {
		strategy[combo] = 1.0 // Always cooperate
	}

	return strategy
}

// Generate the Always Defect strategy
func generateAlwaysDefect(memoryLength int) map[string]float64 {
	combinations := generateCombinations(memoryLength)
	strategy := make(map[string]float64)

	for _, combo := range combinations {
		strategy[combo] = 0.0 // Always defect
	}

	return strategy
}

// Generate the Tit for Tat strategy
func generateTitForTat(memoryLength int) map[string]float64 {
	combinations := generateCombinations(memoryLength)
	strategy := make(map[string]float64)

	for _, combo := range combinations {
		if combo == "" {
			strategy[combo] = 1.0 // Always cooperate on the first move
		} else if combo[len(combo)-1] == 'C' {
			strategy[combo] = 0.9 // High probability to cooperate if the opponent cooperated
		} else {
			strategy[combo] = 0.1 // Low probability to cooperate if the opponent defected
		}
	}

	return strategy
}

// Generate the Majority Rule strategy
func generateMajorityRule(memoryLength int) map[string]float64 {
	combinations := generateCombinations(memoryLength)
	strategy := make(map[string]float64)

	for _, combo := range combinations {
		if combo == "" {
			strategy[combo] = 1.0 // Always cooperate on the first move
		} else {
			// Count occurrences of C and D in the opponent's history
			countC := 0
			countD := 0
			for _, move := range combo {
				if move == 'C' {
					countC++
				} else if move == 'D' {
					countD++
				}
			}

			// Decide based on the majority
			if countC > countD {
				strategy[combo] = 0.9 // High probability to cooperate if C is the majority
			} else if countD > countC {
				strategy[combo] = 0.1 // Low probability to cooperate if D is the majority
			} else {
				strategy[combo] = 0.5 // Moderate probability to cooperate in case of a tie
			}
		}
	}

	return strategy
}
