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

// CLI flag variables
var dwm, battery, clock, audio, memory, diskSpace, temperature, tray bool

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

type symbols struct {
	status    string
	rShift    string
	dropbox   string
	volIcon   string
	syncthing string
	bolt      string
}

type batInfo struct {
	passed   float64
	fiveMins float64
}

func init() {
	// Setup and define cli flags
	flag.BoolVar(&dwm, "dwm", false, "Used to enable output for DWM's status bar.\n Example: --dwm=true")
	flag.BoolVar(&battery, "battery", false, "Used to enable battery module.\n Example: --battery=true")
	flag.BoolVar(&clock, "dateTime", true, "Used to disable the date and time module.\n Example: --dateTime=false")
	flag.BoolVar(&audio, "volume", true, "Used to disable the volume module.\n Example: --volume=false")
	flag.BoolVar(&memory, "ram", true, "Used to disable the RAM module.\n Example: --ram=false")
	flag.BoolVar(&diskSpace, "storage", true, "Used to disable the home directory storage module.\n Example: --storage=false")
	flag.BoolVar(&temperature, "temp", true, "Used to disable the temperature module.\n Example: --temp=false")
	flag.BoolVar(&tray, "tray", true, "Used to disable the custom tray module.\n Example: --tray=false")
}

func main() {
	flag.Parse()

	stats := info{}
	ico := symbols{}
	baty := batInfo{}

	// Each Go routine has it's own timer to delay the execution of the command.
	// A Go routine will run unless it's CLI flag is set to false.
	cWttr := make(chan string)
	eWttr := make(chan error)
	if temperature {
		go blocks.Wttr(cWttr, eWttr)
	}

	cRAM := make(chan string)
	eRAM := make(chan error)
	if memory {
		go blocks.FreeRam(cRAM, eRAM)
	}

	cHomeDisk := make(chan string)
	eHomeDisk := make(chan error)
	if diskSpace {
		go blocks.HomeDisk(cHomeDisk, eHomeDisk)
	}

	start := time.Now() // for batter time math
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

		// Status bar information as defined by the CLI flags.
		status := "" // reset status on every run.
		if temperature {
			status = fmt.Sprintf("%s %s%s ", status, icons.Temp(), stats.weather)
		}
		if diskSpace {
			status = fmt.Sprintf("%s %s%s ", status, icons.Dir(), stats.homeSpace)
		}
		if memory {
			status = fmt.Sprintf("%s %s%s ", status, icons.Mem(), stats.ramFree)
		}
		if audio {
			stats.volText, _ = blocks.VolumeText()
			ico.volIcon, _ = icons.Volume()
			status = fmt.Sprintf("%s %s%s ", status, ico.volIcon, stats.volText)
		}
		if battery {
			if baty.fiveMins == 0 || baty.passed < 10 {
				stats.power, _ = blocks.Battery()
			}
			status = fmt.Sprintf("%s %s%s ", status, icons.Power(), stats.power)
		}
		if clock {
			status = fmt.Sprintf("%s %s ", status, stats.hTime)
		}
		if tray {
			ico.rShift, _ = icons.Redshift()
			ico.dropbox, _ = icons.Dropbox()
			ico.syncthing, _ = icons.Syncthing()
			status = fmt.Sprintf("%s %s%s%s", status, ico.dropbox, ico.syncthing, ico.rShift)
		}

		// Output methods as specified by CLI flags.
		if dwm {
			ui.Dwm(status)
		} else {
			ui.Default(status)
		}
	}
}
