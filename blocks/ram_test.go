package blocks

import (
	"os/exec"
	"strings"
	"testing"
)

// A basic test to make sure the function runs
// and outputs data instead of an error
func TestFreeRAM(t *testing.T) {
	cRAM := make(chan string)
	eRAM := make(chan error)
	var got string
	var runError error

	go FreeRam(cRAM, eRAM)

	ramCmd := "free -h | gawk '/Mem:/ {print $7}'"             // set shell command
	expected, err := exec.Command("sh", "-c", ramCmd).Output() // run and save the output
	if err != nil {
		eRAM <- err
	}

	select {
	case got = <-cRAM:
	case runError = <-eRAM:
	}

	if runError != nil {
		t.Error("Expected Free Ram got:", runError.Error())
	}

	if strings.TrimSpace(string(expected)) != got {
		t.Errorf("\nExpected %s but got %s", expected, got)
	}
	return
}
