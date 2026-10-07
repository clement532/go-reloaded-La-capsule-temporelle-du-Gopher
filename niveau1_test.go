package main

import "testing"

func TestProcessLevel1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "hexadecimal",
			input: "1E (hex) files were added",
			want:  "30 files were added",
		},
		{
			name:  "binary",
			input: "It has been 10 (bin) years",
			want:  "It has been 2 years",
		},
		{
			name:  "hexadecimal lowercase",
			input: "ff (hex) red",
			want:  "255 red",
		},
		{
			name:  "invalid hexadecimal",
			input: "We need ZZ (hex) more",
			want:  "We need ZZ more",
		},
		{
			name:  "invalid binary",
			input: "Call 102 (bin) now",
			want:  "Call 102 now",
		},
		{
			name:  "multiple conversions",
			input: "Simply add 42 (hex) and 10 (bin) and you will see",
			want:  "Simply add 66 and 2 and you will see",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := processLevel1(test.input)

			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}