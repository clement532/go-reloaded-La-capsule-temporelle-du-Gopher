package main

import "testing"

// check vérifie une liste de paires {entrée, sortie attendue}.
func check(t *testing.T, smart bool, cases [][2]string) {
	t.Helper()
	for _, c := range cases {
		if got := processLine(c[0], smart); got != c[1] {
			t.Errorf("entrée %q : obtenu %q, attendu %q", c[0], got, c[1])
		}
	}
}

func TestLevel1(t *testing.T) {
	check(t, false, [][2]string{
		{"1E (hex) files were added", "30 files were added"},
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"ff (hex) red", "255 red"},
		{"We need ZZ (hex) more", "We need ZZ more"},
		{"Call 102 (bin) now", "Call 102 now"},
		{"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			"Simply add 66 and 2 and you will see the result is 68."},
		{"(hex) nothing before", "nothing before"},
	})
}

func TestLevel2(t *testing.T) {
	check(t, false, [][2]string{
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"Welcome to the brooklyn bridge (cap)", "Welcome to the brooklyn Bridge"},
		{"hELLO (cap) there", "Hello there"},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"hello , world (up, 2)", "HELLO, WORLD"},
		{"only two (cap, 5)", "Only Two"},
		{"(up) nothing before", "nothing before"},
		{"1010 (bin) (cap)", "10"},
		{"loud (low) (up)", "LOUD"},
		{"keep calm (up, 0)", "keep calm"},
		{"go (up)! and (up,2) stay", "go (up)! and (up,2) stay"},
		{"10 (hex, 2) stays", "10 (hex, 2) stays"},
	})
}
