package main 

import "fmt"

func TimeLimit() (int, error) {
	var result int
	var error int 
	done := make(chan struct{})

	go func() {
		result, err := doSomeWork() 
		close(done)
	}()
	select {
	case <- done:
		return result, err
	case <- time.After(2 * time.Second):
		return 0, errors.New("Work time out")
	}
}