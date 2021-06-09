package main

import (
	"log"
	"math"
	"os"
	"sync"
	"time"
	"fmt"

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
		mutex     = &sync.Mutex{}
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
			writeToLog(info.WttrErr.Error(), mutex)
		case info.RamFree = <-cRAM:
		case info.RamErr = <-eRAM:
			writeToLog(info.RamErr.Error(), mutex)
		case info.HomeSpace = <-cHomeDisk:
		case info.HomeErr = <-eHomeDisk:
			writeToLog(info.HomeErr.Error(), mutex)
		default:
		}

		// Status bar information as defined by the CLI flags.
		// reset status on every run.
		status := o.Output("", info, ico, bat)

		// Output methods as specified by CLI flags.
		if o.Dwm {
			ui.Dwm(status)
		} else {
			ui.Default(status)
		}
	}
}

func writeToLog(errMsg string, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()

	cache := os.Getenv("XDG_CACHE_HOME")
	if cache == "" {
		cache = "."
	}

	path := fmt.Sprintf("%s/breto.log", cache)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "prefix", log.LstdFlags)
	logger.Println(errMsg)
}
