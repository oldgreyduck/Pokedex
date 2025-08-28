package main

import "testing"

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
    	{
        	input:    "  hello  world  ",
        	expected: []string{"hello", "world"},
    	},
		{
	    	input:    " charmander BULBAsaur PIKACHU   ",
	    	expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
	    	input:    "Magic The GATHERING",
	    	expected: []string{"magic", "the", "gathering"},
		},
    }
		
    for _, c := range cases {
		actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("For input '%s': expected %d, got %d", c.input, len(c.expected), len(actual))
    }
    for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
	    	t.Errorf("For input '%s' at index %d: expected %v, got %v", c.input, i, expectedWord, word)
    	}
	}
    }
}

