package main

import (
    "testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
	input string
	expected []string
    }{
	{
	    input: "  hello world  ",
	    expected: []string{"hello", "world"},
	},
	{
	    input: "Charmander Bulbasaur PIKACHU",
	    expected: []string{"charmander", "bulbasaur", "pikachu"},
	},
    }

    for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
	    t.Errorf("lengths don't match: %v, %v", actual, c.input)
	    continue
	}

	for i := range actual {
	    word := actual[i]
	    expectedWord := c.expected[i]

	    if word != expectedWord {
		t.Errorf("word don't match %v, %v", word, expectedWord)
	    }
	}
    }
}
