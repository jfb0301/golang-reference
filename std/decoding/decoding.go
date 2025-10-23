package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"Age"`
}

func main() {
    data := `{"name":"Fred","Age":40}
{"name":"Camille","Age":41}
{"name":"William","Age":42}`

    dec := json.NewDecoder(strings.NewReader(data))
    var people []Person
    for dec.More() {
        var t Person
        err := dec.Decode(&t)
        if err != nil {
            panic(err)
        }
        people = append(people, t) // Process: collect in a slice
    }

    // Print all decoded people
    for _, p := range people {
        fmt.Printf("%+v\n", p)
    }
}
