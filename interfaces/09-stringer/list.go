// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strings"
)

type list []fmt.Stringer

// list satisfies the fmt.Stringer
func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚."
	}

	// use strings.Builder when you're combining unknown
	// list of strings together.
	var str strings.Builder

	for _, it := range l {
		// the builder doesn't know about the stringer interface.
		// that's why you need to call String method here.
		str.WriteString(it.String())
		str.WriteRune('\n')
	}

	return str.String()
}

func (l list) discount(ratio float64) {
	type discounter interface {
		discount(float64)
	}

	for _, it := range l {
		if it, ok := it.(discounter); ok {
			it.discount(ratio)
		}
	}
}
