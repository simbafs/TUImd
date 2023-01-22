package util

import (
	"fmt"
	"os"
)

var file *os.File

func Log(format string, a ...any) {
	if file == nil {
		file, _ = os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	}
	file.WriteString(fmt.Sprintf(format, a...))
}
