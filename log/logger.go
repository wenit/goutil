package log

import (
	"log"
	"os"
	"fmt"
	"strings"
)

type Logger struct {
	log *log.Logger
}

func New() *Logger {
	logger := &Logger{}
	logger.log = log.New(os.Stdout, "", log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	return logger
}

func (logger *Logger)Debug(v ...interface{}) {
	outPut(logger, 1, v...)
}

func (logger *Logger)Info(v ...interface{}) {
	outPut(logger, 2, v...)
}

func (logger *Logger)Warn(v ...interface{}) {
	outPut(logger, 3, v...)
}

func (logger *Logger)Error(v ...interface{}) {
	outPut(logger, 4, v...)
}

func outPut(logger *Logger, logType int, v ...interface{}) {
	tag := "[INFO ] "
	switch logType {
	case 1:
		tag = "[DEBUG] "
	case 2:
		tag = "[INFO ] "
	case 3:
		tag = "[WARN ] "
	case 4:
		tag = "[ERROR] "
	default:
		tag = "[INFO ] "
	}

	s := ""
	index := -1

	length := len(v)

	var first interface{}
	format := ""
	if length > 0 {
		first = v[0]

	}

	switch first.(type) {
	case string:
		format = first.(string)
		index = strings.Index(format, "%")
	default:

	}

	if index >= 0 {
		subArgs := v[1:]
		s = fmt.Sprintf(tag + format, subArgs...)
	} else {
		s = fmt.Sprint(tag + fmt.Sprintln(v...))
	}

	logger.log.Output(3, s)
}
