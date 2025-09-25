package main 

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
	fmt.Println(i, v)
	}

	// Ignoring key 

	evenVals2 := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals2 {
		fmt.Println(v)
	}

	// Ignoring value 

	uniqueNames := map[string]bool{"Fred" : true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}

