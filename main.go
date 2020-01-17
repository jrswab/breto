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
var dwmIsSet, batteryIsSet, clock, audio, memory, diskSpace, temperature, tray, emoji bool

type batInfo struct {
	passed   float64
	fiveMins float64
}

func init() {
	// Setup and define cli flags
	flag.BoolVar(&dwmIsSet, "dwm", false, "Used to enable output for DWM's status bar.\n Example: --dwm=true")
	flag.BoolVar(&batteryIsSet, "battery", false, "Used to enable battery module.\n Example: --battery=true")
	flag.BoolVar(&clock, "dateTime", true, "Used to disable the date and time module.\n Example: --dateTime=false")
	flag.BoolVar(&audio, "volume", true, "Used to disable the volume module.\n Example: --volume=false")
	flag.BoolVar(&memory, "ram", true, "Used to disable the RAM module.\n Example: --ram=false")
	flag.BoolVar(&diskSpace, "storage", true, "Used to disable the home directory storage module.\n Example: --storage=false")
	flag.BoolVar(&temperature, "temp", true, "Used to disable the temperature module.\n Example: --temp=false")
	flag.BoolVar(&tray, "tray", true, "Used to disable the custom tray module.\n Example: --tray=false")
	flag.BoolVar(&emoji, "emoji", false, "Used to enable emoji icons instead of Awosome Font.\n Example: --emoji=true")
}

func main() {
	flag.Parse()

	stats := blocks.Info{}
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
		hTime := time.Now().Format("Jan 02 15:04")

		if batteryIsSet {
			baty.passed = time.Since(start).Seconds()
			baty.fiveMins = math.Floor(math.Remainder(baty.passed, 300))
		}

		select { // updates the go routine channels as they send data
		case stats.Weather = <-cWttr:
		case stats.WttrErr = <-eWttr:
			log.Println(stats.WttrErr.Error())
			stats.Weather = "N/A"
		case stats.RamFree = <-cRAM:
		case stats.RamErr = <-eRAM:
			log.Println(stats.RamErr.Error())
		case stats.HomeSpace = <-cHomeDisk:
		case stats.HomeErr = <-eHomeDisk:
			log.Println(stats.HomeErr.Error())
		default:
		}

		// Status bar information as defined by the CLI flags.
		status := "" // reset status on every run.
		if temperature {
			status = fmt.Sprintf("%s %s%s ", status, icons.Temp(emoji), stats.Weather)
		}
		if diskSpace {
			status = fmt.Sprintf("%s %s%s ", status, icons.Dir(emoji), stats.HomeSpace)
		}
		if memory {
			status = fmt.Sprintf("%s %s%s ", status, icons.Mem(emoji), stats.RamFree)
		}
		if audio {
			volText, _ := blocks.VolumeText()
			volIcon, _ := icons.Volume(emoji)
			status = fmt.Sprintf("%s %s%s ", status, volIcon, volText)
		}
		if batteryIsSet {
			if baty.fiveMins == 0 || baty.passed < 10 {
				stats.Power, _ = blocks.Battery()
			}
			status = fmt.Sprintf("%s %s%s ", status, icons.Power(emoji), stats.Power)
		}
		if clock {
			status = fmt.Sprintf("%s %s ", status, hTime)
		}
		if tray {
			redShift, _ := icons.Redshift(emoji)
			dropbox, _ := icons.Dropbox(emoji)
			syncthing, _ := icons.Syncthing(emoji)
			status = fmt.Sprintf("%s %s%s%s", status, dropbox, syncthing, redShift)
		}

		// Output methods as specified by CLI flags.
		if dwmIsSet {
			ui.Dwm(status)
		} else {
			ui.Default(status)
		}
	}
}
