// Wrapping errors with defer

package main 

import "fmt"

func DoSomeThings(val1 int, val2 int) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThings: %w", err)
		}
	}()
	val3, err := doThin1(val1)
	if err != nil {
		return "", err
	}
	return doThing3(val3, val4)
}