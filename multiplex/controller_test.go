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

package multiplex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tforce-io/tf-golib/diag"
)

func TestServiceController_Register(t *testing.T) {
	logger := diag.NewDebugLogger(10)
	svc := NewServiceController(logger)
	echo := NewEchoService(logger)
	isRegistered := svc.Register(echo)
	assert.False(t, isRegistered, "service can't be registered without correct router")
	echo.SetRouter(svc)
	isRegistered = svc.Register(echo)
	assert.True(t, isRegistered)
}

func TestServiceController_Unregister(t *testing.T) {
	logger := diag.NewDebugLogger(10)
	svc := NewServiceController(logger)
	echo := NewEchoService(logger)
	echo.SetRouter(svc)
	svc.Register(echo)
	_, found := svc.services[echo.ServiceID()]
	assert.True(t, found, "service should be registered")
	svc.Unregister(echo.ServiceID())
	_, found = svc.services[echo.ServiceID()]
	assert.False(t, found, "service must be unregistered")
	svc.Unregister(echo.ServiceID())
	_, found = svc.services[echo.ServiceID()]
	assert.False(t, found, "service must be unregistered")
}

func TestServiceController_Run_Background(t *testing.T) {
	logger := diag.NewDebugLogger(10)
	svc := NewServiceController(logger)
	svc.SetWorker(1)
	go func() {
		time.Sleep(100 * time.Millisecond)
		svc.Exec("exit", ExecParams{})
	}()
	svc.Run(true)
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "INFO Controller#1: Process exited.", logger.LastMessage(), "routine not exited properly")
}

func TestServiceController_Dispatch(t *testing.T) {
	logger := diag.NewDebugLogger(10)
	svc := NewServiceController(logger)
	svc.SetWorker(1)
	logger2 := diag.NewDebugLogger(10)
	echo := NewEchoService(logger2)
	echo.SetWorker(1)
	echo.SetRouter(svc)
	svc.Register(echo)
	svc.Dispatch(echo.ServiceID(), "", ExecParams{"message": "Hello, World!"})
	go func() {
		time.Sleep(100 * time.Millisecond)
		svc.Exec("exit", ExecParams{})
	}()
	svc.Run(true)
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "INFO Controller#1: Process exited.", logger.LastMessage(), "routine not exited properly")
	assert.Equal(t, "INFO Message received: Hello, World!", logger2.LastMessage(), "message is not dispatched properly")
}
