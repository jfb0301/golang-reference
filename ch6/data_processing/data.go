package main 

import(
	"fmt"
	"os"
)

func process(r io.Reader) error 

r, err := os.Open(fileName) // Meets io.Reader interface 
if err != nil {
	return err
}

defer r.Close()
return process(r)
return nil 


// We can wrap io.Reader into another io.Reader 

defer r.Close()
gz, err := gzip.NewReader(r)
if err != nil {
	return err 
}

defer gz.Close()
return process(gz)