package rotate

import (
	"log"
	"os"
)

type ActiveLog struct {
	Path    string
	File    *os.File
	Loggers []*log.Logger
}
