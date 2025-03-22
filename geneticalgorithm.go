package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
)

func geneticAlgorithm(populationSize, generations, rounds int, crossoverRate, mutationRate float64, fixedStrategies []map[string]float64, memoryLength int, csvFileName string) {
	// Initialize population
	population := make([]Agent, populationSize)
	for i := 0; i < populationSize; i++ {
		population[i] = initializeAgent(memoryLength)
	}

	// Create a CSV file to store total fitness and diversity per generation
	csvFile, err := os.Create(csvFileName)
	if err != nil {
		logger.Error("Error creating CSV file", "error", err)
		return
	}
	defer csvFile.Close()

	// Create a CSV writer
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write the header row
	writer.Write([]string{fmt.Sprintf("Total Fitness (Memory Length: %d)", memoryLength), "Diversity"})

	// Evolve over generations
	for gen := 0; gen < generations; gen++ {
		// Evaluate fitness
		totalFitness := 0
		for i := range population {
			population[i].Fitness = 0 // Reset fitness for this generation

			// Evaluate fitness against each fixed strategy
			for _, fixedStrategy := range fixedStrategies {
				population[i].Fitness += playIPD(population[i].Strategy, fixedStrategy, rounds)
			}

			totalFitness += population[i].Fitness
		}

		// Calculate diversity
		diversity := calculateDiversity(population)

		// Write the total fitness and diversity for this generation to the CSV file
		writer.Write([]string{fmt.Sprintf("%d", totalFitness), fmt.Sprintf("%d", diversity)})

		// Sort population by fitness (descending)
		sort.Slice(population, func(i, j int) bool {
			return population[i].Fitness > population[j].Fitness
		})

		// Create new population
		newPopulation := []Agent{}

		// Elitism: Retain top Y% of agents
		elitismCount := int(float64(populationSize) * (1 - crossoverRate))
		newPopulation = append(newPopulation, population[:elitismCount]...)

		// Crossover: Generate new agents from top X% of agents
		for len(newPopulation) < populationSize {
			parent1 := population[rand.Intn(elitismCount)]
			parent2 := population[rand.Intn(elitismCount)]
			child := crossover(parent1, parent2)
			newPopulation = append(newPopulation, child)
		}

		// Mutate the new population
		for i := range newPopulation {
			newPopulation[i] = mutate(newPopulation[i], mutationRate)
		}

		// Replace old population with new population
		population = newPopulation

		// Log the best fitness, strategy, and diversity
		logger.Info("Generation completed", "generation", gen+1, "bestFitness", population[0].Fitness, "diversity", diversity, "memoryLength", population[0].MemoryLength,
			"bestStrategy", population[0].Strategy)

		// Reset fitness scores
		for i := range population {
			population[i].Fitness = 0
		}
	}
}

// Calculate diversity in the population
func calculateDiversity(population []Agent) int {
	uniqueStrategies := make(map[string]bool)
	for _, agent := range population {
		strategyKey := fmt.Sprintf("%v", agent.Strategy) // Serialize the strategy as a string
		uniqueStrategies[strategyKey] = true
	}
	return len(uniqueStrategies)
}

// Perform crossover between two agents
func crossover(parent1, parent2 Agent) Agent {
	childStrategy := map[string]float64{}
	for key := range parent1.Strategy {
		if rand.Float64() < 0.5 {
			childStrategy[key] = parent1.Strategy[key]
		} else {
			childStrategy[key] = parent2.Strategy[key]
		}
	}
	return Agent{Strategy: childStrategy, Fitness: 0, MemoryLength: parent1.MemoryLength}
}

// Mutate an agent's strategy
func mutate(agent Agent, mutationRate float64) Agent {
	for key := range agent.Strategy {
		if rand.Float64() < mutationRate {
			// Adjust the probability slightly, ensuring it stays within [0.0, 1.0]
			agent.Strategy[key] += (rand.Float64() - 0.5) * 0.1
			if agent.Strategy[key] < 0.0 {
				agent.Strategy[key] = 0.0
			} else if agent.Strategy[key] > 1.0 {
				agent.Strategy[key] = 1.0
			}
		}
	}
	return agent
}
