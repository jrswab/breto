package main

import (
	"fmt"
	"gitlab.com/jrswab/go-status/blocks"
	"gitlab.com/jrswab/go-status/ui"
	"log"
	"time"
)

func main() {
	var status, hTime, weather, ramFree string
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

		// Change by editing variables & `%s`
		status = fmt.Sprintf(" %s%s%s ", ramFree, weather, hTime)
		ui.Dwm(status) // change this to the UI of choice
	}
}
