package utils

import (
	"runtime"

	log "github.com/sirupsen/logrus"
)

func PrintFunctionName() {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	log.WithFields(log.Fields{"function": f.Name()}).Debug()
}
