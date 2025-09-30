package main 

import(
	"fmt"
	"strconv"
)

func add(i int, j int) int { return i + j}

func sub(i int, j int) int { return i - j}

func mul(i int, j int) int { return i * j}

func div(i int, j int) int { return i / j}


// Map: Associates a math operator to each function 

var opMap = map[string]func(int, int) int {
	"+" : add,
	"-" : sub,
	"*" : mul,
	"/" : div,
}


func main() {
	
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"3", "-", "5"},
		[]string{"3" ,"*", "5"},
		[]string{"10", "/", "2"},
		[]string{"4" ,"-", "1"},
		[]string{"3" ,"*", "5"},
		[]string{"3" ,"%", "5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expression: ", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("Unsupported operator", op)
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}