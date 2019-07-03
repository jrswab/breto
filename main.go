package main

import (
	"fmt"
	"github.com/jrswab/go-status/blocks" // if pulled from github change to github
	"github.com/jrswab/go-status/ui"     // if pulled from github change to github
	"log"
	"math"
	"time"
)

func main() {
	var status, hTime, battery string
	var weather, ramFree, homeSpace string                        // Go routine blocks
	var wttrErr, ramErr, homeErr error                            // Go routine errors
	var rShift, dropbox, volIcon, volText, syncthing, bolt string // Icon blocks
	var passed, fiveMins float64

	// These are static icons and only need defined at the start
	homeDir := blocks.DirIcon()
	memIco := blocks.MemIcon()
	tempIco := blocks.TempIcon()

	// Each Go routine has it's own timer to delay the execution of the command.
	cWttr := make(chan string) // start weather data routine
	eWttr := make(chan error)
	go blocks.Wttr(cWttr, eWttr)

	cRam := make(chan string) // start free ram data routine
	eRam := make(chan error)
	go blocks.FreeRam(cRam, eRam)

	cHomeDisk := make(chan string) // start free home dir space routine
	eHomeDisk := make(chan error)
	go blocks.HomeDisk(cHomeDisk, eHomeDisk)

	ticker := time.NewTicker(time.Second)
	start := time.Now()
	for range ticker.C {
		// add year & seconds with "Jan 02, 2006 15:04:05"
		hTime = time.Now().Format("Jan 02 15:04")
		passed = time.Since(start).Seconds()
		fiveMins = math.Floor(math.Remainder(passed, 300))

		select { // update the go routine channels as they send data
		case weather = <-cWttr:
		case wttrErr = <-eWttr:
			log.Println(wttrErr.Error())
		case ramFree = <-cRam:
		case ramErr = <-eRam:
			log.Println(ramErr.Error())
		case homeSpace = <-cHomeDisk:
		case homeErr = <-eHomeDisk:
			log.Println(homeErr.Error())
		default:
		}

		// Assign Icons & non Go Routine blocks every round
		rShift, _ = blocks.RedshiftIcon()
		dropbox, _ = blocks.DropboxIcon()
		volText, _ = blocks.VolumeText()
		volIcon, _ = blocks.VolumeIcon()
		syncthing, _ = blocks.SyncthingIcon()
		bolt = blocks.PowerIcon()
		if fiveMins == 0 || passed < 10 {
			battery, _ = blocks.Battery()
		}

		// Change by editing variables & `%s`
		status = fmt.Sprintf(" %s%s %s%s %s%s %s%s %s%s %s %s%s%s",
			tempIco, weather, homeDir, homeSpace, memIco, ramFree, volIcon, volText,
			bolt, battery, hTime, dropbox, syncthing, rShift)
		ui.Dwm(status) // change this to the UI of choice
	}
}
