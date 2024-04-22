package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

const LOGFILE = "gomediamanager.log"

func Initiate() {
	fileLog := flag.Bool("filelog", true, "Log to file"+LOGFILE)
	errLog := flag.Bool("errlog", true, "Log errors to the error streams")

	logggers := []io.Writer{}

	if *fileLog {
		file, err := os.OpenFile(LOGFILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			Error(fmt.Errorf("%v %v", err.Error(), "Failed to open log file"), 0)
		}
		logggers = append(logggers, file)
	}
	if *errLog {
		logggers = append(logggers, os.Stderr)
	}
	log.SetOutput(io.MultiWriter(logggers...))
}

func Error(err error, stackRewind int) {
	_, file, line, ok := runtime.Caller(stackRewind + 1)
	if ok {
		log.Printf("%s:%d %v\n", file, line, err)
	}
}

func Fatal(err error, stackRewind int) {
	_, file, line, ok := runtime.Caller(stackRewind + 1)
	if ok {
		log.Panic("%s:%d %v\n", file, line, err)
	}
}
