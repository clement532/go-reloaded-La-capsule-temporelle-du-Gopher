package main

import "testing"

func TestProcessLevel2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "up",
			input: "hello (up)",
			want:  "HELLO",
		},
		{
			name:  "low",
			input: "HELLO (low)",
			want:  "hello",
		},
		{
			name:  "cap",
			input: "hello (cap)",
			want:  "Hello",
		},
		{
			name:  "multiple modifiers",
			input: "hello (up) WORLD (low) test (cap)",
			want:  "HELLO world Test",
		},

		// Tests avec plusieurs mots
		{
			name:  "up multiple words",
			input: "hello world (up, 2)",
			want:  "HELLO WORLD",
		},
		{
			name:  "low multiple words",
			input: "HELLO WORLD (low, 2)",
			want:  "hello world",
		},
		{
			name:  "cap multiple words",
			input: "hello WORLD (cap, 2)",
			want:  "Hello World",
		},
		{
			name:  "up three words",
			input: "one two three (up, 3)",
			want:  "ONE TWO THREE",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := processLevel2(test.input)

			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}
