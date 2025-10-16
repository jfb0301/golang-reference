package main 

import "fmt"

// Reading from an empty interface 

type MyInt int 

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine 
	i2, ok := i.(string)	// If type don't coincide with value - panic: interface conversion: interface{} is not string
	if !ok {
		return fmt.Errorf("Unexpected type for", i)
	}
	fmt.Println(i2 + 1)
}



