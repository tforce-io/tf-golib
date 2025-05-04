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
	"errors"
	"testing"
)

func TestDebugLogger_Error(t *testing.T) {
	logger := NewDebugLogger(10)
	logger.Error(errors.New("invalid"), "Message")
	output := logger.LastMessage()
	if output != "ERROR invalid Message" {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDebugLogger_Errorf(t *testing.T) {
	logger := NewDebugLogger(10)
	logger.Errorf(errors.New("invalid"), "%s", "Messagef")
	output := logger.LastMessage()
	if output != "ERROR invalid Messagef" {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDebugLogger_Warn(t *testing.T) {
	logger := NewDebugLogger(10)
	logger.Warn("Message")
	output := logger.LastMessage()
	if output != "WARN Message" {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDebugLogger_Warnf(t *testing.T) {
	logger := NewDebugLogger(10)
	logger.Warnf("%s", "Messagef")
	output := logger.LastMessage()
	if output != "WARN Messagef" {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDebugLogger_Info(t *testing.T) {
	logger := NewDebugLogger(10)
	logger.Info("Message")
	output := logger.LastMessage()
	if output != "INFO Message" {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDebugLogger_Infof(t *testing.T) {
	logger := NewDebugLogger(10)
	logger.Infof("%s", "Messagef")
	output := logger.LastMessage()
	if output != "INFO Messagef" {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDebugLogger_Debug(t *testing.T) {
	logger := NewDebugLogger(10)
	logger.Debug("Message")
	output := logger.LastMessage()
	if output != "DEBUG Message" {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDebugLogger_Debugf(t *testing.T) {
	logger := NewDebugLogger(10)
	logger.Debugf("%s", "Messagef")
	output := logger.LastMessage()
	if output != "DEBUG Messagef" {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDebugLogger_Trace(t *testing.T) {
	logger := NewDebugLogger(10)
	logger.Trace("Message")
	output := logger.LastMessage()
	if output != "TRACE Message" {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDebugLogger_Tracef(t *testing.T) {
	logger := NewDebugLogger(10)
	logger.Tracef("%s", "Messagef")
	output := logger.LastMessage()
	if output != "TRACE Messagef" {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDebugLogger_Last_Empty(t *testing.T) {
	logger := NewDebugLogger(10)
	output := logger.Last()
	if output != nil {
		t.Errorf("Expected nil output for empty logger, got: %v", output)
	}
}

func TestDebugLogger_LastMessage_Empty(t *testing.T) {
	logger := NewDebugLogger(10)
	output := logger.LastMessage()
	if output != "" {
		t.Errorf("Expected empty output for empty logger, got: %v", output)
	}
}

func TestDebugLogger_All(t *testing.T) {
	logger := NewDebugLogger(5)
	logger.Info("Message 1")
	logger.Warn("Message 2")
	logger.Debug("Message 3")
	logger.Info("Message 4")

	contexts := logger.All()
	if len(contexts) != 4 {
		t.Errorf("Expected 4 log contexts, got: %d", len(contexts))
	}

	expectedMessages := []string{"INFO Message 1", "WARN Message 2", "DEBUG Message 3", "INFO Message 4"}
	for i, ctx := range contexts {
		if ctx.Message != expectedMessages[i] {
			t.Errorf("Expected message '%s', got: '%s'", expectedMessages[i], ctx.Message)
		}
	}
}

func TestDebugLogger_All_Overflown(t *testing.T) {
	logger := NewDebugLogger(3)
	logger.Info("Message 1")
	logger.Warn("Message 2")
	logger.Debug("Message 3")
	logger.Info("Message 4")

	contexts := logger.All()
	if len(contexts) != 3 {
		t.Errorf("Expected 3 log contexts, got: %d", len(contexts))
	}

	expectedMessages := []string{"WARN Message 2", "DEBUG Message 3", "INFO Message 4"}
	for i, ctx := range contexts {
		if ctx.Message != expectedMessages[i] {
			t.Errorf("Expected message '%s', got: '%s'", expectedMessages[i], ctx.Message)
		}
	}
}

func TestDebugLogger_AllMessages(t *testing.T) {
	logger := NewDebugLogger(5)
	logger.Info("Message 1")
	logger.Warn("Message 2")
	logger.Debug("Message 3")
	logger.Info("Message 4")

	messages := logger.AllMessages()
	if len(messages) != 4 {
		t.Errorf("Expected 4 log messages, got: %d", len(messages))
	}

	expectedMessages := []string{"INFO Message 1", "WARN Message 2", "DEBUG Message 3", "INFO Message 4"}
	for i, msg := range messages {
		if msg != expectedMessages[i] {
			t.Errorf("Expected message '%s', got: '%s'", expectedMessages[i], msg)
		}
	}
}

func TestDebugLogger_AllMessages_Overflown(t *testing.T) {
	logger := NewDebugLogger(3)
	logger.Info("Message 1")
	logger.Warn("Message 2")
	logger.Debug("Message 3")
	logger.Info("Message 4")

	messages := logger.AllMessages()
	if len(messages) != 3 {
		t.Errorf("Expected 3 log messages, got: %d", len(messages))
	}

	expectedMessages := []string{"WARN Message 2", "DEBUG Message 3", "INFO Message 4"}
	for i, msg := range messages {
		if msg != expectedMessages[i] {
			t.Errorf("Expected message '%s', got: '%s'", expectedMessages[i], msg)
		}
	}
}

func TestDebugLogger_AllReverse(t *testing.T) {
	logger := NewDebugLogger(5)
	logger.Info("Message 1")
	logger.Warn("Message 2")
	logger.Debug("Message 3")
	logger.Info("Message 4")

	contexts := logger.AllReverse()
	if len(contexts) != 4 {
		t.Errorf("Expected 4 log contexts, got: %d", len(contexts))
	}

	expectedMessages := []string{"INFO Message 4", "DEBUG Message 3", "WARN Message 2", "INFO Message 1"}
	for i, ctx := range contexts {
		if ctx.Message != expectedMessages[i] {
			t.Errorf("Expected message '%s', got: '%s'", expectedMessages[i], ctx.Message)
		}
	}
}

func TestDebugLogger_AllReverse_Overflown(t *testing.T) {
	logger := NewDebugLogger(3)
	logger.Info("Message 1")
	logger.Warn("Message 2")
	logger.Debug("Message 3")
	logger.Info("Message 4")

	contexts := logger.AllReverse()
	if len(contexts) != 3 {
		t.Errorf("Expected 3 log contexts, got: %d", len(contexts))
	}

	expectedMessages := []string{"INFO Message 4", "DEBUG Message 3", "WARN Message 2"}
	for i, ctx := range contexts {
		if ctx.Message != expectedMessages[i] {
			t.Errorf("Expected message '%s', got: '%s'", expectedMessages[i], ctx.Message)
		}
	}
}

func TestDebugLogger_AllMessagesReverse(t *testing.T) {
	logger := NewDebugLogger(5)
	logger.Info("Message 1")
	logger.Warn("Message 2")
	logger.Debug("Message 3")
	logger.Info("Message 4")

	messages := logger.AllMessagesReverse()
	if len(messages) != 4 {
		t.Errorf("Expected 4 log messages, got: %d", len(messages))
	}

	expectedMessages := []string{"INFO Message 4", "DEBUG Message 3", "WARN Message 2", "INFO Message 1"}
	for i, msg := range messages {
		if msg != expectedMessages[i] {
			t.Errorf("Expected message '%s', got: '%s'", expectedMessages[i], msg)
		}
	}
}

func TestDebugLogger_AllMessagesReverse_Overflown(t *testing.T) {
	logger := NewDebugLogger(3)
	logger.Info("Message 1")
	logger.Warn("Message 2")
	logger.Debug("Message 3")
	logger.Info("Message 4")

	messages := logger.AllMessagesReverse()
	if len(messages) != 3 {
		t.Errorf("Expected 3 log messages, got: %d", len(messages))
	}

	expectedMessages := []string{"INFO Message 4", "DEBUG Message 3", "WARN Message 2"}
	for i, msg := range messages {
		if msg != expectedMessages[i] {
			t.Errorf("Expected message '%s', got: '%s'", expectedMessages[i], msg)
		}
	}
}
