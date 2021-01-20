package idiom_cmdln

import (
	"flag"
	"os"
)

var (
	Name string
)

var (
	Verbose bool = false
	VeryVerbose bool = false
	VeryVeryVerbose bool = false
)

var (
	Args []string
)

func init() {
	Name = os.Args[0]

	flag.BoolVar(&Verbose,         "v",   false, "tells program to display verbose log messages while running")
	flag.BoolVar(&VeryVerbose,     "vv",  false, "tells program to display very verbose log messages while running")
	flag.BoolVar(&VeryVeryVerbose, "vvv", false, "tells program to display very very verbose log messages while running")

	flag.Parse()

	Args = flag.Args()
}
