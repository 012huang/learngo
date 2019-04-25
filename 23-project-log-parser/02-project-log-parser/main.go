// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		sum     map[string]int // total visits per domain
		domains []string       // unique domain names
		total   int            // total visits to all domains
		lines   int            // number of parsed lines (for the error messages)
	)

	sum = make(map[string]int)

	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		lines++

		// Parse the fields
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", fields, lines)
			return
		}

		// Sum the total visits per domain
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", fields[1], lines)
			return
		}

		name := fields[0]

		// Collect the unique domains
		if _, ok := sum[name]; !ok {
			domains = append(domains, name)
		}

		// Keep track of total and per domain visits
		total += visits
		sum[name] += visits
	}

	// Print the visits per domain
	sort.Strings(domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, name := range domains {
		visits := sum[name]
		fmt.Printf("%-30s %10d\n", name, visits)
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)
}
