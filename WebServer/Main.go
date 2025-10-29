package main 

import(
	"context"
	"os"

	"example.com/WebServer/clients"
	"example.com/WebServer/servers"
	
)

func main() {
    ss := servers.SlowServer() 
    defer ss.Close()
    fs := servers.FastServer() 
    defer fs.Close()

    ctx := context.Background()

    // Check if argument is provided
    if len(os.Args) < 2 {
        println("Usage: go run . <error>")
        println("Example: go run . true")
        return
    }

    clients.CallBoth(ctx, os.Args[1], ss.URL, fs.URL)
}