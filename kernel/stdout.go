package idiom_kernel

import (
	"os"
)

var Stdout *os.File

func init() {
	Stdout = os.Stdout
}
