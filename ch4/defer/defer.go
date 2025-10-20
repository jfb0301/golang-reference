package main 

import(
	"fmt"
	"os"
	"log"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No file is specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != io.EOF {
			log.Fatal(err)
		}
		break 
	}
}