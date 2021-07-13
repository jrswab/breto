package format

import (
	"flag"
	"fmt"
	"log"

	"git.swab.dev/breto.git/blocks"
	"git.swab.dev/breto.git/icons"
	"git.swab.dev/breto.git/stats"
)

// CLI flag variables
type Options struct {
	Dwm, Battery, Clock, Audio, CPU, Memory, DiskSpace, Temperature, Tray, Emoji, DynamicVol bool
}

func (o *Options) ParseFlags() *Options {
	// Setup and define cli flags
	flag.BoolVar(&o.Clock, "dateTime", true, "Used to disable the date and time module.\n Example: --dateTime=false")
	flag.BoolVar(&o.Audio, "volume", true, "Used to disable the volume-text module.\n Example: --volume=false")
	flag.BoolVar(&o.Memory, "ram", true, "Used to disable the RAM module.\n Example: --ram=false")
	flag.BoolVar(&o.CPU, "cpu", true, "Used to disable the CPU module.\n Example: --cpu=false")

	// Flags disabled by default:
	flag.BoolVar(&o.DiskSpace, "storage", false, "Used to enable the home directory storage module.\n Example: --storage=true")
	flag.BoolVar(&o.Temperature, "temp", false, "Used to enable the temperature module.\n Example: --temp=true")
	flag.BoolVar(&o.Tray, "tray", false, "Used to enable the custom tray module.\n Example: --tray=true")
	flag.BoolVar(&o.Emoji, "emoji", false, "Used to enable Openmoji icons instead of Awesome Font.\n Example: --emoji=true")
	flag.BoolVar(&o.Battery, "battery", false, "Used to enable battery module.\n Example: --battery=true")
	flag.BoolVar(&o.DynamicVol, "volume-icon-only", false, "Used to enable the dynamic volume icon module. (This disables the volume text.)\n Example: --volume-icon-only=false")

	// UI flags:
	flag.BoolVar(&o.Dwm, "dwm", false, "Used to enable output for DWM's status bar.\n Example: --dwm=true")

	flag.Parse()

	return o
}

func (o *Options) Output(status string, info *stats.Info, ico *icons.Symbols, bat *stats.Battery) string {
	if o.Temperature {
		status = fmt.Sprintf("%s %s%s ", status, icons.Temp(o.Emoji), info.Weather)
	}
	if o.CPU {
		status = fmt.Sprintf("%s %s%s %s ", status, icons.CPU(o.Emoji), info.CPUMHz, info.CPUTemp)
	}
	if o.DiskSpace {
		status = fmt.Sprintf("%s %s%s ", status, icons.Dir(o.Emoji), info.HomeSpace)
	}
	if o.Memory {
		status = fmt.Sprintf("%s %s%s ", status, icons.Mem(o.Emoji), info.RamFree)
	}
	if o.Audio {
		info.VolText, _ = blocks.Volume(o.Emoji)
		volumeIcon := icons.VolumeSingleIcon(o.Emoji)
		var err error
		if o.DynamicVol {
			volumeIcon, err = icons.VolumeDynamic(o.Emoji)
			if err != nil {
				log.Printf("dynamic volume icon encountered an error: %s\n", err)
			}
		}
		status = fmt.Sprintf("%s %s%s ", status, volumeIcon, info.VolText)
	}
	if o.Battery {
		if bat.Minute == 0 || bat.Passed < 10 {
			info.Power, _ = blocks.Battery()
		}
		status = fmt.Sprintf("%s %s%s ", status, icons.Power(o.Emoji), info.Power)
	}
	if o.Clock {
		status = fmt.Sprintf("%s %s ", status, info.HTime)
	}
	if o.Tray {
		ico.RShift, _ = icons.Redshift(o.Emoji)
		ico.Dropbox, _ = icons.Dropbox(o.Emoji)
		ico.Syncthing, _ = icons.Syncthing(o.Emoji)
		status = fmt.Sprintf("%s %s%s%s", status, ico.Dropbox, ico.Syncthing, ico.RShift)
	}
	return status
}
