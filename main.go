package main

import (
	"fmt"
	"gitlab.com/jrswab/go-status/blocks" // if pulled from github change to github
	"gitlab.com/jrswab/go-status/ui"     // if pulled from github change to github
	"log"
	"time"
)

func main() {
	// Other Blocks:
	var status, hTime, weather, ramFree, rShift, dropbox, volIcon, volText string
	var wttrErr, ramErr error

	cWttr := make(chan string) // start weather data routine
	eWttr := make(chan error)
	go blocks.Wttr(cWttr, eWttr)

	cRam := make(chan string) // start free ram data routine
	eRam := make(chan error)
	go blocks.FreeRam(cRam, eRam)

	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		// add seconds with "Jan 02, 2006 15:04:05"
		hTime = time.Now().Format("Jan 02, 2006 15:04")

		select { // update the go routine channels as they send data
		case weather = <-cWttr:
		case wttrErr = <-eWttr:
			log.Println(wttrErr.Error())
		case ramFree = <-cRam:
		case ramErr = <-eRam:
			log.Println(ramErr.Error())
		default:
		}

		// Assign Icons & non Go Routine blocks every round
		rShift, _ = blocks.RedshiftIcon()
		dropbox, _ = blocks.DropboxIcon()
		volText, _ = blocks.VolumeText()
		volIcon, _ = blocks.VolumeIcon()

		// Change by editing variables & `%s`
		status = fmt.Sprintf(" %s%s%s%s%s %s%s",
			ramFree, weather, volIcon, volText, hTime, dropbox, rShift)
		ui.Dwm(status) // change this to the UI of choice
	}
}
