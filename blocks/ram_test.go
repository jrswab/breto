package blocks

import (
	"fmt"
	"testing"
)

// A basic test to make sure the function runs
// and outputs data instead of an error
func TestFreeRam(t *testing.T) {
	cRam := make(chan string)
	eRam := make(chan error)
	var ram string
	var runError error

	go FreeRam(cRam, eRam)

	select {
	case ram = <-cRam:
	case runError = <-eRam:
	}

	if runError != nil {
		t.Error("Expected Free Ram got:", runError.Error())
	}
	fmt.Println(ram)
}
