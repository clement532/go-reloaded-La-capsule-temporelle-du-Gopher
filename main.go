package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// isPunctChar dit si r est l'un des six signes de ponctuation : . , ! ? : ;
func isPunctChar(r rune) bool {
	return strings.ContainsRune(".,!?:;", r)
}

// isWord applique la définition 5 du sujet : un mot n'est pas composé
// uniquement de signes de ponctuation et/ou d'apostrophes.
func isWord(tok string) bool {
	for _, r := range tok {
		if !isPunctChar(r) && r != '\'' {
			return true
		}
	}
	return false
}

// splitTokens découpe une ligne en éléments (définitions 2 et 3).
func splitTokens(line string) []string {
	return strings.FieldsFunc(line, func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\r'
	})
}

// allDigits dit si s est non vide et composé uniquement de chiffres 0-9.
func allDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// modifier décrit un modificateur reconnu dans la ligne.
type modifier struct {
	kind string // hex, bin, up, low, cap ou rev
	n    int    // nombre de mots concernés
	size int    // nombre d'éléments occupés dans la ligne (1 ou 2)
}

// parseModifier regarde si tokens[i] commence un modificateur valide.
// "(up, 2)" occupe deux éléments : "(up," puis "2)".
func parseModifier(tokens []string, i int, smart bool) (modifier, bool) {
	tok := tokens[i]
	switch tok {
	case "(hex)", "(bin)", "(up)", "(low)", "(cap)":
		return modifier{strings.Trim(tok, "()"), 1, 1}, true
	case "(rev)":
		if smart {
			return modifier{"rev", 1, 1}, true
		}
	case "(up,", "(low,", "(cap,", "(rev,":
		kind := strings.Trim(tok, "(,")
		if kind == "rev" && !smart {
			return modifier{}, false
		}
		if i+1 >= len(tokens) || !strings.HasSuffix(tokens[i+1], ")") {
			return modifier{}, false
		}
		digits := strings.TrimSuffix(tokens[i+1], ")")
		if !allDigits(digits) {
			return modifier{}, false
		}
		n, err := strconv.Atoi(digits)
		if err != nil {
			n = math.MaxInt // nombre énorme : on s'applique à tous les mots
		}
		return modifier{kind, n, 2}, true
	}
	return modifier{}, false
}

// convert lit tok dans la base donnée (16 ou 2) et renvoie sa valeur
// décimale. Le booléen vaut false si tok n'est pas valide dans cette base.
func convert(tok string, base int) (string, bool) {
	maxLen := 15
	if base == 2 {
		maxLen = 63
	}
	if len(tok) > maxLen {
		return tok, false
	}
	n, err := strconv.ParseUint(tok, base, 64)
	if err != nil {
		return tok, false
	}
	return strconv.FormatUint(n, 10), true
}

// capitalize met la première lettre en majuscule et les autres en minuscules.
func capitalize(s string) string {
	runes := []rune(s)
	first := true
	for i, r := range runes {
		if !unicode.IsLetter(r) {
			continue
		}
		if first {
			runes[i] = unicode.ToUpper(r)
			first = false
		} else {
			runes[i] = unicode.ToLower(r)
		}
	}
	return string(runes)
}

// reverse écrit les caractères d'un mot à l'envers (bonus).
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// transform applique un modificateur de casse (ou rev) à un mot.
func transform(kind, s string) string {
	switch kind {
	case "up":
		return strings.ToUpper(s)
	case "low":
		return strings.ToLower(s)
	case "cap":
		return capitalize(s)
	case "rev":
		return reverse(s)
	}
	return s
}

// applyModifier agit sur les éléments déjà construits (out) : il remonte
// depuis la fin en ignorant tout ce qui n'est pas un mot.
func applyModifier(out []string, m modifier) {
	if m.kind == "hex" || m.kind == "bin" {
		base := 16
		if m.kind == "bin" {
			base = 2
		}
		for i := len(out) - 1; i >= 0; i-- {
			if isWord(out[i]) {
				if v, ok := convert(out[i], base); ok {
					out[i] = v
				}
				return
			}
		}
		return
	}
	count := 0
	for i := len(out) - 1; i >= 0 && count < m.n; i-- {
		if !isWord(out[i]) {
			continue
		}
		out[i] = transform(m.kind, out[i])
		count++
	}
}

