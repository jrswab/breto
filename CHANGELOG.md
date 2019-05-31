# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.2.5] - 
### Added
- Go test for wttr package
- Another Go channel for wttr errors

### Changed
- location of package directories to top level

## [1.2.4] - May 29, 2019
### Changed
- Removed the counters from the bottom of main.go and cleaned up comments
- Now each lib has a ticker and sends data to main ever 'n' seconds

## [1.2.3] - May 28, 2019
### Changed
- Passing varibles and memory adderss for go routines

## [1.2.2] - May 27, 2019
### Change
- the passing of weather data to passing the memory adderess

## [1.2.1] - May 23, 2019
### Changed
- status.go to main.go
- Functions to custom packages for a cleaner removal if a module is unwanted
- Space formatting at the start and end of the status info

### Removed
- functions from main.go

## [1.2.0] - May 22, 2019
### Added
- Hourly weather update function

### Changed
- RAM to a function call named getRam()
- RAM to update every three seconds

### Removed
- The binary file

## [1.1.0] - May 21, 2019
### Added
- Total GiB of Ram remaining on the system.

## [1.0.0] - May 20, 2019
### Added
- Time updated each second
