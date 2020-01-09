# Breto
A status bar written in Go.
Currently tested with DWM, i3wm, and tmux.

## Current Features:
### Master Branch:
#### Blocks:
- Date & Time
- Weather via wttr.in
- Total RAM not used
- Total disk space left in the binary's present partition.
- Total battery percentage (off by default)
- Icons
#### Current UIs:
- DWM
- i3wm
- tmux
- Polybar

#### CLI Flags
- `-dwm=true` to use in DWM's status bar
- `-battery=true` to enable battery block

### Icons:
#### To display icons in DWM.
1. Install Font Awesome 5
2. Set FontAwesome as the second font in `dwm/config.h`
   - eg. `static const char *fonts[] = { "Source Code Pro:size=13", "Font Awesome 5 Free:style=Regular:size=14" };`
3. Relaunch DWM

#### To display in tmux or i3wm:
1. Install Font Awesome 5
   - Keep in mind that the icons are currently very small.
   - Research in progess.

#### To display in Polybar:
1. Install Font Awesome 5:
2. Add to your Polybar config:
```
font-0 = {base font here}
font-1 = "Font Awesome 5 Free:style=Regular:pixelsize=12;1"
font-2 = "Font Awesome 5 Free:style=Solid:pixelsize=12;1"
font-3 = "Font Awesome 5 Brands:pixelsize=12;1"
```
3. If not using Font Awesome 5 search for your version with `fc-list | grep Awesome`.

#### Current Icons:
- Tempurature: 
- Disk: 
- RAM: 
- Volume:    
- Syncthing: 
- Dropbox: 
- Redshift: 
- Battery: 

Currently, the icons are much too small to be usefull in Tmux or i3wm without further configuration.
If you have an easy way to display FontAwesome icons at the same scale as the terminal text please submit a pull request.

## How To Use:
1. Open main.go in a text editor.
2. Edit the last `status` variable to contain the blocks you wish to use.
3. Build the binary.
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

## To-Do:
1. Scale icons in tmux and i3wm
   - [This Unix Stack Exchange post may help](http://unix.stackexchange.com/questions/49779/ddg#49823)

## License:
MIT
