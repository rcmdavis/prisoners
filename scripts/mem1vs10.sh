#!/bin/bash

# Create directories for results
mkdir -p memory1 memory10

# Define opponents
opponents=("alwaysCooperate" "alwaysDefect" "titForTat" "allThree")

# Run for memory length 1
for opponent in "${opponents[@]}"; do
    output_file="memory1/${opponent}_results.csv"
    echo "Running with memory=1 against $opponent..."
    ./prisoners -memory=1 -opponent="$opponent" -csvFile="$output_file"
done

# Run for memory length 10
for opponent in "${opponents[@]}"; do
    output_file="memory10/${opponent}_results.csv"
    echo "Running with memory=10 against $opponent..."
    ./prisoners -memory=10 -opponent="$opponent" -csvFile="$output_file"
done

echo "All runs completed. Results saved in 'memory1' and 'memory10' directories."