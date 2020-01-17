package blocks

// Info holds the variables for the block package.
type Info struct {
	Weather   string
	RamFree   string
	HomeSpace string
	VolText   string
	Power     string
	WttrErr   error
	RamErr    error
	HomeErr   error
}
