package main 

import  "fmt"

func main() {
	select {
	case v := <- ch:
		fmt.Println(v)
	case v := <- ch2:
		fmt.Println(v)
	case ch3 := x:
		fmt.Println("wrote", x)
	case <- ch4:
		fmt.Println("Got value on ch4, but ignored it.") 
	}
}