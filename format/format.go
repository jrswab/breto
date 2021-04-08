package format

import (
	"flag"
	"fmt"

	"git.swab.dev/breto.git/blocks"
	"git.swab.dev/breto.git/icons"
	"git.swab.dev/breto.git/stats"
)

// CLI flag variables
type Options struct {
	Dwm, Battery, Clock, Audio, Memory, DiskSpace, Temperature, Tray, Emoji bool
}

func (o *Options) ParseFlags() *Options {
	// Setup and define cli flags
	flag.BoolVar(&o.Dwm, "dwm", false, "Used to enable output for DWM's status bar.\n Example: --dwm=true")
	flag.BoolVar(&o.Battery, "battery", false, "Used to enable battery module.\n Example: --battery=true")
	flag.BoolVar(&o.Clock, "dateTime", true, "Used to disable the date and time module.\n Example: --dateTime=false")
	flag.BoolVar(&o.Audio, "volume", true, "Used to disable the volume module.\n Example: --volume=false")
	flag.BoolVar(&o.Memory, "ram", true, "Used to disable the RAM module.\n Example: --ram=false")
	flag.BoolVar(&o.DiskSpace, "storage", true, "Used to disable the home directory storage module.\n Example: --storage=false")
	flag.BoolVar(&o.Temperature, "temp", true, "Used to disable the temperature module.\n Example: --temp=false")
	flag.BoolVar(&o.Tray, "tray", true, "Used to disable the custom tray module.\n Example: --tray=false")
	flag.BoolVar(&o.Emoji, "emoji", false, "Used to enable Openmoji icons instead of Awesome Font.\n Example: --emoji=true")

	flag.Parse()

	return o
}

func (o *Options) Output(status string, info *stats.Info, ico *icons.Symbols, bat *stats.Battery) string {
	if o.Temperature {
		status = fmt.Sprintf("%s %s%s ", status, icons.Temp(o.Emoji), info.Weather)
	}
	if o.DiskSpace {
		status = fmt.Sprintf("%s %s%s ", status, icons.Dir(o.Emoji), info.HomeSpace)
	}
	if o.Memory {
		status = fmt.Sprintf("%s %s%s ", status, icons.Mem(o.Emoji), info.RamFree)
	}
	if o.Audio {
		info.VolText, _ = blocks.VolumeText()
		ico.VolIcon, _ = icons.Volume(o.Emoji)
		status = fmt.Sprintf("%s %s%s ", status, ico.VolIcon, info.VolText)
	}
	if o.Battery {
		if bat.FiveMins == 0 || bat.Passed < 10 {
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
