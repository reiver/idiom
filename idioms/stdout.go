package idioms

import (
	"os"
)

var Stdout *os.File

func init() {
	Stdout = os.Stdout
}
