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
