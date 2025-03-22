package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags for the parameters
	memoryLength := flag.Int("memory", 2, "Memory length for strategies")
	populationSize := flag.Int("populationSize", 100, "Size of the population")
	generations := flag.Int("generations", 50, "Number of generations")
	rounds := flag.Int("rounds", 10, "Number of rounds per game")
	crossoverRate := flag.Float64("crossoverRate", 0.8, "Crossover rate (0.0 to 1.0)")
	mutationRate := flag.Float64("mutationRate", 0.01, "Mutation rate (0.0 to 1.0)")
	opponentFlag := flag.String("opponent", "alwaysDefect", "Fixed opponent strategy: alwaysCooperate, alwaysDefect, titForTat, majorityRule, or allFour")

	// Parse the flags
	flag.Parse()

	// Create strategies based on the memory length
	alwaysCooperate := generateStrategy(*memoryLength, "C")
	alwaysDefect := generateStrategy(*memoryLength, "D")
	titForTat := generateTitForTat(*memoryLength)
	majorityRule := generateMajorityRule(*memoryLength)

	// Define the fixed strategies
	strategies := map[string]interface{}{
		"alwaysCooperate": []map[string]string{alwaysCooperate},
		"alwaysDefect":    []map[string]string{alwaysDefect},
		"titForTat":       []map[string]string{titForTat},
		"majorityRule":    []map[string]string{majorityRule},
		"allFour":         []map[string]string{alwaysCooperate, alwaysDefect, titForTat, majorityRule},
	}

	// Get the selected strategy or list of strategies
	fixedStrategies, exists := strategies[*opponentFlag]
	if !exists {
		fmt.Printf("Invalid opponent strategy: %s\n", *opponentFlag)
		fmt.Println("Valid options are: alwaysCooperate, alwaysDefect, titForTat, majorityRule, allFour")
		os.Exit(1)
	}

	// Run the genetic algorithm
	geneticAlgorithm(*populationSize, *generations, *rounds, *crossoverRate, *mutationRate, fixedStrategies.([]map[string]string), *memoryLength)
}
