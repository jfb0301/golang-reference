package main 

import "fmt"

type MyFuncOps struct {
	FirstName string
	LastName string
	Age int 
}

func MyFunc(opts MyFuncOps) error {
	// Do something 
}

func main() {
	MyFunc(MyFuncOps {
		LastName : "Pastel",
		Age : 50,
	})

	MyFunc(MyFuncOps {
		FirstName : "Joe",
		LastName : "Bianchi",
	})
	
}