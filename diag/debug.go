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

import (
	"container/ring"
	"fmt"
	"sync"
	"time"
)

// DebugLogger implement Logger interface that store log contents in a circular list.
//
// Available since v0.5.2
type DebugLogger struct {
	i *DebugLoggerInternal
}

// DebugLoggerInternal stores internal data of a DebugLogger.
//
// Available since v0.5.2
type DebugLoggerInternal struct {
	Cache *ring.Ring
	LogMu *sync.Mutex
}

// Return new DebugLogger with maximum capacity for cache.
//
// Available since v0.5.2
func NewDebugLogger(capacity int) *DebugLogger {
	return &DebugLogger{
		&DebugLoggerInternal{
			Cache: ring.New(capacity),
			LogMu: &sync.Mutex{},
		},
	}
}

// Print a message with Error level.
//
// Available since v0.5.2
func (l DebugLogger) Error(err error, v ...interface{}) {
	v2 := append([]interface{}{"ERROR ", err, " "}, v...)
	l.write(ErrorLevel, v2...)

}

// Print a message with Error level with format.
//
// Available since v0.5.2
func (l DebugLogger) Errorf(err error, format string, v ...interface{}) {
	v2 := append([]interface{}{"ERROR ", err}, v...)
	l.writef(ErrorLevel, "%s%v "+format, v2...)
}

// Print a message with Warn level.
//
// Available since v0.5.2
func (l DebugLogger) Warn(v ...interface{}) {
	v2 := append([]interface{}{"WARN "}, v...)
	l.write(WarnLevel, v2...)
}

// Print a message with Warn level with format.
//
// Available since v0.5.2
func (l DebugLogger) Warnf(format string, v ...interface{}) {
	v2 := append([]interface{}{"WARN "}, v...)
	l.writef(WarnLevel, "%s"+format, v2...)
}

// Print a message with Info level.
//
// Available since v0.5.2
func (l DebugLogger) Info(v ...interface{}) {
	v2 := append([]interface{}{"INFO "}, v...)
	l.write(InfoLevel, v2...)
}

// Print a message with Info level with format.
//
// Available since v0.5.2
func (l DebugLogger) Infof(format string, v ...interface{}) {
	v2 := append([]interface{}{"INFO "}, v...)
	l.writef(InfoLevel, "%s"+format, v2...)
}

// Print a message with Debug level.
//
// Available since v0.5.2
func (l DebugLogger) Debug(v ...interface{}) {
	v2 := append([]interface{}{"DEBUG "}, v...)
	l.write(DebugLevel, v2...)
}

// Print a message with Debug level with format.
//
// Available since v0.5.2
func (l DebugLogger) Debugf(format string, v ...interface{}) {
	v2 := append([]interface{}{"DEBUG "}, v...)
	l.writef(DebugLevel, "%s"+format, v2...)
}

// Print a message with Trace level.
//
// Available since v0.5.2
func (l DebugLogger) Trace(v ...interface{}) {
	v2 := append([]interface{}{"TRACE "}, v...)
	l.write(TraceLevel, v2...)
}

// Print a message with Trace level with format.
//
// Available since v0.5.2
func (l DebugLogger) Tracef(format string, v ...interface{}) {
	v2 := append([]interface{}{"TRACE "}, v...)
	l.writef(TraceLevel, "%s"+format, v2...)
}

// Return the context for last log message.
//
// Available since v0.5.2
func (l *DebugLogger) Last() *LogContext {
	l.i.LogMu.Lock()
	defer l.i.LogMu.Unlock()
	v := l.i.Cache.Value
	if v == nil {
		return nil
	}
	return v.(*LogContext)
}

// Return the message for last log message.
//
// Available since v0.5.2
func (l *DebugLogger) LastMessage() string {
	last := l.Last()
	if last == nil {
		return ""
	} else {
		return last.Message
	}
}

// Return the context for all log messages in chronological order.
//
// Available since v0.5.2
func (l *DebugLogger) All() []*LogContext {
	l.i.LogMu.Lock()
	defer l.i.LogMu.Unlock()
	cap := l.i.Cache.Len()
	contexts := []*LogContext{}
	for i := 0; i < cap; i++ {
		l.i.Cache = l.i.Cache.Next()
		if l.i.Cache.Value == nil {
			continue
		}
		contexts = append(contexts, l.i.Cache.Value.(*LogContext))
	}
	return contexts
}

// Return the message for all log messages in chronological order.
//
// Available since v0.5.2
func (l *DebugLogger) AllMessages() []string {
	l.i.LogMu.Lock()
	defer l.i.LogMu.Unlock()
	cap := l.i.Cache.Len()
	messages := []string{}
	for i := 0; i < cap; i++ {
		l.i.Cache = l.i.Cache.Next()
		if l.i.Cache.Value == nil {
			continue
		}
		messages = append(messages, l.i.Cache.Value.(*LogContext).Message)
	}
	return messages
}

// Return the context for all log messages in reverse chronological order.
//
// Available since v0.5.2
func (l *DebugLogger) AllReverse() []*LogContext {
	l.i.LogMu.Lock()
	defer l.i.LogMu.Unlock()
	cap := l.i.Cache.Len()
	contexts := []*LogContext{}
	for i := 0; i < cap; i++ {
		if l.i.Cache.Value == nil {
			break
		}
		contexts = append(contexts, l.i.Cache.Value.(*LogContext))
		l.i.Cache = l.i.Cache.Prev()
	}
	return contexts
}

// Return the message for all log messages in reverse chronological order.
//
// Available since v0.5.2
func (l *DebugLogger) AllMessagesReverse() []string {
	l.i.LogMu.Lock()
	defer l.i.LogMu.Unlock()
	cap := l.i.Cache.Len()
	messages := []string{}
	for i := 0; i < cap; i++ {
		if l.i.Cache.Value == nil {
			break
		}
		messages = append(messages, l.i.Cache.Value.(*LogContext).Message)
		l.i.Cache = l.i.Cache.Prev()
	}
	return messages
}

// Write the message context into cache.
func (l *DebugLogger) write(level LogLevel, v ...interface{}) {
	l.i.LogMu.Lock()
	defer l.i.LogMu.Unlock()
	l.i.Cache = l.i.Cache.Next()
	l.i.Cache.Value = &LogContext{
		Time:    time.Now(),
		Level:   level,
		Message: fmt.Sprint(v...),
	}
}

// Write the message context into cache with format.
func (l *DebugLogger) writef(level LogLevel, format string, v ...interface{}) {
	l.i.LogMu.Lock()
	defer l.i.LogMu.Unlock()
	l.i.Cache = l.i.Cache.Next()
	l.i.Cache.Value = &LogContext{
		Time:    time.Now(),
		Level:   level,
		Message: fmt.Sprintf(format, v...),
	}
}
