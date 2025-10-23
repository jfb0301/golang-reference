package main

import (
    "bytes"
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"Age"`
}

func process(input Person) Person {
    // Example processing: increment age
    return Person{Name: input.Name, Age: input.Age + 1}
}

func main() {
    allInputs := []Person{
        {Name: "Fred", Age: 39},
        {Name: "Camille", Age: 40},
        {Name: "William", Age: 41},
    }

    var b bytes.Buffer
    enc := json.NewEncoder(&b)
    for _, input := range allInputs {
        t := process(input)
        err := enc.Encode(t)
        if err != nil {
            panic(err)
        }
    }

    out := b.String()
    fmt.Println(out)
}