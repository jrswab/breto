package main

import (
	"fmt"
	"gitlab.com/jrswab/dwm-status/libs/ram"
	"gitlab.com/jrswab/dwm-status/libs/wttr"
	//"os/exec"
	"time"
)

func main() {
	// initial run
	weather := ""
	ramFree := ""

	// get weather data
	c := make(chan string)
	go wttr.Local(c)

	// get free ram data
	cRam := make(chan string)
	go ram.Free(cRam)

	for {
		// to add seconds use "Jan 02, 2006 15:04:05"
		hTime := time.Now().Format("Jan 02, 2006 15:04")

		// Change what is displayed by adding or removing variables (eg. weather).
		// Change the text output by editing bewteen the `""`.
		// Make sure to have the same number of `%s` as variables.
		status := fmt.Sprintf(" %s%s%s ", ramFree, weather, hTime)

		//cmd := exec.Command("xsetroot", "-name", status)
		//cmd.Run()
		fmt.Println(status)

		time.Sleep(1 * time.Second)

		// set channel data for next loop
		ramFree = <-cRam
		weather = <-c
	}
}
