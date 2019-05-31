package ram

import (
	"fmt"
	"testing"
)

// A basic test to make sure the function runs
// and outputs data instead of an error
func TestFree(t *testing.T) {
	cRam := make(chan string)
	eRam := make(chan error)
	var ram string
	var runError error

	go Free(cRam, eRam)

	select { // grab first avalible channel
	case ram = <-cRam:
	case runError = <-eRam:
	}

	if runError != nil {
		t.Error("Expected Free Ram got:", runError.Error())
	}
	fmt.Println(ram)
}
