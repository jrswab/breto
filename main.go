package main

import (
	"fmt"
	"gitlab.com/jrswab/dwm-status/libs/ram"
	"gitlab.com/jrswab/dwm-status/libs/wttr"
	"os/exec"
	"time"
)

func main() {
	var status, hTime, weather, ramFree string
	var cmd *exec.Cmd

	cWttr := make(chan string) // start weather data routine
	go wttr.Local(cWttr)

	cRam := make(chan string) // start free ram data routine
	go ram.Free(cRam)

	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		hTime = time.Now().Format("Jan 02, 2006 15:04") // add seconds with "Jan 02, 2006 15:04:05"

		select { // update the go routine channels as they send data
		case weather = <-cWttr:
		case ramFree = <-cRam:
		default:
		}

		status = fmt.Sprintf(" %s%s%s ", ramFree, weather, hTime) // Change by editing variables & `%s`
		cmd = exec.Command("xsetroot", "-name", status)
		cmd.Run()
	}
}
