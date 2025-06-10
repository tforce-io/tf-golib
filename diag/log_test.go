// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package diag

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"testing"
)

func TestDefaultLogger_Error(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureStdout(func() {
		logger.Error(errors.New("invalid"), "Message")
	})
	if !strings.HasSuffix(output, "ERROR invalid Message\n") {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDefaultLogger_Errorf(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureStdout(func() {
		logger.Errorf(errors.New("invalid"), "%s", "Messagef")
	})
	if !strings.HasSuffix(output, "ERROR invalid Messagef\n") {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDefaultLogger_Warn(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureStdout(func() {
		logger.Warn("Message")
	})
	if !strings.HasSuffix(output, "WARN Message\n") {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDefaultLogger_Warnf(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureStdout(func() {
		logger.Warnf("%s", "Messagef")
	})
	if !strings.HasSuffix(output, "WARN Messagef\n") {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDefaultLogger_Info(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureStdout(func() {
		logger.Info("Message")
	})
	if !strings.HasSuffix(output, "INFO Message\n") {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDefaultLogger_Infof(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureStdout(func() {
		logger.Infof("%s", "Messagef")
	})
	if !strings.HasSuffix(output, "INFO Messagef\n") {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDefaultLogger_Debug(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureStdout(func() {
		logger.Debug("Message")
	})
	if !strings.HasSuffix(output, "DEBUG Message\n") {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDefaultLogger_Debugf(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureStdout(func() {
		logger.Debugf("%s", "Messagef")
	})
	if !strings.HasSuffix(output, "DEBUG Messagef\n") {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDefaultLogger_Trace(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureStdout(func() {
		logger.Trace("Message")
	})
	if !strings.HasSuffix(output, "TRACE Message\n") {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func TestDefaultLogger_Tracef(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureStdout(func() {
		logger.Tracef("%s", "Messagef")
	})
	if !strings.HasSuffix(output, "TRACE Messagef\n") {
		t.Errorf("Unexpected log output: %s", output)
	}
}

func captureStdout(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)
	f()
	return buf.String()
}
