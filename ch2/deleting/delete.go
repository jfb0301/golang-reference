package main 

import "fmt"


func main() {
	m := map[string]int {
		"Hello" : 5,
		"World" : 10, 
	}

	fmt.Println(m)
	delete(m, "Hello")
	fmt.Println(m)
}