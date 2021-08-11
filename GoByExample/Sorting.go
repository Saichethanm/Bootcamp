package main

import (
	"fmt"
	"sort"
)

func main() {

	// stringSlice := []string{"c", "a", "b", "d"}
	stringSlice := []string{"hwabcd", "helloworld", "hsomething", "hletssee"}
	intSlice := []int{5, 4, 3, 2, 1}

	sort.Strings(stringSlice)
	fmt.Println(stringSlice)
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	fmt.Println(sort.IntsAreSorted(intSlice))
}
