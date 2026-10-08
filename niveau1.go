package main

import (
	"strconv"
	"strings"
)

func processLevel1(line string) string {
	words := strings.Fields(line)
	result := []string{}

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(hex)" || word == "(bin)" {
			if len(result) == 0 {
				continue
			}

			previous := result[len(result)-1]

			base := 16
			maxLength := 15

			if word == "(bin)" {
				base = 2
				maxLength = 63
			}

			if len(previous) <= maxLength {
				value, err := strconv.ParseUint(previous, base, 64)

				if err == nil {
					result[len(result)-1] = strconv.FormatUint(value, 10)
				}
			}

			continue
		}

		result = append(result, word)
	}

	return strings.Join(result, " ")
}
