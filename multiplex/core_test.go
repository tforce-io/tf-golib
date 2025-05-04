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

func TestServiceCore_ServiceID(t *testing.T) {
	logger := diag.NewDebugLogger(10)
	svc := NewRandomService(logger)
	assert.Equal(t, "Random", svc.ServiceID(), "invalid service ID")
}

func TestServiceCore_Router(t *testing.T) {
	logger := diag.NewDebugLogger(10)
	svc := NewRandomService(logger)
	assert.Nil(t, svc.Router(), "unexpected value for router, must be nil")
	controller := NewServiceController(logger)
	svc.SetRouter(controller)
	assert.NotNil(t, svc.Router(), "unexpected value for router, must not be nil")
}

func TestServiceCore_Worker(t *testing.T) {
	logger := diag.NewDebugLogger(10)
	svc := NewRandomService(logger)
	assert.Equal(t, uint64(0), svc.WorkerCount(), "mismatch worker count")
	svc.SetWorker(17)
	assert.Equal(t, uint64(17), svc.WorkerCount(), "mismatch worker count")
	time.Sleep(100 * time.Millisecond)
	svc.SetWorker(10)
	assert.Equal(t, uint64(10), svc.WorkerCount(), "mismatch worker count")
}

func TestServiceCore_Exec(t *testing.T) {
	logger := diag.NewDebugLogger(10)
	svc := NewEchoService(logger)
	svc.SetWorker(1)
	svc.Exec("", ExecParams{
		"message": "Hello, World!",
	})
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, "INFO Message received: Hello, World!", logger.LastMessage(), "invalid message")
}

type EchoService struct {
	ServiceCore
	i *ServiceCoreInternal
}

func NewEchoService(logger diag.Logger) *EchoService {
	svc := &EchoService{}
	svc.i = svc.InitServiceCore("Echo", logger, svc.coreProcessHook)
	return svc
}

func (s *EchoService) coreProcessHook(workerID uint64, msg *ServiceMessage) *HookState {
	message := msg.Params["message"].(string)
	s.i.Logger.Infof("Message received: %s", message)
	return &HookState{Handled: true}
}

type RandomService struct {
	ServiceCore
	i *ServiceCoreInternal
}

func NewRandomService(logger diag.Logger) *RandomService {
	svc := &RandomService{}
	svc.i = svc.InitServiceCore("Random", logger, svc.coreProcessHook)
	return svc
}

func (s *RandomService) coreProcessHook(workerID uint64, msg *ServiceMessage) *HookState {
	return &HookState{Handled: true}
}
