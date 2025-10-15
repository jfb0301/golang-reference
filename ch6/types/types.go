package main 

import "fmt"

type LogicProvider struct { }		// Nothing declared here indicates that it meets the interf

func (lp LogicProvider) Process(data string) string {
	// Business Logic 
}

type Logic interface {			// Interfaces specify what the client needs 
	Process(data string) string 
}

type Client struct {		// Only the caller (Client) knows about the interface
	L Logic 
}

func(c Client) program() {
	// Get data from somewhere 
	c.L.Process(data)
}

func main() {
	c := client{
		L : LogicProvider{},
	}
	c.Program()
}

