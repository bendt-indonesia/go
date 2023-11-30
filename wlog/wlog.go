package wlog

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func InitLog() {
	lvl, _ := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	trace := map[int]log.Level {
		0: log.TraceLevel,
		1: log.DebugLevel,
		2: log.InfoLevel,
		3: log.WarnLevel,
		4: log.ErrorLevel,
		5: log.FatalLevel,
		6: log.PanicLevel,
	}

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(trace[lvl])
}
