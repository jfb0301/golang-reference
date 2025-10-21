package main 

import "fmt"

for {
	select {
	case v, ok := <-in:
		if !ok {
			in = nil // The case will never succeed again!
			continue
		}
		// Process the v that was read from in 
	case v. ok := <-in2:
		if !ok {
			in2 = nil 
			continue 
		}
		// Process the v that was read from in2
	case <-done:
		return 
	}
}


