package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <fichier entrée> <fichier sortie>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")

	var result []string

	for _, line := range lines {
		result = append(result, processLevel1(line))
	}

	err = os.WriteFile(outputFile, []byte(strings.Join(result, "\n")), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}
}
