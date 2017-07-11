package glog

// default logger
var defaultLogger *Glog

func init() {
	defaultLogger = New(false)
	defaultLogger.SetLevel(LogLevelError)
	defaultLogger.SetInterval(10)
}

// SetInterval default logger set the interval of log level value when reload environmental opt
func SetInterval(val int) {
	defaultLogger.SetInterval(val)
}

// SetLevel default logger set level
func SetLevel(logLevel int) error {
	return defaultLogger.SetLevel(logLevel)
}

// SetReload default logger set the reload boolean label
func SetReload(isReload bool) {
	defaultLogger.SetReload(isReload)
}

// Debugf default logger use the debug log level output
func Debugf(format string, v ...interface{}) string {
	return defaultLogger.Debugf(format, v...)
}

// Infof default logger use the debug log level output
func Infof(format string, v ...interface{}) string {
	return defaultLogger.Infof(format, v...)
}

// Warnf default logger use the warn log level output
func Warnf(format string, v ...interface{}) string {
	return defaultLogger.Warnf(format, v...)
}

// Errorf default logger use the error log level output
func Errorf(format string, v ...interface{}) string {
	return defaultLogger.Errorf(format, v...)
}
