package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Check if the file path is provided
	if len(os.Args) < 2 {
		fmt.Println("Please specify the file path.")
		return
	}

	// Retrieve the file path from command-line arguments
	filePath := os.Args[1]

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Schedule the file to be closed at the end of the function
	defer file.Close()

	// Initialize the line count
	lineCount := 0

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		fmt.Printf("The file '%s' contains %d lines.\n", filePath, lineCount)
	}
}
