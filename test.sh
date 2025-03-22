#!/bin/bash

# Ensure the script exits on errors
set -e

# Define the range of memory lengths
START_MEMORY=1
END_MEMORY=30

# Define other parameters
OPPONENT="allFour"

# Create an output directory for the CSV files
OUTPUT_DIR="results"
mkdir -p "$OUTPUT_DIR"

# Loop through memory lengths and run the program
for MEMORY_LENGTH in $(seq $START_MEMORY $END_MEMORY); do
    # Define the CSV file name
    CSV_FILE="$OUTPUT_DIR/fitness_memory_${MEMORY_LENGTH}.csv"

    # Run the program
    echo "Running genetic algorithm with memory length $MEMORY_LENGTH..."
    ./prisoners \
        -memory="$MEMORY_LENGTH" \
        -opponent="$OPPONENT" \
        -csvFile="$CSV_FILE" 

    echo "Results saved to $CSV_FILE"
done

echo "All runs completed. Results are in the '$OUTPUT_DIR' directory."