package main

import (
	"fmt"
	"gitlab.com/jrswab/go-status/blocks"
	"log"
	"os/exec"
	"time"
)

func main() {
	var status, hTime, weather, ramFree string
	var wttrErr, ramErr error
	var cmd *exec.Cmd

	cWttr := make(chan string) // start weather data routine
	eWttr := make(chan error)
	go blocks.Wttr(cWttr, eWttr)

	cRam := make(chan string) // start free ram data routine
	eRam := make(chan error)
	go blocks.FreeRam(cRam, eRam)

	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		hTime = time.Now().Format("Jan 02, 2006 15:04") // add seconds with "Jan 02, 2006 15:04:05"

		select { // update the go routine channels as they send data
		case weather = <-cWttr:
		case wttrErr = <-eWttr:
			log.Println(wttrErr.Error())
		case ramFree = <-cRam:
		case ramErr = <-eRam:
			log.Println(ramErr.Error())
		default:
		}

		status = fmt.Sprintf(" %s%s%s ", ramFree, weather, hTime) // Change by editing variables & `%s`
		cmd = exec.Command("xsetroot", "-name", status)
		cmd.Run()
	}
}
