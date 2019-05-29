package ram

import (
	"fmt"
	"os/exec"
	"strings"
)

func Free(cRam chan string) {
	ramFree := ""
	// set shell command
	ramCmd := "free -h | gawk '/Mem:/ {print $4}'"
	// run and save the output of ramCmd to ramGib
	ramGib, err := exec.Command("sh", "-c", ramCmd).Output()
	if err != nil {
		cRam <- err.Error()
	}
	// return ramGib as string
	ramFree = fmt.Sprintf("Ram: %s free | ", strings.TrimSpace(string(ramGib)))
	cRam <- ramFree
}
