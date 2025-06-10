// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package multiplex

import (
	"github.com/tforce-io/tf-golib/diag"
)

const (
	// Number of pending requests a service supports.
	MainChainCapacity = 256
	// Number of pending requests a service supports.
	ExtraChanCapacity = 16
)

// Service interface defines minimum functions for a service.
//
// Available since v0.5.0
type Service interface {
	// Return Service Identifier string
	//
	// Available since v0.5.0
	ServiceID() string

	// Return Router for inter-service operability.
	//
	// Available since v0.5.0
	Router() *ServiceRouter

	// Set number of Process routines the service should use to handle requests.
	//
	// Available since v0.5.0
	SetWorker(workerCount uint64)

	// Return number of Process routines the service is managing.
	//
	// Available since v0.5.0
	WorkerCount() uint64

	// Enqueue the request.
	//
	// Available since v0.5.0
	Exec(command string, params ExecParams)

	// Request other service to handle the request via configurated Router.
	//
	// Available since v0.5.0
	Dispatch(serviceID string, command string, params ExecParams)
}

// ServiceCore is the base struct for deriving new service.
// New service need to embed ServiceCore to access to pre-defined pattern.
//
// Available since v0.5.0
type ServiceCore struct {
	i *ServiceCoreInternal
}

// ServiceCoreInternal stores internal data of a ServiceCore.
//
// Available since v0.5.0
type ServiceCoreInternal struct {
	ServiceID string
	WorkerID  uint64
	Router    *ServiceRouter

	MainChan   chan *ServiceMessage
	ExitChan   chan bool
	Background bool

	WorkerCounter *Uint64ThreadSafe
	WorkerCount   uint64

	Logger diag.Logger

	CoreProcessHook func(workerID uint64, msg *ServiceMessage) *HookState
}

// Init ServiceCore internal and return the reference for later access.
//
// Available since v0.5.0
func (s *ServiceCore) InitServiceCore(serviceID string, logger diag.Logger, processHook func(workerID uint64, msg *ServiceMessage) *HookState) *ServiceCoreInternal {
	if s.i != nil {
		return s.i
	}
	s.i = &ServiceCoreInternal{
		ServiceID:     serviceID,
		MainChan:      make(chan *ServiceMessage, MainChainCapacity),
		ExitChan:      make(chan bool, ExtraChanCapacity),
		WorkerCounter: &Uint64ThreadSafe{},

		Logger: logger,

		CoreProcessHook: processHook,
	}
	return s.i
}

// Return Service Identifier string.
//
// Available since v0.5.0
func (s ServiceCore) ServiceID() string {
	return s.i.ServiceID
}

// Return Router for inter-service operability.
//
// Available since v0.5.0
func (s ServiceCore) Router() *ServiceRouter {
	return s.i.Router
}

// Set router to use for inter-service operability.
//
// Available since v0.5.0
func (s ServiceCore) SetRouter(controller *ServiceController) {
	s.i.Router = controller.i.Router
}

// Set number of Process routines the service should use to handle requests.
//
// Available since v0.5.0
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
// Available since v0.5.0
func (s ServiceCore) WorkerCount() uint64 {
	return s.i.WorkerCount
}

// Enqueue the request.
//
// Available since v0.5.0
func (s ServiceCore) Exec(command string, params ExecParams) {
	msg := &ServiceMessage{
		Command: command,
		Params:  params,
	}
	s.i.MainChan <- msg
}

// Request other service to handle the request via configurated Router.
//
// Available since v0.5.0
func (s ServiceCore) Dispatch(serviceID string, command string, params ExecParams) {
	s.i.Router.Forward(serviceID, command, params)
}

// Process routine to handle the request.
//
// Available since v0.5.0
func (s ServiceCore) process(workerID uint64) {
	s.i.WorkerCounter.Add(1)
	s.i.Logger.Infof("%s#%d: Process started.", s.i.ServiceID, workerID)
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
	s.i.Logger.Infof("%s#%d: Process exited.", s.i.ServiceID, workerID)
	if s.i.Background {
		s.i.ExitChan <- true
	}
}
