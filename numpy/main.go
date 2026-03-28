package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <inputfile> <outputfile>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read input
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Process text
	result := ProcessText(string(data))

	// Write output
	err = os.WriteFile(outputFile, []byte(result+"\n"), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}
}
