package module

import "fmt"

type FmtLogger struct {
}

func (this FmtLogger) Tracef(format string, params ...interface{}) {
	fmt.Printf(format, params...)
}
func (this FmtLogger) Debugf(format string, params ...interface{}) {
	fmt.Printf(format, params...)
}
func (this FmtLogger) Infof(format string, params ...interface{}) {
	fmt.Printf(format, params...)
}

// Warnf formats message according to format specifier
// and writes to log with level = Warn.
func (this FmtLogger) Warnf(format string, params ...interface{}) error {
	fmt.Printf(format, params...)
	return nil
}

// Errorf formats message according to format specifier
// and writes to log with level = Error.
func (this FmtLogger) Errorf(format string, params ...interface{}) error {
	fmt.Printf(format, params...)
	return nil
}

// Criticalf formats message according to format specifier
// and writes to log with level = Critical.
func (this FmtLogger) Criticalf(format string, params ...interface{}) error {
	fmt.Printf(format, params...)
	return nil
}

// Trace formats message using the default formats for its operands
// and writes to log with level = Trace
func (this FmtLogger) Trace(v ...interface{}) {
	fmt.Println(v...)
}

// Debug formats message using the default formats for its operands
// and writes to log with level = Debug
func (this FmtLogger) Debug(v ...interface{}) {
	fmt.Println(v...)
}

// Info formats message using the default formats for its operands
// and writes to log with level = Info
func (this FmtLogger) Info(v ...interface{}) {
	fmt.Println(v...)
}

// Warn formats message using the default formats for its operands
// and writes to log with level = Warn
func (this FmtLogger) Warn(v ...interface{}) error {
	fmt.Println(v...)
	return nil
}

// Error formats message using the default formats for its operands
// and writes to log with level = Error
func (this FmtLogger) Error(v ...interface{}) error {
	fmt.Println(v...)
	return nil
}

// Critical formats message using the default formats for its operands
// and writes to log with level = Critical
func (this FmtLogger) Critical(v ...interface{}) error {
	fmt.Println(v...)
	return nil
}

func (this FmtLogger) Close() {

}

// Flush flushes all the messages in the logger.
func (this FmtLogger) Flush() {

}

// Closed returns true if the logger was previously closed.
func (this FmtLogger) Closed() bool {
	return true
}
