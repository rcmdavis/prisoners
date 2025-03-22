package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags for the parameters
	memoryLength := flag.Int("memoryLength", 3, "Memory length for strategies")
	populationSize := flag.Int("populationSize", 100, "Size of the population")
	generations := flag.Int("generations", 50, "Number of generations")
	rounds := flag.Int("rounds", 10, "Number of rounds per game")
	crossoverRate := flag.Float64("crossoverRate", 0.8, "Crossover rate (0.0 to 1.0)")
	mutationRate := flag.Float64("mutationRate", 0.01, "Mutation rate (0.0 to 1.0)")
	opponentFlag := flag.String("opponent", "alwaysDefect", "Fixed opponent strategy: alwaysCooperate, alwaysDefect, or titForTat")

	// Parse the flags
	flag.Parse()

	// Create strategies based on the memory length
	alwaysCooperate := generateStrategy(*memoryLength, "C")
	alwaysDefect := generateStrategy(*memoryLength, "D")
	titForTat := generateTitForTat(*memoryLength)

	// Define the fixed strategies
	strategies := map[string]map[string]string{
		"alwaysCooperate": alwaysCooperate,
		"alwaysDefect":    alwaysDefect,
		"titForTat":       titForTat,
	}

	// Get the selected strategy
	fixedStrategy, exists := strategies[*opponentFlag]
	if !exists {
		fmt.Printf("Invalid opponent strategy: %s\n", *opponentFlag)
		fmt.Println("Valid options are: alwaysCooperate, alwaysDefect, titForTat")
		os.Exit(1)
	}

	// Run the genetic algorithm
	geneticAlgorithm(*populationSize, *generations, *rounds, *crossoverRate, *mutationRate, fixedStrategy, *memoryLength)
}
