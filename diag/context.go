// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package diag

import "time"

// LogLevel indicates level of a log message.
//
// Available since v0.5.2
type LogLevel int8

const (
	NoLevel LogLevel = iota
	FatalLevel
	PanicLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

// LogContext is a struct to store raw log message.
//
// Available since v0.5.2
type LogContext struct {
	Time    time.Time
	Level   LogLevel
	Message string
}
