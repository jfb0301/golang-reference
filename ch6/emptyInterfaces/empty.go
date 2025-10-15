package main 

import "fmt"

data := map[string]interface{}
contents, err := ioutil.ReadFile("testdata/sample.json")
if err != nil {
	return err 
}

defer contents.Close()

json.Unmarshal(contents, &data)

// The contents are now in the data map 

