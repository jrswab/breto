package stats

type Info struct {
	HTime     string
	Weather   string
	RamFree   string
	HomeSpace string
	VolText   string
	Power     string

	WttrErr error
	RamErr  error
	HomeErr error
}

type Battery struct {
	Passed   float64
	FiveMins float64
}
