package main

import (
	"flag"
	"fmt"
	"log"
	"math"
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
	power     string
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
	bolt      string
}

// batInfo holds information for battery capacity
type batInfo struct {
	passed   float64
	fiveMins float64
}

func main() {
	var dwm bool
	flag.BoolVar(&dwm, "dwm", false, "Used to enable output for DWM's status bar\n Example: -dwm=true")
	var battery bool
	flag.BoolVar(&battery, "battery", false, "Used to enable battery module.\n Example: -battery=true")
	flag.Parse()

	var status string
	stats := info{}
	ico := symbols{}
	baty := batInfo{}
	start := time.Now()

	// These are static icons and only need defined at the start
	homeDir := icons.Dir()
	memIco := icons.Mem()
	tempIco := icons.Temp()
	ico.bolt = icons.Power()

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

		if battery {
			baty.passed = time.Since(start).Seconds()
			baty.fiveMins = math.Floor(math.Remainder(baty.passed, 300))
		}

		select { // updates the go routine channels as they send data
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

		if battery {
			if baty.fiveMins == 0 || baty.passed < 10 {
				stats.power, _ = blocks.Battery()
			}
		}

		// Change by editing variables & `%s`
		status = fmt.Sprintf("%s%s", tempIco, stats.weather)
		status = fmt.Sprintf(" %s %s%s", status, homeDir, stats.homeSpace)
		status = fmt.Sprintf(" %s %s%s", status, memIco, stats.ramFree)
		status = fmt.Sprintf(" %s %s%s", status, ico.volIcon, stats.volText)
		if battery {
			status = fmt.Sprintf(" %s %s%s", status, ico.bolt, stats.power)
		}
		status = fmt.Sprintf(" %s %s", status, stats.hTime)
		status = fmt.Sprintf(" %s %s%s%s", status, ico.dropbox, ico.syncthing, ico.rShift)

		if dwm {
			ui.Dwm(status)
		} else {
			ui.Default(status)
		}
	}
}
