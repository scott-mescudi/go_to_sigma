package main

import (
	"fmt"
	"os"
	"amgis/lexer"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: gosigma build <input.go> <output.sigma>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	test := string(file)
	err = os.WriteFile(outputFile, []byte(lexer.Words(test)), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("Successfully processed %s and wrote to %s\n", inputFile, outputFile)
}
