package systype

import (
	"runtime"
	"time"
)

// return the system type

func GetSysType() string {
	return runtime.GOOS
}

func GetTime() string {
	return time.Now().Format("2006-01-02-15-04-05")
}