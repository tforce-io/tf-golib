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

func TestDefaultLogger_Error(t *testing.T) {
	logger := &DefaultLogger{}
	logger.Error(errors.New("invalid"), "Message")
}

func TestDefaultLogger_Errorf(t *testing.T) {
	logger := &DefaultLogger{}
	logger.Errorf(errors.New("invalid"), "%s", "Messagef")
}

func TestDefaultLogger_Warn(t *testing.T) {
	logger := &DefaultLogger{}
	logger.Warn("Message")
}

func TestDefaultLogger_Warnf(t *testing.T) {
	logger := &DefaultLogger{}
	logger.Warnf("%s", "Messagef")
}

func TestDefaultLogger_Info(t *testing.T) {
	logger := &DefaultLogger{}
	logger.Info("Message")
}

func TestDefaultLogger_Infof(t *testing.T) {
	logger := &DefaultLogger{}
	logger.Infof("%s", "Messagef")
}

func TestDefaultLogger_Debug(t *testing.T) {
	logger := &DefaultLogger{}
	logger.Debug("Message")
}

func TestDefaultLogger_Debugf(t *testing.T) {
	logger := &DefaultLogger{}
	logger.Debugf("%s", "Messagef")
}

func TestDefaultLogger_Trace(t *testing.T) {
	logger := &DefaultLogger{}
	logger.Trace("Message")
}

func TestDefaultLogger_Tracef(t *testing.T) {
	logger := &DefaultLogger{}
	logger.Tracef("%s", "Messagef")
}
