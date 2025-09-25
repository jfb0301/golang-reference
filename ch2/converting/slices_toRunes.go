package main 

import "fmt"

func main(){
	var s string = "Hello, ☉"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)

	fmt.Println(bs)		// [72 101 108 108 111 44 32 226 152 137]
	fmt.Println(rs)		// [72 101 108 108 111 44 32 9737]
}