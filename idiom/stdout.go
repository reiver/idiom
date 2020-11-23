package idiom

import (
	"os"
)

var Stdout *os.File

func init() {
	Stdout = os.Stdout
}
