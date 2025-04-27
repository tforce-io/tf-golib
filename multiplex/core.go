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
	"sync"

	"github.com/tforce-io/tf-golib/diag"
)

const (
	// Number of pending requests a service supports.
	MainChainCapacity = 256
)

// Service interface defines minimum functions for a service.
//
// Available since vTBD
type Service interface {
	// Return Service Identifier string
	//
	// Available since vTBD
	ServiceID() string

	// Return Router for inter-service operability.
	//
	// Available since vTBD
	Router() *ServiceRouter

	// Set number of Process routines the service should use to handle requests.
	//
	// Available since vTBD
	SetWorker(workerCount uint64)

	// Return number of Process routines the service is managing.
	//
	// Available since vTBD
	WorkerCount() uint64

	// Enqueue the request.
	//
	// Available since vTBD
	Exec(command string, params ExecParams)

	// Request other service to handle the request via configurated Router.
	//
	// Available since vTBD
	Dispatch(serviceID string, command string, params ExecParams)
}

// ServiceCore is the base struct for deriving new service.
// New service need to embed ServiceCore to access to pre-defined pattern.
//
// Available since vTBD
type ServiceCore struct {
	i *ServiceCoreInternal
}

// Init ServiceCore internal and return the reference for later access.
//
// Available since vTBD
func (s *ServiceCore) InitServiceCore(serviceID string, logger diag.Logger, processHook func(workerID uint64, msg *ServiceMessage) *HookState) *ServiceCoreInternal {
	if s.i != nil {
		return s.i
	}
	s.i = &ServiceCoreInternal{
		ServiceID:     serviceID,
		MainChan:      make(chan *ServiceMessage, MainChainCapacity),
		WorkerCounter: &Uint64ThreadSafe{},

		Logger: logger,

		CoreProcessHook: processHook,
	}
	return s.i
}

// ServiceCoreInternal stores internal data of a ServiceCore.
//
// Available since vTBD
type ServiceCoreInternal struct {
	ServiceID string
	WorkerID  uint64
	Router    *ServiceRouter

	MainChan chan *ServiceMessage

	WorkerCounter *Uint64ThreadSafe
	WorkerCount   uint64

	Logger diag.Logger

	CoreProcessHook func(workerID uint64, msg *ServiceMessage) *HookState
}

// Return Service Identifier string.
//
// Available since vTBD
func (s ServiceCore) ServiceID() string {
	return s.i.ServiceID
}

// Return Router for inter-service operability.
//
// Available since vTBD
func (s ServiceCore) Router() *ServiceRouter {
	return s.i.Router
}

// Set router to use for inter-service operability.
//
// Available since vTBD
func (s ServiceCore) SetRouter(controller *ServiceController) {
	s.i.Router = controller.router
}

// Set number of Process routines the service should use to handle requests.
//
// Available since vTBD
func (s ServiceCore) SetWorker(workerCount uint64) {
	s.i.WorkerCounter.Lock()
	defer s.i.WorkerCounter.Unlock()
	if s.i.WorkerCounter.ValueNoLock() != s.i.WorkerCount {
		return
	}
	if workerCount > s.i.WorkerCounter.ValueNoLock() {
		for i := s.i.WorkerCounter.ValueNoLock(); i < workerCount; i++ {
			s.i.WorkerID++
			go s.process(s.i.WorkerID)
		}
		s.i.WorkerCount = workerCount
	}
	if workerCount < s.i.WorkerCounter.ValueNoLock() {
		for i := s.i.WorkerCounter.ValueNoLock(); i > workerCount; i-- {
			cmd := &ServiceMessage{
				Command: "exit",
			}
			s.i.MainChan <- cmd
		}
		s.i.WorkerCount = workerCount
	}
}

// Return number of Process routines the service is managing.
//
// Available since vTBD
func (s ServiceCore) WorkerCount() uint64 {
	return s.i.WorkerCount
}

// Enqueue the request.
//
// Available since vTBD
func (s ServiceCore) Exec(command string, params ExecParams) {
	msg := &ServiceMessage{
		Command: command,
		Params:  params,
	}
	s.i.MainChan <- msg
}

// Request other service to handle the request via configurated Router.
//
// Available since vTBD
func (s ServiceCore) Dispatch(serviceID string, command string, params ExecParams) {
	s.i.Router.Forward(serviceID, command, params)
}

// Process routine to handle the request.
//
// Available since vTBD
func (s ServiceCore) process(workerID uint64) {
	s.i.WorkerCounter.Add(1)
	s.i.Logger.Infof("%s#%d Process started.", s.i.ServiceID, workerID)
	status := InitState
	for status != ExitState {
		msg := <-s.i.MainChan
		if s.i.CoreProcessHook != nil {
			hookState := s.i.CoreProcessHook(workerID, msg)
			if hookState.Handled {
				continue
			}
		}
		if msg.Command == "exit" {
			status = ExitState
			continue
		}
	}
	s.i.WorkerCounter.Sub(1)
	s.i.Logger.Infof("%s#%d Process exited.", s.i.ServiceID, workerID)
}

// ServiceMessage defines a request for processing by Service.
//
// Available since vTBD
type ServiceMessage struct {
	Command string
	Params  ExecParams
	Returns *ReturnParams
	Extra   interface{}
}

// Indicate that the request expect returning result.
// This is for sender side.
//
// Available since vTBD
func (p *ServiceMessage) ExpectReturns() {
	p.Returns = &ReturnParams{
		signal: new(sync.WaitGroup),
	}
}

// Set the returning result then signal listener that the request has been completed.
// Nothing will be done if the sender doesn't expect returns.
// This is for recipient side.
//
// Available since vTBD
func (p *ServiceMessage) CompleteReturns(result interface{}) {
	if p.Returns != nil {
		p.Returns.result = result
		p.Returns.signal.Done()
	}
}

// Listen to the signal and return received result.
// The routine won't be blocked and receive nil if it doesn't expect returns.
// This is for sender side.
//
// Available since vTBD
func (p *ServiceMessage) WaitForReturns() interface{} {
	if p.Returns != nil {
		p.Returns.signal.Wait()
		return p.Returns.result
	}
	return nil
}

// Collection of parameters as key-value mapping.
//
// Available since vTBD
type ExecParams map[string]interface{}

// Return parameter value if any, and a bool indicate if the key exists.
//
// Available since vTBD
func (p ExecParams) Get(key string) (interface{}, bool) {
	if val, ok := p[key]; ok {
		return val, true
	}
	return nil, false
}

// Set parameter with specified key.
//
// Available since vTBD
func (p ExecParams) Set(key string, val interface{}) {
	p[key] = val
}

// Delete parameter with specified key.
//
// Available since vTBD
func (p ExecParams) Delete(key string) {
	delete(p, key)
}

// ReturnParams comprises of a dyanmic type result and a signal for synchronous support.
//
// Available since vTBD
type ReturnParams struct {
	signal *sync.WaitGroup
	result interface{}
}
