package wttr

import (
	"os/exec"
	"strings"
)

func Local() string {
	wttrCmd := "curl -s 'wttr.in/?format=%t+%w'"
	wttrRun, err := exec.Command("sh", "-c", wttrCmd).Output()
	if err != nil {
		return "wttr command returned error!"
	}
	return strings.TrimSpace(string(wttrRun))
}
