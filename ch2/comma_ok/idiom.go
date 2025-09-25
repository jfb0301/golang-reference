package main

import "fmt"

func main() {
	m := map[string]int {
		"Hello " : 5,
		"World " : 10, 
	}
	v, ok := m["Hello"]
	fmt.Println(v, ok)

	v, ok = m["World"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)
}