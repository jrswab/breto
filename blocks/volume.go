package blocks

import (
	"fmt"
	"os/exec"
	"strings"
)

// Volume sends back the current volume percent.
func Volume(emoji bool) (string, error) {
	volCmd := "pulsemixer --get-volume | awk '{print $1}'"
	runVol, err := exec.Command("sh", "-c", volCmd).Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s", strings.TrimSpace(string(runVol)), "%"), nil
}
