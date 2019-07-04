# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- In code instructions to enable battery status.

### Changed
- Commented out all battery status code in main.go
- wttr custom error message to used fmt.Errorf instead for errors.New

### Removed
- The errors package from wttr.go

## [0.8.1] - July 3, 2019
## Changed
- Wttr from checking for server error to checking for desired output.

## [0.8.0] - July 2, 2019
### Added
- Battery percentage for laptops
- Batter power icon
### Changed
- Shell command for disk space from "home" to the partition the binary is run within.
- Month, day, year, time to month, day, time

## [0.7.2] - June 7, 2019
### Added
- contributing.md
### Changed
- To a custom error message for the error channel.
### Fixed
- strings.Contains() due to missing the last `s` before the `()`
- converted bodyData to a string

## [0.7.1] - June 6, 2019
### Changed
- The check statement in wttr.go that looks for the server error webpage.

## [0.7.0] - June 5, 2019
### Added
- Volume percent number (blocks/volume.go)
- Remaining home directory storage (blocks/disk.go)
- More icons

### Changed
- Block output formatting to allow for a less crammed Printf in main.go
- Made blocks more modular for custom formatting

## [0.6.0] - June 3, 2019
## Added
- FontAwesome icons for DWM
- Special error case for wttr.in for when it returns a webpage explaining a server error and not an http get error.
- Volume percent number as a block

## [0.5.0] - June 2, 2019
### Added
- UI package to allow for easy switching between user interfaces.

### Changed
- Moved weather and ram into a 'blocks' package.

## [0.4.0] - May 31, 2019
### Added
- Go test for wttr and ram packages
- Extra Go channels for errors
- Instructions for using in Tmux

### Changed
- Location of package directories to top level
- Project name from dwm-status to Go Status

## [0.3.4] - May 29, 2019
### Changed
- Removed the counters from the bottom of main.go and cleaned up comments
- Now each lib has a ticker and sends data to main ever 'n' seconds

## [0.3.3] - May 28, 2019
### Changed
- Passing varibles and memory adderss for go routines

## [0.3.2] - May 27, 2019
### Change
- The passing of weather data to passing the memory adderess

## [0.3.1] - May 23, 2019
### Changed
- Status.go to main.go
- Functions to custom packages for a cleaner removal if a module is unwanted
- Space formatting at the start and end of the status info

### Removed
- Functions from main.go

## [0.3.0] - May 22, 2019
### Added
- Hourly weather update function

### Changed
- RAM to a function call named getRam()
- RAM to update every three seconds

### Removed
- The binary file

## [0.2.0] - May 21, 2019
### Added
- Total GiB of Ram remaining on the system.

## [0.1.0] - May 20, 2019
### Added
- Time updated each second
