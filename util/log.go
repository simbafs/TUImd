package util

import (
	"os"
)

var file *os.File

func LogLn(log string) {
	if file == nil {
		file, _ = os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	}
	file.WriteString(log + "\n")
}
