

package main

import (
	"strings"
	"unicode"
)

// Vérifie si un caractère est un signe de ponctuation.
func isPunctuation(r rune) bool {
	return strings.ContainsRune(".,!?:;", r)
}

// Corrige la ponctuation et les espaces d'une ligne.
func normalizeLine(line string) string {
	runes := []rune(line)

	var result strings.Builder
	var word strings.Builder

	// Ajoute le mot en cours au résultat.
	flushWord := func() {
		if word.Len() == 0 {
			return
		}

		if result.Len() > 0 {
			result.WriteByte(' ')
		}

		result.WriteString(word.String())
		word.Reset()
	}

	for i := 0; i < len(runes); {
		r := runes[i]

		// Ignore les espaces multiples.
		if unicode.IsSpace(r) {
			flushWord()
			i++
			continue
		}

		// Gestion des signes de ponctuation.
		if isPunctuation(r) {
			start := i

			// Regroupe les signes consécutifs.
			for i < len(runes) && isPunctuation(runes[i]) {
				i++
			}

			group := string(runes[start:i])

			// Ponctuation à l'intérieur d'un mot.
			// Exemples : 3.14, 1,000, this,that
			insideWord := start > 0 &&
				i < len(runes) &&
				!unicode.IsSpace(runes[start-1]) &&
				!unicode.IsSpace(runes[i])

			if insideWord {
				word.WriteString(group)
			} else {
				flushWord()
				result.WriteString(group)
			}

			continue
		}

		word.WriteRune(r)
		i++
	}

	flushWord()

	return result.String()
}

// Traite toutes les lignes du texte.
func ProcessLevel3(text string) string {
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		lines[i] = normalizeLine(line)
	}

	return strings.Join(lines, "\n")
}

