package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

// Assuming Person is a struct like this
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    // Sample data to encode
    toFile := Person{Name: "Alice", Age: 30}

    // Create a temporary file
    tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
    if err != nil {
        panic(err)
    }
    defer os.Remove(tmpFile.Name()) // Clean up the file after use

    // Encode Person struct to JSON and write to file
    err = json.NewEncoder(tmpFile).Encode(toFile)
    if err != nil {
        panic(err)
    }

    // Close the file
    err = tmpFile.Close()
    if err != nil {
        panic(err)
    }

    // Reopen the file for reading
    tmpFile2, err := os.Open(tmpFile.Name())
    if err != nil {
        panic(err)
    }

    // Decode JSON from file into Person struct
    var fromFile Person
    err = json.NewDecoder(tmpFile2).Decode(&fromFile) // Fixed: NewDecoder, not NewEncoder
    if err != nil {
        panic(err)
    }

    // Close the file
    err = tmpFile2.Close()
    if err != nil {
        panic(err)
    }

    // Print the decoded struct
    fmt.Printf("%+v\n", fromFile)
}