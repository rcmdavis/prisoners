package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"
)

var logger *slog.Logger

func main() {
	// Define flags for the parameters
	memoryLength := flag.Int("memory", 2, "Memory length for strategies")
	populationSize := flag.Int("populationSize", 100, "Size of the population")
	generations := flag.Int("generations", 50, "Number of generations")
	rounds := flag.Int("rounds", 100, "Number of rounds per game")
	crossoverRate := flag.Float64("crossoverRate", 0.8, "Crossover rate (0.0 to 1.0)")
	mutationRate := flag.Float64("mutationRate", 0.01, "Mutation rate (0.0 to 1.0)")
	opponentFlag := flag.String("opponent", "alwaysDefect", "Fixed opponent strategy: alwaysCooperate, alwaysDefect, titForTat, majorityRule, or allFour")
	csvFile := flag.String("csvFile", "fitness_per_generation.csv", "Name of the CSV file to store fitness data")
	logLevel := flag.String("logLevel", "WARN", "Log level: DEBUG, INFO, WARN, or ERROR")

	// Parse the flags
	flag.Parse()

	// Configure the logger based on the log level
	var level slog.Level
	switch *logLevel {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		fmt.Printf("Invalid log level: %s. Defaulting to INFO.\n", *logLevel)
		level = slog.LevelInfo
	}

	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))

	// Log the start of the program
	logger.Info("Starting the genetic algorithm", "memoryLength", *memoryLength, "populationSize", *populationSize, "generations", *generations)

	// Create strategies based on the memory length
	alwaysCooperate := generateStrategy(*memoryLength, "C")
	alwaysDefect := generateStrategy(*memoryLength, "D")
	titForTat := generateTitForTat(*memoryLength)

	// Define the fixed strategies
	strategies := map[string]interface{}{
		"alwaysCooperate": []map[string]string{alwaysCooperate},
		"alwaysDefect":    []map[string]string{alwaysDefect},
		"titForTat":       []map[string]string{titForTat},
		"allThree":        []map[string]string{alwaysCooperate, alwaysDefect, titForTat},
	}

	// Get the selected strategy or list of strategies
	fixedStrategies, exists := strategies[*opponentFlag]
	if !exists {
		logger.Error("Invalid opponent strategy", "opponent", *opponentFlag)
		fmt.Println("Valid options are: alwaysCooperate, alwaysDefect, titForTat, majorityRule, allFour")
		os.Exit(1)
	}

	startTime := time.Now()
	// Run the genetic algorithm
	geneticAlgorithm(*populationSize, *generations, *rounds, *crossoverRate, *mutationRate, fixedStrategies.([]map[string]string), *memoryLength, *csvFile)
	// Calculate and log the total time taken
	elapsedTime := time.Since(startTime)
	logger.Info("Genetic algorithm completed", "timeTaken", elapsedTime)
	fmt.Printf("Time taken for this run: %s\n", elapsedTime)
}
