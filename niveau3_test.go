package main

import "testing"

func TestProcessLevel3(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "virgule",
			input:    "Hello , world",
			expected: "Hello, world",
		},
		{
			name:     "point",
			input:    "Hello . World",
			expected: "Hello. World",
		},
		{
			name:     "exclamation",
			input:    "Hello ! World",
			expected: "Hello! World",
		},
		{
			name:     "interrogation",
			input:    "Really ! ? No way",
			expected: "Really!? No way",
		},
		{
			name:     "points_de_suspension",
			input:    "I was thinking ... You were right",
			expected: "I was thinking... You were right",
		},
		{
			name:     "double_exclamation",
			input:    "BAMM !!",
			expected: "BAMM!!",
		},
		{
			name:     "ponctuation_multiple",
			input:    "Hello ! ! ? World",
			expected: "Hello!!? World",
		},
		{
			name:     "ingredients",
			input:    "Ingredients : flour , eggs ; milk",
			expected: "Ingredients: flour, eggs; milk",
		},
		{
			name:     "nombre_decimal",
			input:    "Pi is 3.14",
			expected: "Pi is 3.14",
		},
		{
			name:     "nombre_avec_virgule",
			input:    "1,000 is big",
			expected: "1,000 is big",
		},
		{
			name:     "ponctuation_dans_mot",
			input:    "e.g. this,that",
			expected: "e.g. this,that",
		},
		{
			name:     "espaces_multiples",
			input:    "too    many   spaces   ",
			expected: "too many spaces",
		},
		{
			name:     "ponctuation_debut_ligne",
			input:    "! Hello",
			expected: "! Hello",
		},
		{
			name:     "ponctuation_fin_ligne",
			input:    "Hello !   ",
			expected: "Hello!",
		},
		{
			name:     "phrase_complete",
			input:    "I was sitting over there ,and then BAMM !!",
			expected: "I was sitting over there, and then BAMM!!",
		},
		{
			name:     "plusieurs_lignes",
			input:    "Hello , world !\nReally ! ? No way",
			expected: "Hello, world!\nReally!? No way",
		},
		{
			name:     "ligne_vide",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProcessLevel3(tt.input)

			if result != tt.expected {
				t.Errorf(
					"\nEntrée   : %q\nAttendu  : %q\nObtenu   : %q",
					tt.input,
					tt.expected,
					result,
				)
			}
		})
	}
}