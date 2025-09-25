package main

import "fmt"

// Using a map as set

func main() {
	intSet := map[int]bool{}					// Initialize map 
	vals := []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 100}	// Populate the set 
	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in a set")
	}
}