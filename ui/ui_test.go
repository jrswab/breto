package ui

import (
	"fmt"
	"testing"
)

// Test DWM xsetroot -name command
func TestDwm(t *testing.T) {
	status := Dwm("Test Status")
	if status != nil {
		t.Error("Expected 'Test Status' but got:", status.Error())
	}
	fmt.Println(status)
}
