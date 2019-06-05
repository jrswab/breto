package blocks

import (
	"fmt"
	"os/exec"
	"strings"
)

func VolumeText() (string, error) {
	volCmd := "amixer -D pulse sget Master | awk '/Front Right:/ {print $5}' | grep -o '[0-9].'"
	runVol, err := exec.Command("sh", "-c", volCmd).Output()
	if err != nil {
		return "", err
	}

	percent := "%"
	return fmt.Sprintf("%s%s |", strings.TrimSpace(string(runVol)), percent), nil
}
