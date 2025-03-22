package main

import "math/rand"

// Agent represents an individual in the population
type Agent struct {
	Strategy     map[string]string // Strategy as a map
	Fitness      int               // Fitness score
	MemoryLength int               // Memory length of the agent
}

// Initialize a random strategy for an agent
func initializeAgent(memoryLength int) Agent {
	strategy := generateRandomStrategy(memoryLength)
	return Agent{Strategy: strategy, Fitness: 0, MemoryLength: memoryLength}
}

// Generate a random strategy for a given memory length
func generateRandomStrategy(memoryLength int) map[string]string {
	combinations := generateCombinations(memoryLength)
	strategy := make(map[string]string)

	for _, combo := range combinations {
		strategy[combo] = randomMove()
	}

	return strategy
}

// Generate a random move ("C" or "D")
func randomMove() string {
	if rand.Intn(2) == 0 {
		return "C"
	}
	return "D"
}
