package main

import (
	"fmt" //"os/exec"
	"gitlab.com/jrswab/dwm-status/libs/ram"
	"gitlab.com/jrswab/dwm-status/libs/wttr"
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

		// store desired items as strings
		// delete or comment out a line to remove from status bar
		cat := []string{"", // do not delete or comment out this line
			"RAM:", ramFree, "free", "|",
			weather, "|",
			hTime,
			""} // do not delete or comment out this line

		// concatinate all strings to one line for output
		status := strings.Join(cat, " ")

		//cmd := exec.Command("xsetroot", "-name", status)
		//cmd.Run()
		fmt.Println(status)

		time.Sleep(1 * time.Second)

		// reset i to loop for time delayed updates
		if i > 3600 {
			i = 0
		} else {
			i++
		}
	}
}
