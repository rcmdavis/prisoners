package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
)

func geneticAlgorithm(populationSize, generations, rounds int, crossoverRate, mutationRate float64, fixedStrategies []map[string]string, memoryLength int, csvFileName string) {
	// Initialize population
	population := make([]Agent, populationSize)
	for i := 0; i < populationSize; i++ {
		population[i] = initializeAgent(memoryLength)
	}

	// Create a CSV file to store total fitness per generation
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
	writer.Write([]string{fmt.Sprintf("Total Fitness (Memory Length: %d)", memoryLength)})

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

		// Write the total fitness for this generation to the CSV file
		writer.Write([]string{fmt.Sprintf("%d", totalFitness)})

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

		// Log the best fitness and strategy
		logger.Debug("Generation completed", "generation", gen+1, "bestFitness", population[0].Fitness, "bestStrategy", population[0].Strategy, "memoryLength", population[0].MemoryLength)

		// Reset fitness scores
		for i := range population {
			population[i].Fitness = 0
		}
	}
}

// Perform crossover between two agents
func crossover(parent1, parent2 Agent) Agent {
	childStrategy := map[string]string{}
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
			agent.Strategy[key] = randomMove()
		}
	}
	return agent
}
