package util

import "os"

var file, _ = os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
var Log = file.WriteString
