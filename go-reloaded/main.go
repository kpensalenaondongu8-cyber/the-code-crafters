package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Correct")
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	text := string(data)
	words := Tokenize(text)

	words = ProcessHexBin(words)
	words = ProcessCase(words)
	words = ProcessPunctuation(words)
	words = ProcessArticles(words)

	result := strings.Join(words, " ")

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Correct Output saved to", outputFile)
}
