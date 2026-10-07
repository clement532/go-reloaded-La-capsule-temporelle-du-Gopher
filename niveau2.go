package main

import (
	"strconv"
	"strings"
)

func processLevel2(line string) string {
	words := strings.Fields(line)
	result := []string{}

	for i := 0; i < len(words); i++ {
		word := words[i]

		// (up)
		if word == "(up)" {
			applyModifier(result, 1, "up")
			continue
		}

		// (low)
		if word == "(low)" {
			applyModifier(result, 1, "low")
			continue
		}

		// (cap)
		if word == "(cap)" {
			applyModifier(result, 1, "cap")
			continue
		}

		// (up, N)
		if word == "(up," && i+1 < len(words) {
			n := getNumber(words[i+1])

			applyModifier(result, n, "up")
			i++
			continue
		}

		// (low, N)
		if word == "(low," && i+1 < len(words) {
			n := getNumber(words[i+1])

			applyModifier(result, n, "low")
			i++
			continue
		}

		// (cap, N)
		if word == "(cap," && i+1 < len(words) {
			n := getNumber(words[i+1])

			applyModifier(result, n, "cap")
			i++
			continue
		}

		result = append(result, word)
	}

	return strings.Join(result, " ")
}

func getNumber(word string) int {
	word = strings.TrimSuffix(word, ")")
	n, err := strconv.Atoi(word)

	if err != nil {
		return 0
	}

	return n
}

func applyModifier(words []string, n int, modifier string) {
	if n <= 0 {
		return
	}

	start := len(words) - n

	if start < 0 {
		start = 0
	}

	for i := start; i < len(words); i++ {
		switch modifier {
		case "up":
			words[i] = strings.ToUpper(words[i])

		case "low":
			words[i] = strings.ToLower(words[i])

		case "cap":
			words[i] = capitalize(words[i])
		}
	}
}

func capitalize(word string) string {
	if word == "" {
		return word
	}

	runes := []rune(strings.ToLower(word))
	runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]

	return string(runes)
}
