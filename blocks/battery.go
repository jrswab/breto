package blocks

import (
	"fmt"
	"os/exec"
	"strings"
)

// Battery returns the current percentage remaining.
func Battery() (string, error) {
	cmd := "cat /sys/class/power_supply/BAT0/capacity"
	runCmd, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil && err.Error() != "exit status 1" {
		return "", err
	}

	perc := "%"
	level := fmt.Sprintf("%s%s", strings.TrimSpace(string(runCmd)), perc)
	return level, nil
}
