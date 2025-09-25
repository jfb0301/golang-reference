package main 

import "fmt"

func main() {
	// Declaring and initializing a map using a map literal 

	ages := map[string]int{
		"Alice" : 30,
		"Bob" : 305,
		"Alices" : 50,
	}

	fmt.Println(ages)

	// Map literal with struct 

	type City struct {
		Population int
		Country string
	}

	cities := map[string]City{
		"London" : {Population: 9000000, Country : "UK"},
		"Paris" : {Population: 9000000, Country : "France"},
	}

	fmt.Println(cities)
}