#!/bin/bash

# Output CSV file
output_csv="time_results.csv"

# Write the header row to the CSV file
echo "Memory,TimeTaken(s)" > "$output_csv"

# Loop through memory values from 1 to 10
for memory in {1..10}; do
    echo "Running with memory=$memory..."
    
    # Record the start time with nanosecond precision
    start_time=$(date +%s.%N)
    
    # Run the program
    ./prisoners -memory="$memory" -csvFile="results_memory${memory}.csv"
    
    # Record the end time with nanosecond precision
    end_time=$(date +%s.%N)
    
    # Calculate the elapsed time with greater precision
    elapsed_time=$(echo "$end_time - $start_time" | bc)
    
    # Write the memory and time taken to the CSV file
    echo "$memory,$elapsed_time" >> "$output_csv"
    
    echo "Time taken for memory=$memory: ${elapsed_time}s"
done

echo "All runs completed. Time results saved in $output_csv."