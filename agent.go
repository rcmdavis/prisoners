package main

import "math/rand"

// Agent represents an individual in the population
type Agent struct {
	Strategy     map[string]float64 // Strategy as a map of probabilities
	Fitness      int                // Fitness score
	MemoryLength int                // Memory length of the agent
}

// Initialize a random strategy for an agent
func initializeAgent(memoryLength int) Agent {
	strategy := generateRandomStrategy(memoryLength)
	return Agent{Strategy: strategy, Fitness: 0, MemoryLength: memoryLength}
}

// Generate a random strategy for a given memory length
func generateRandomStrategy(memoryLength int) map[string]float64 {
	combinations := generateCombinations(memoryLength)
	strategy := make(map[string]float64)

	for _, combo := range combinations {
		strategy[combo] = rand.Float64() // Random probability between 0.0 and 1.0
	}

	return strategy
}
