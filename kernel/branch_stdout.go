package idiom_kernel

import (
	"os"
)

func (receiver Branch) Stdout() *os.File {
	return os.Stdout
}
