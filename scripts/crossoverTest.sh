#!/bin/bash

# Ensure the script exits on errors
set -e

# Define the fixed memory length
MEMORY_LENGTH=3

# Define the range of elitism rates
START_CROSSOVER=0.1
END_CROSSOVER=0.9
STEP=0.1

# Define other parameters
OPPONENT="allThree"

# Create an output directory for the CSV files
OUTPUT_DIR="1000diversityCrossover"
mkdir -p "$OUTPUT_DIR"

# Loop through elitism rates in steps of 0.1 and run the program
for CROSSOVER_RATE in $(seq $START_CROSSOVER $STEP $END_CROSSOVER); do
    # Format the elitism rate to one decimal place
    FORMATTED_CROSSOVER=$(printf "%.1f" $CROSSOVER_RATE)

    # Define the CSV file name
    CSV_FILE="$OUTPUT_DIR/fitness_crossover_${FORMATTED_CROSSOVER}.csv"

    # Run the program
    echo "Running genetic algorithm with crossover rate $FORMATTED_CROSSOVER..."
    ./prisoners \
        -memory="$MEMORY_LENGTH" \
        -opponent="$OPPONENT" \
        -csvFile="$CSV_FILE" \
        -populationSize 1000 \
        -rounds 100 \
        -generations 50 \
        -crossoverRate="$FORMATTED_CROSSOVER"

    echo "Results saved to $CSV_FILE"
done

echo "All runs completed. Results are in the '$OUTPUT_DIR' directory."