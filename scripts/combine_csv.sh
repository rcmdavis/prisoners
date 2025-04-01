#!/bin/bash

# Ensure the script exits on errors
set -e

# Check if the correct number of arguments is provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <input_directory> <output_file>"
    exit 1
fi

# Get the input directory and output file from arguments
INPUT_DIR="$1"
OUTPUT_FILE="$2"

# Find all CSV files in the input directory and sort them by memory length
CSV_FILES=$(ls "$INPUT_DIR"/*.csv | sort -V)

# Combine the files line by line, separated by commas
paste -d ',' $CSV_FILES > "$OUTPUT_FILE"

echo "Combined fitness data saved to $OUTPUT_FILE"