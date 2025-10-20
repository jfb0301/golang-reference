package main 

import "fmt"

func process(val int) int {
	// do something with val 
}


func runThingConcurrently(in <- chan int, out chan<- int) {
	go func() {
		for val := range in {
			result := process(val)
			out <- result
		}
	} ()
}

