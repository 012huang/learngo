// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

type game struct {
	title    string
	price    money
	playTime int
}

func (g *game) print() {
	fmt.Printf("%-15s: %s\n", g.title, g.price.string())
}

// func (g *game) String() string {
// 	return fmt.Sprintf("%-15s: $%.2f", g.title, g.price)
// }

func (g *game) discount(ratio float64) {
	g.price *= money(1 - ratio)
}

// TODO: NEW
func (g *game) sum() money {
	return g.price
}
