package main

import(
	"fmt"
	"math/rand"
)

func main() {
	// Scoping to an if statement 

	if n := rand.Intn(5); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("Thats too big")
	} else {
		fmt.Println("Thats a good number", n)
	}

	// Out of scope 

	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big: ", n)
	} else {
		fmt.Println("That's too low: ", n)
	}
	fmt.Println(n)		// Out of scope 
}

