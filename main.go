package main

import (
	"log"
	"math"
	"time"

	"git.swab.dev/breto.git/blocks"
	"git.swab.dev/breto.git/format"
	"git.swab.dev/breto.git/icons"
	"git.swab.dev/breto.git/stats"
	"git.swab.dev/breto.git/ui"
)

func main() {

	var (
		opt       = new(format.Options)
		info      = new(stats.Info)
		ico       = new(icons.Symbols)
		bat       = new(stats.Battery)
		cWttr     = make(chan string)
		eWttr     = make(chan error)
		cRAM      = make(chan string)
		eRAM      = make(chan error)
		cHomeDisk = make(chan string)
		eHomeDisk = make(chan error)
	)

	o := opt.ParseFlags()

	// Each Go routine has it's own timer to delay the execution of the command.
	// A Go routine will run unless it's CLI flag is set to false.
	if o.Temperature {
		go blocks.Wttr(cWttr, eWttr)
	}

	if o.Memory {
		go blocks.FreeRam(cRAM, eRAM)
	}

	if o.DiskSpace {
		go blocks.HomeDisk(cHomeDisk, eHomeDisk)
	}

	start := time.Now() // for batter time math
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		// add year & seconds with "Jan 02, 2006 15:04:05"
		info.HTime = time.Now().Format("Jan 02 15:04")

		if o.Battery {
			bat.Passed = time.Since(start).Seconds()
			bat.FiveMins = math.Floor(math.Remainder(bat.Passed, 300))
		}

		select { // updates the go routine channels as they send data
		case info.Weather = <-cWttr:
		case info.WttrErr = <-eWttr:
			log.Println(info.WttrErr.Error())
		case info.RamFree = <-cRAM:
		case info.RamErr = <-eRAM:
			log.Println(info.RamErr.Error())
		case info.HomeSpace = <-cHomeDisk:
		case info.HomeErr = <-eHomeDisk:
			log.Println(info.HomeErr.Error())
		default:
		}

		// Status bar information as defined by the CLI flags.
		// reset status on every run.
		finalStatus := o.Output("", info, ico, bat)

		// Output methods as specified by CLI flags.
		if o.Dwm {
			ui.Dwm(finalStatus)
		} else {
			ui.Default(finalStatus)
		}
	}
}
