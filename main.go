package main

import (
	"gitlab.com/jrswab/dwm-status/libs/ram"
	"gitlab.com/jrswab/dwm-status/libs/wttr"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// initial run
	i := 0
	weather := wttr.Local()
	ramFree := ram.Free()

	for i < 3700 {
		hTime := time.Now().Format("Jan 02, 2006 15:04")
		// time delayed retrievals:
		// get weather once per hour
		if i == 3600 {
			weather = wttr.Local()
		}
		// get free Ram every 3 seconds
		if i%3 == 0 {
			ramFree = ram.Free()
		}

		// store desired items as strings
		// delete or comment out a line to remove from status bar
		cat := []string{"",
			"RAM:", ramFree, "free", "|",
			weather, "|",
			hTime,
			""}

		// concatinate all strings to one line for output
		status := strings.Join(cat, " ")

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
