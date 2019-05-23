package ram

import (
	"os/exec"
	"strings"
)

func Free() string {
	ramCmd := "free -h | gawk '/Mem:/ {print $4}'"
	ramGib, err := exec.Command("sh", "-c", ramCmd).Output()
	if err != nil {
		return "Error in 'free -h' command."
	}
	return strings.TrimSpace(string(ramGib))
}
