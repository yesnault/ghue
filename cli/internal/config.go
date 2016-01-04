package internal

import "os"

var (
	// Verbose conditions the quantity of output of api requests
	Verbose bool

	// Home fetches the user home directory
	Home = os.Getenv("HOME")

	// Format to use for output. One of 'json', 'yaml', 'pretty'
	Format string
)
