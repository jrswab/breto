[![Go Report Card](https://goreportcard.com/badge/github.com/jrswab/breto)](https://goreportcard.com/report/github.com/jrswab/breto)
# Breto
A status bar written in Go.
Currently tested with DWM, i3wm, tmux, and Polybar.

## Dependencies:
- pulsemixer
- lm_sensors
- bc

## Current Features:
### Master Branch:
#### Blocks:
- CPU MHz & tempurature
- Date & Time
- Weather via wttr.in
- Total RAM not used
- Total disk space left in the binary's present partition.
- Audio volume percentage (pamixer needs installed)
- Total battery percentage (off by default)
- Icons

#### Current UIs Tested:
- DWM
- i3wm
- tmux
- Polybar

#### CLI Flags
- `-battery=true` to enable battery block
- `-dateTime=false` to remove the date/time block
- `-dwm=true` to use in DWM's status bar
- `-emoji=true` to enable emoji icons instead of Font Awesome
- `-ram=false` to remove RAM remaining block
- `-storage=false` to remove the home directory storage remaining block
- `-temp=false` to remove the weather block
- `-tray=false` to remove the "system" tray block
- `-volume=false` to remove volume percentage block

### Icons:
#### To display icons in DWM.
1. Install Font Awesome 5 (or an emoji font)
2. Set FontAwesome as the second font in `dwm/config.h`
   - eg. `static const char *fonts[] = { "Source Code Pro:size=13", "Font Awesome 5 Free:style=Regular:size=14" };`
   - if using an emoji font, replace the second font name with the emoji set
3. Relaunch DWM

#### To display in tmux or i3wm:
1. Install Font Awesome 5 (or an emoji font)
   - Keep in mind that the icons are currently very small.
   - Research in progess.

#### To display in Polybar:
1. Install Font Awesome 5 (or an emoji font)
2. Add to your Polybar config:
   - If not using Font Awesome search for your emoji set with `fc-list` and replace "Font Awesome 5".
```
font-0 = {base font here}
font-1 = "Font Awesome 5 Free:style=Regular:pixelsize=12;1"
font-2 = "Font Awesome 5 Free:style=Solid:pixelsize=12;1"
font-3 = "Font Awesome 5 Brands:pixelsize=12;1"
```

## How To Use:
1. Clone repository.
2. `cd` into the directory.
3. Build the binary.
   - `go build`
4. Edit your config file to use the new binary.
   - If using DWM execute with the `--dwm=true` flag.

### Adding to DWM:
- If you already have a startup script for DWM just add a new line with the path to this binary.

### Addng to Tmux:
- `set -g status-right "#($HOME/PATHTO/tmux-status)"`
 - If you have colors in this setting add the path at the end of the string
 - Be sure to use the correct path and name of the file you built with GO.
 - Running `mv breto ~/tmux-status` will allow you to use `"($HOME/tmux-status)"` in your config.
- `set -g status-right-length 53`
 - If you are not using all the custom packages this number can be lower
 - This also will vary based on screen size. 53 is the minimum that worked for all current blocks on my montior.
 - If you notice the status getting cut off just increase the number and reload tmux.

### Adding to i3wm:
- In the i3wm confige file, change `status_command ...` to `status_command PATH/TO/breto`

### Adding to Polybar:
- Add the following module to your Polybar configuration file:
```
[module/breto]
type = custom/script
exec = /path/to/breto
tail = true
```

## Wttr.in Options:
1. Add your area to the weather function
   - Area Code: 'wttr.in/~00000?format=2' 
   - City: 'wttr.in/~Paris?format=2'
	 - More information can be found at [wttr.in/:help](https://wttr.in/:help)
2. Add tweaks to `blocks/wttr.go`

## TODO:
- Add CLI flag for user input for weather location.

## License:
MIT
