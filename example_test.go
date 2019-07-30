package rabinkarp_test

import (
	"fmt"
	"rabinkarp"
	"sort"
)

func Example() {
	pattern := "bro"
	inputs := []string{"slowbro", "browser", "bob", "st", "bebrobaum"}
	matcher := rabinkarp.Make(rabinkarp.NaiveHasher)
	matches := matcher.MatchAll(pattern, inputs)

	sort.Sort(matches)
	fmt.Println(matches)

	// Output:
	// [{0 browser} {2 bebrobaum} {4 slowbro}]
}
