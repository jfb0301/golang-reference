package main

import "fmt"

func main() {
	n := map[string]int{
		"a" : 1, 
		"b" : 2, 
		"c" : 3, 
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i) 
			for k, v := range n {
				fmt.Println(k, v)
			}
	}


}