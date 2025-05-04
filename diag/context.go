// Copyright (C) 2025 T-Force I/O
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
