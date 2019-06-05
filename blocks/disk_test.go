package blocks

import (
	"fmt"
	"testing"
)

// A basic test to make sure the function runs
// and outputs data instead of an error
func TestHomeDisk(t *testing.T) {
	cHomeDisk := make(chan string)
	eHomeDisk := make(chan error)
	var homeFree string
	var runError error

	go HomeDisk(cHomeDisk, eHomeDisk)

	select {
	case homeFree = <-cHomeDisk:
	case runError = <-eHomeDisk:
	}

	if runError != nil {
		t.Error("Expected Free Home Directory Space got:", runError.Error())
	}
	fmt.Println(homeFree)
}
