package main

import (
	"sort"
	"fmt"
)

func Sorting() {

	// Sort methods are specific to the buildin type,
	// Note that sorting is in-place,
	// so it changes the given slice and doesn't return a new one.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints	", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

}