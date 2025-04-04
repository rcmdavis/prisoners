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

// Generate a strategy map for a given memory length and default move
func generateStrategy(memoryLength int, defaultMove string) map[string]string {
	combinations := generateCombinations(memoryLength)
	strategy := make(map[string]string)

	for _, combo := range combinations {
		strategy[combo] = defaultMove
	}

	return strategy
}

// Generate the Tit for Tat strategy
func generateTitForTat(memoryLength int) map[string]string {
	combinations := generateCombinations(memoryLength)
	strategy := make(map[string]string)

	for _, combo := range combinations {
		// Tit for Tat mimics the opponent's last move
		if combo == "" {
			strategy[combo] = "C" // Default to cooperate for the first move
		} else {
			strategy[combo] = string(combo[len(combo)-1]) // Use the last move in the memory
		}
	}

	return strategy
}
