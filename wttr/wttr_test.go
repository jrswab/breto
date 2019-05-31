package wttr

import (
	"fmt"
	"testing"
)

// A basic test to make sure the function runs
// and outputs data over instead of an error
func TestLocal(t *testing.T) {
	cWttr := make(chan string)
	eWttr := make(chan error)
	var weather string
	var wttrError error

	go Local(cWttr, eWttr)

	// wait for weather or error
	select {
	case weather = <-cWttr:
	case wttrError = <-eWttr:
	}

	if wttrError != nil {
		t.Error("Expected Weather Data got:", wttrError.Error())
	}
	fmt.Println(weather)
}
