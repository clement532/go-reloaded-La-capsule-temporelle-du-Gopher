
package main

import "strings"

func isPunctuation(c byte) bool {
    return strings.ContainsRune(".,!?:;", rune(c))
}

func fixPunctuation(text string) string {
    words := strings.Fields(text)
    result := ""

    for _, word := range words {
        onlyPunctuation := true

        for i := 0; i < len(word); i++ {
            if !isPunctuation(word[i]) {
                onlyPunctuation = false
                break
            }
        }

        if onlyPunctuation {
            result += word
        } else {
            if result != "" {
                if !isPunctuation(result[len(result)-1]) {
                    result += " "
                } else {
                    result += " "
                }
            }
            result += word
        }
    }

    return result
}
