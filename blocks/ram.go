package blocks

import (
	"fmt"
	"math"
	"os/exec"
	"strings"
	"time"
)

func FreeRam(cRam chan string, eRam chan error) {
	var passed, tenSecs float64
	start := time.Now() // set to determine seconds passed
	ticker := time.NewTicker(time.Second)

	for range ticker.C { // run every 10 seconds
		passed = time.Since(start).Seconds() // get total seconds passed
		tenSecs = math.Floor(math.Remainder(passed, 10))

		if passed < 5 || tenSecs == 0 { // trigger: asap or divisible by ten
			ramFree := ""
			ramCmd := "free -h | gawk '/Mem:/ {print $4}'" // set shell command

			ramGib, err := exec.Command("sh", "-c", ramCmd).Output() // run and save the output
			if err != nil {
				eRam <- err
			}

			ramFree = fmt.Sprintf("Ram: %s free | ",
				strings.TrimSpace(string(ramGib))) // set string

			cRam <- ramFree // send string
		}
	}
}
