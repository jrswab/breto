package main

import (
	"fmt"
	"gitlab.com/jrswab/dwm-status/libs/ram"
	"gitlab.com/jrswab/dwm-status/libs/wttr"
	"math"
	"os/exec"
	"time"
)

func main() {
	// initial run
	start := time.Now()
	var weather string
	var ramFree string
	hTime := time.Now().Format("Jan 02, 2006 15:04")
	status := fmt.Sprintf(" %s ", hTime)
	// run first execution without the go routines
	cmd := exec.Command("xsetroot", "-name", status)
	cmd.Run()

	// get weather data
	cWttr := make(chan string)
	go wttr.Local(cWttr)

	// get free ram data
	cRam := make(chan string)
	go ram.Free(cRam)

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		// get number of seconds since execution
		tPassed := time.Since(start).Seconds()
		// to add seconds use "Jan 02, 2006 15:04:05"
		hTime := time.Now().Format("Jan 02, 2006 15:04")

		// Change what is displayed by adding or removing variables (eg. weather).
		// Change the text output by editing bewteen the `""`.
		// Make sure to have the same number of `%s` as variables.
		// Using select here to update the go routine channels as they send data
		select {
		case weather = <-cWttr:
			status = fmt.Sprintf(" %s%s%s ", ramFree, weather, hTime)
		case ramFree = <-cRam:
			status = fmt.Sprintf(" %s%s%s ", ramFree, weather, hTime)
		default:
			status = fmt.Sprintf(" %s%s%s ", ramFree, weather, hTime)
		}

		// Update output
		cmd = exec.Command("xsetroot", "-name", status)
		cmd.Run()

		// round down to nearest int, and use the number to set refresh rate per
		// library. This is needed to not over use wttr.in
		fiveSec := math.Floor(math.Remainder(tPassed, 5))
		hourPassed := math.Floor(math.Remainder(tPassed, 3600))

		if fiveSec == 0 {
			go ram.Free(cRam)
		}
		if hourPassed == 0 {
			go wttr.Local(cWttr)
		}
	}
}