// applyModifiers traite toute la ligne de gauche à droite.
func applyModifiers(tokens []string, smart bool) []string {
	out := make([]string, 0, len(tokens))
	for i := 0; i < len(tokens); {
		m, ok := parseModifier(tokens, i, smart)
		if !ok {
			out = append(out, tokens[i])
			i++
			continue
		}
		applyModifier(out, m)
		i += m.size
	}
	return out
}

// fixPunctuation colle les signes à la fin du mot précédent.
func fixPunctuation(tokens []string) []string {
	out := make([]string, 0, len(tokens))
	for _, tok := range tokens {
		runes := []rune(tok)
		k := 0
		for k < len(runes) && isPunctChar(runes[k]) {
			k++
		}
		lead, rest := string(runes[:k]), string(runes[k:])
		if lead != "" {
			if len(out) == 0 {
				out = append(out, lead) // signe en début de ligne
			} else {
				out[len(out)-1] += lead
			}
		}
		if rest != "" {
			out = append(out, rest)
		}
	}
	return out
}

// fixQuotes associe les guillemets deux par deux et supprime les espaces
// juste après l'ouvrant et juste avant le fermant. Un ' entre deux lettres
// est une apostrophe et n'est jamais touché.
func fixQuotes(line string) string {
	runes := []rune(line)
	var quotes []int
	for i, r := range runes {
		if r != '\'' {
			continue
		}
		isApostrophe := i > 0 && i < len(runes)-1 &&
			unicode.IsLetter(runes[i-1]) && unicode.IsLetter(runes[i+1])
		if !isApostrophe {
			quotes = append(quotes, i)
		}
	}
	remove := make([]bool, len(runes))
	for p := 0; p+1 < len(quotes); p += 2 {
		open, closing := quotes[p], quotes[p+1]
		for j := open + 1; j < closing && runes[j] == ' '; j++ {
			remove[j] = true
		}
		for j := closing - 1; j > open && runes[j] == ' '; j-- {
			remove[j] = true
		}
	}
	var sb strings.Builder
	for i, r := range runes {
		if !remove[i] {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func isArticle(tok string) bool {
	return tok == "a" || tok == "A" || tok == "an" || tok == "An"
}

// chooseArticle renvoie "a" ou "an" d'après l'élément suivant. Le booléen
// vaut false quand l'article ne doit pas changer.
func chooseArticle(next string, smart bool) (string, bool) {
	if smart {
		word := strings.ToLower(strings.Trim(next, ".,!?:;'"))
		switch word {
		case "hour", "honest", "honor", "heir":
			return "an", true
		case "university", "unicorn", "unique", "user", "useful", "one":
			return "a", true
		}
	}
	runes := []rune(strings.TrimLeft(next, "'"))
	if len(runes) == 0 || !unicode.IsLetter(runes[0]) {
		return "", false
	}
	if strings.ContainsRune("aeiouAEIOU", runes[0]) {
		return "an", true
	}
	return "a", true
}

func fixArticles(tokens []string, smart bool) []string {
	for i := 0; i+1 < len(tokens); i++ {
		if !isArticle(tokens[i]) {
			continue
		}
		article, ok := chooseArticle(tokens[i+1], smart)
		if !ok {
			continue
		}
		if tokens[i][0] == 'A' {
			article = "A" + article[1:]
		}
		tokens[i] = article
	}
	return tokens
}

// processLine applique les niveaux dans l'ordre : 1-2, 3, 4, 5.
func processLine(line string, smart bool) string {
	tokens := splitTokens(line)
	tokens = applyModifiers(tokens, smart)
	tokens = fixPunctuation(tokens)
	line = fixQuotes(strings.Join(tokens, " "))
	tokens = fixArticles(splitTokens(line), smart)
	return strings.Join(tokens, " ")
}

// run fait tout le travail et renvoie le code de sortie du programme.
func run(args []string) int {
	smart := false
	if len(args) > 0 && args[0] == "--smart" {
		smart = true
		args = args[1:]
	}
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go run . [--smart] <fichier d'entrée> <fichier de sortie>")
		return 1
	}
	data, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "erreur de lecture :", err)
		return 1
	}
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		lines[i] = processLine(line, smart)
	}
	err = os.WriteFile(args[1], []byte(strings.Join(lines, "\n")), 0o644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "erreur d'écriture :", err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
