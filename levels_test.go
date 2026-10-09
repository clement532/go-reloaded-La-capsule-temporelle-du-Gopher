package main

import "testing"

func TestLevel3(t *testing.T) {
	check(t, false, [][2]string{
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"Really ! ? No way", "Really!? No way"},
		{"Ingredients : flour , eggs ; milk", "Ingredients: flour, eggs; milk"},
		{"Pi is 3.14 and 1,000 is big", "Pi is 3.14 and 1,000 is big"},
		{"e.g. this,that", "e.g. this,that"},
		{"Punctuation tests are ... kinda boring ,what do you think ?",
			"Punctuation tests are... kinda boring, what do you think?"},
		{"too   many \t spaces  ", "too many spaces"},
		{"", ""},
	})
}

func TestLevel4(t *testing.T) {
	check(t, false, [][2]string{
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"As Yoda said: ' Do or do not , there is no try '",
			"As Yoda said: 'Do or do not, there is no try'"},
		{"He said ' hello ' !", "He said 'hello'!"},
		{"Don't touch it , it's ' fragile '", "Don't touch it, it's 'fragile'"},
		{"' one ' and ' two '", "'one' and 'two'"},
		{"a lonely ' quote", "a lonely ' quote"},
	})
}

func TestLevel5(t *testing.T) {
	check(t, false, [][2]string{
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"There is no greater agony than bearing a untold story inside you.",
			"There is no greater agony than bearing an untold story inside you."},
		{"She ate an banana and a orange", "She ate a banana and an orange"},
		{"It was a ' epic ' fail", "It was an 'epic' fail"},
		{"Give me a", "Give me a"},
		{"a house and a hour", "a house and a hour"},
		{"a 8-bit game", "a 8-bit game"},
		{"a (up) apple", "An apple"},
	})
}

func TestAllTogether(t *testing.T) {
	in := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, " +
		"it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, " +
		"it was the season of Light, it was the season of darkness, it was the spring of hope, " +
		"IT WAS THE (low, 3) winter of despair."
	want := "It was the best of times, it was the worst of TIMES, it was the age of wisdom, " +
		"It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, " +
		"it was the season of Light, it was the season of darkness, it was the spring of hope, " +
		"it was the winter of despair."
	check(t, false, [][2]string{{in, want}})
}

func TestSmart(t *testing.T) {
	check(t, true, [][2]string{
		{"olleh (rev) dlrow (rev) !", "hello world!"},
		{"olleh dlrow (rev, 2)", "hello world"},
		{"a hour , an university and a honest unicorn", "an hour, a university and an honest unicorn"},
	})
	// sans --smart, (rev) est du texte ordinaire
	check(t, false, [][2]string{
		{"olleh (rev) dlrow", "olleh (rev) dlrow"},
	})
}
