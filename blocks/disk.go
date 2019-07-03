package blocks

import (
	"fmt"
	"math"
	"os/exec"
	"strings"
	"time"
)

func HomeDisk(cHomeDisk chan string, eHomeDisk chan error) {
	// df -h | awk '/home/ {print $4}
	var passed, hour float64
	var homeFree string
	homeCmd := "df -Ph .| awk '/d*G/ {print $4}'"
	start := time.Now()
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		passed = time.Since(start).Seconds()
		hour = math.Floor(math.Remainder(passed, 3600))

		if passed < 5 || hour == 0 {
			homeOut, err := exec.Command("sh", "-c", homeCmd).Output()
			if err != nil {
				eHomeDisk <- err
			}

			homeFree = fmt.Sprintf("%s |", strings.TrimSpace(string(homeOut)))
			cHomeDisk <- homeFree
		}
	}
}
