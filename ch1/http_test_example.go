package main

import (
    "fmt"
    "os/exec"
)

func main() {
    fmt.Println("Starting load test with hey...")

    // Run hey command
    cmd := exec.Command("hey", "-n", "100", "-c", "10", "https://nytimes.com")
    // -n 100: 100 total requests
    // -c 10: 10 concurrent requests

    // Capture output
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Error running hey: %v\n", err)
        return
    } 

    // Print the output
    fmt.Println(string(output))
}