package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: Missing filename argument.")
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <filename>")
		os.Exit(1)
	}

	input := os.Args[1]

	content, err := os.ReadFile(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file '%s': %v\n", input, err)
		os.Exit(1)
	}

	fmt.Print(string(content))

	output := "output.txt"
	err = os.WriteFile(output, content, 0644)
}
