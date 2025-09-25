package main 

import "fmt"

func main() {
	samples := []string{"hello", "Apple_n!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}