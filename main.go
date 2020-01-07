package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jrswab/breto/blocks"
	"github.com/jrswab/breto/icons"
	"github.com/jrswab/breto/ui"
)

// To Add Battery Status:
// Uncomment all lines that follow the message:
// "Uncomment for battery status"
// Also, add the correct string formatting to
// the status assignment at the end of the file and add "math" to the imports

// info holds dynamic information
type info struct {
	hTime     string
	weather   string
	ramFree   string
	homeSpace string
	volText   string
	wttrErr   error
	ramErr    error
	homeErr   error
}

// icons holds the text of dynamic icons
type symbols struct {
	status    string
	rShift    string
	dropbox   string
	volIcon   string
	syncthing string
}

// batInfo holds information for battery capacity
type batInfo struct {
	passed   float64
	fiveMins float64
}

func main() {
	var status string
	stats := info{}
	ico := symbols{}
	// Uncomment for battery status
	/*
		baty := batInfo{}
		start := time.Now()
	*/

	// These are static icons and only need defined at the start
	homeDir := icons.Dir()
	memIco := icons.Mem()
	tempIco := icons.Temp()

	// Each Go routine has it's own timer to delay the execution of the command.
	cWttr := make(chan string) // start weather data routine
	eWttr := make(chan error)
	go blocks.Wttr(cWttr, eWttr)

	cRAM := make(chan string) // start free ram data routine
	eRAM := make(chan error)
	go blocks.FreeRam(cRAM, eRAM)

	cHomeDisk := make(chan string) // start free home dir space routine
	eHomeDisk := make(chan error)
	go blocks.HomeDisk(cHomeDisk, eHomeDisk)

	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		// add year & seconds with "Jan 02, 2006 15:04:05"
		stats.hTime = time.Now().Format("Jan 02 15:04")

		// Uncomment for battery status
		/*
			baty.passed = time.Since(start).Seconds()
			baty.fiveMins = math.Floor(math.Remainder(passed, 300))
		*/

		select { // update the go routine channels as they send data
		case stats.weather = <-cWttr:
		case stats.wttrErr = <-eWttr:
			log.Println(stats.wttrErr.Error())
		case stats.ramFree = <-cRAM:
		case stats.ramErr = <-eRAM:
			log.Println(stats.ramErr.Error())
		case stats.homeSpace = <-cHomeDisk:
		case stats.homeErr = <-eHomeDisk:
			log.Println(stats.homeErr.Error())
		default:
		}

		// Assign Icons & non Go Routine blocks every round
		ico.rShift, _ = icons.Redshift()
		ico.dropbox, _ = icons.Dropbox()
		stats.volText, _ = blocks.VolumeText()
		ico.volIcon, _ = icons.Volume()
		ico.syncthing, _ = icons.Syncthing()

		// Uncomment for battery status
		/*
			bolt = icons.Power()
			if baty.fiveMins == 0 || passed < 10 {
				battery, _ = blocks.Battery()
			}
		*/

		// Change by editing variables & `%s`
		status = fmt.Sprintf(" %s%s %s%s %s%s %s%s %s %s%s%s",
			tempIco, stats.weather, homeDir, stats.homeSpace,
			memIco, stats.ramFree, ico.volIcon, stats.volText,
			stats.hTime, ico.dropbox, ico.syncthing, ico.rShift)
		ui.Polybar(status) // change this to the UI of choice
	}
}
