package ram

import (
	"os/exec"
	"strings"
)

func Free() string {
	// set shell command
	ramCmd := "free -h | gawk '/Mem:/ {print $4}'"
	// run and save the output of ramCmd to ramGib
	ramGib, err := exec.Command("sh", "-c", ramCmd).Output()
	if err != nil {
		return err.Error()
	}
	// return ramGib as string
	return strings.TrimSpace(string(ramGib))
}
