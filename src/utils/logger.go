package utils

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetPrefix("ghlabel:cloner -- ")
	log.SetFlags(log.Ltime | log.Lshortfile | log.Lmsgprefix)
}

func Error(format string, args ...any) {
	log.Printf(fmt.Sprintf("[ERR] %s\n", format), args...)
}

func Info(format string, args ...any) {
	log.Printf(fmt.Sprintf("[INF] %s\n", format), args...)
}

func Debug(format string, args ...any) {
	var debugFlag, _ = os.LookupEnv("DEBUG")
	if debugFlag != "" {
		log.Printf(fmt.Sprintf("[DBG] %s\n", format), args...)
	}
}
