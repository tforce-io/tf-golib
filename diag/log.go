// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package diag

import "log"

// Logger defines required function for all log levels used for application logging.
//
// Available since v0.5.0
type Logger interface {
	// Print a message with Error level.
	Error(err error, v ...interface{})
	// Print a message with Error level with format.
	Errorf(err error, format string, v ...interface{})

	// Print a message with Warn level.
	Warn(v ...interface{})
	// Print a message with Warn level with format.
	Warnf(format string, v ...interface{})

	// Print a message with Info level.
	Info(v ...interface{})
	// Print a message with Info level with format.
	Infof(format string, v ...interface{})

	// Print a message with Debug level.
	Debug(v ...interface{})
	// Print a message with Debug level with format.
	Debugf(format string, v ...interface{})

	// Print a message with Trace level.
	Trace(v ...interface{})
	// Print a message with Trace level with format.
	Tracef(format string, v ...interface{})
}

// DefaultLogger implement Logger interface that prints log message to stdout
// using global logger instance of Go.
//
// Available since v0.5.0
type DefaultLogger struct {
}

// Print a message with Error level.
//
// Available since v0.5.0
func (l DefaultLogger) Error(err error, v ...interface{}) {
	v2 := append([]interface{}{"ERROR ", err, " "}, v...)
	log.Print(v2...)
}

// Print a message with Error level with format.
//
// Available since v0.5.0
func (l DefaultLogger) Errorf(err error, format string, v ...interface{}) {
	v2 := append([]interface{}{"ERROR ", err}, v...)
	log.Printf("%s%v "+format, v2...)
}

// Print a message with Warn level.
//
// Available since v0.5.0
func (l DefaultLogger) Warn(v ...interface{}) {
	v2 := append([]interface{}{"WARN "}, v...)
	log.Print(v2...)
}

// Print a message with Warn level with format.
//
// Available since v0.5.0
func (l DefaultLogger) Warnf(format string, v ...interface{}) {
	v2 := append([]interface{}{"WARN "}, v...)
	log.Printf("%s"+format, v2...)
}

// Print a message with Info level.
//
// Available since v0.5.0
func (l DefaultLogger) Info(v ...interface{}) {
	v2 := append([]interface{}{"INFO "}, v...)
	log.Print(v2...)
}

// Print a message with Info level with format.
//
// Available since v0.5.0
func (l DefaultLogger) Infof(format string, v ...interface{}) {
	v2 := append([]interface{}{"INFO "}, v...)
	log.Printf("%s"+format, v2...)
}

// Print a message with Debug level.
//
// Available since v0.5.0
func (l DefaultLogger) Debug(v ...interface{}) {
	v2 := append([]interface{}{"DEBUG "}, v...)
	log.Print(v2...)
}

// Print a message with Debug level with format.
//
// Available since v0.5.0
func (l DefaultLogger) Debugf(format string, v ...interface{}) {
	v2 := append([]interface{}{"DEBUG "}, v...)
	log.Printf("%s"+format, v2...)
}

// Print a message with Trace level.
//
// Available since v0.5.0
func (l DefaultLogger) Trace(v ...interface{}) {
	v2 := append([]interface{}{"TRACE "}, v...)
	log.Print(v2...)
}

// Print a message with Trace level with format.
//
// Available since v0.5.0
func (l DefaultLogger) Tracef(format string, v ...interface{}) {
	v2 := append([]interface{}{"TRACE "}, v...)
	log.Printf("%s"+format, v2...)
}
