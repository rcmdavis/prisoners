#!/bin/bash

# Ensure the script exits on errors
set -e

# Define the directory containing the CSV files
INPUT_DIR="results"
OUTPUT_FILE="combined_fitness.csv"

# Find all CSV files and sort them by memory length
CSV_FILES=$(ls "$INPUT_DIR"/fitness_memory_*.csv | sort -V)

# Combine the files line by line, separated by commas
paste -d ',' $CSV_FILES > "$OUTPUT_FILE"

echo "Combined fitness data saved to $OUTPUT_FILE"