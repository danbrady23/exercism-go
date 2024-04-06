// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

const (
	stanza = "For want of a %s the %s was lost."
	last   = "And all for the want of a %s."
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) (output []string) {

	nLines := len(rhyme)

	if nLines == 0 {
		return
	}

	for i := 0; i < nLines-1; i++ {
		output = append(output, fmt.Sprintf(stanza, rhyme[i], rhyme[i+1]))
	}
	return append(output, fmt.Sprintf(last, rhyme[0]))
}
