package blocks

import (
	"fmt"
	"testing"
)

// A basic test to make sure the function runs
// and outputs data over instead of an error
func TestWttr(t *testing.T) {
	cWttr := make(chan string)
	eWttr := make(chan error)
	var weather string
	var wttrError error

	go Wttr(cWttr, eWttr)

	// wait for weather or error
	select {
	case weather = <-cWttr:
	case wttrError = <-eWttr:
	}

	if wttrError != nil {
		t.Error("Expected Weather Data got:", wttrError.Error())
	} else if weather == " | " {
		t.Error("Weather channel sent an empty string.")
	}
	fmt.Println(weather)
}
