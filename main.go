package main

import (
	"fmt"
	"gitlab.com/jrswab/dwm-status/libs/ram"
	"gitlab.com/jrswab/dwm-status/libs/wttr"
	//"os/exec"
	"strings"
	"time"
)

func main() {
	// initial run
	var i uint16 = 0
	weather := strings.TrimSpace(string(*wttr.Local()))
	ramFree := ram.Free()

	for i < 3700 {
		// to add seconds use "Jan 02, 2006 15:04:05"
		hTime := time.Now().Format("Jan 02, 2006 15:04")
		// time delayed retrievals:
		// get weather once per hour
		if i == 3600 {
			weather = strings.TrimSpace(string(*wttr.Local()))
		}
		// get free Ram every 3 seconds
		if i%3 == 0 {
			ramFree = ram.Free()
		}

		// Change what is displayed by adding or removing variables (eg. weather).
		// Change the text output by editing bewteen the `""`.
		// Make sure to have the same number of `%s` as variables.
		status := fmt.Sprintf(" RAM: %s free | %s | %s ", ramFree, weather, hTime)

		cmd := exec.Command("xsetroot", "-name", status)
		cmd.Run()

		time.Sleep(1 * time.Second)

		// reset i to loop for time delayed updates
		if i > 3600 {
			i = 0
		} else {
			i++
		}
	}
}
