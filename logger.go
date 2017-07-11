package glog

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

// LogLevelDebug, LogLevelInfo, LogLevelWarn, LogLevelError define the logger level
const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

// Glog is a standard error output logger, allow user it to print any log info that user want
type Glog struct {
	*log.Logger
	level    int
	isReload bool
	interval int // reload log level interval(unit: second), default/minimal: 10s
}

// New create new glog
func New(isReload bool) *Glog {
	logger := log.New(os.Stdout, "[LOG]", log.LstdFlags|log.Lmicroseconds)
	gl := &Glog{
		Logger:   logger,
		level:    LogLevelError,
		isReload: isReload,
		interval: 10,
	}
	go gl.ReloadLogLevel()
	return gl
}

// SetInterval set the interval of log level value when reload environmental opt
func (g *Glog) SetInterval(val int) {
	// check the virtual value of interval
	if val < 10 {
		val = 10
	}
	g.interval = val
}

// SetLevel set logger output level
func (g *Glog) SetLevel(logLevel int) error {
	if logLevel < 0 || logLevel > LogLevelError {
		return errors.New("Invalid Log level")
	}
	g.level = logLevel
	if g.level == LogLevelDebug {
		g.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	}
	return nil
}

// SetReload set the reload boolean label
func (g *Glog) SetReload(isReload bool) {
	g.isReload = isReload
}

// ReloadLogLevel reload log level dynamically from system environmental variable
func (g *Glog) ReloadLogLevel() {
	logLevelStr := map[int]string{
		0: "DEBUG",
		1: "INFO",
		2: "WARN",
		3: "ERROR",
	}
	t := time.Tick(time.Second * time.Duration(g.interval))
	for range t {
		if g.isReload {
			logLevel := os.Getenv("LOG_LEVEL")
			switch logLevel {
			case "DEBUG":
				g.level = LogLevelDebug
			case "INFO":
				g.level = LogLevelInfo
			case "WARN":
				g.level = LogLevelWarn
			case "ERROR":
				g.level = LogLevelError
			default:
				// if LOG_LEVEL does not exist, DO NOT do any change
			}
		}
		g.Printf("Current log level : %s", logLevelStr[g.level])
	}
}

// Debugf use the debug log level output
func (g *Glog) Debugf(format string, v ...interface{}) string {
	if g.level <= LogLevelDebug {
		g.SetPrefix("[LOG][DEBUG]: ")
		g.Printf(format, v...)
		return fmt.Sprintf("[LOG][DEBUG]: "+format, v...)
	}
	return ""
}

// Infof use the debug log level output
func (g *Glog) Infof(format string, v ...interface{}) string {
	if g.level <= LogLevelInfo {
		g.SetPrefix("[LOG][INFO]: ")
		g.Printf(format, v...)
		return fmt.Sprintf("[LOG][INFO]: "+format, v...)
	}
	return ""
}

// Warnf use the warn log level output
func (g *Glog) Warnf(format string, v ...interface{}) string {
	if g.level <= LogLevelWarn {
		g.SetPrefix("[LOG][WARN]: ")
		g.Printf(format, v...)
		return fmt.Sprintf("[LOG][WARN]: "+format, v...)
	}
	return ""
}

// Errorf use the error log level output
func (g *Glog) Errorf(format string, v ...interface{}) string {
	if g.level <= LogLevelError {
		g.SetPrefix("[LOG][ERROR]: ")
		g.Printf(format, v...)
		return fmt.Sprintf("[LOG][ERROR]: "+format, v...)
	}
	return ""
}
