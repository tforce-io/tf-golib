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

import "github.com/tforce-io/tf-golib/diag"

// ServiceController is a controller for managing services and routing messages between them.
//
// Available since v0.5.0
type ServiceController struct {
	ServiceCore
	services map[string]Service
}

// Return new ServiceRouter.
//
// Available since v0.5.0
func NewServiceController(logger diag.Logger) *ServiceController {
	svc := &ServiceController{}
	svc.InitServiceCore("Controller", logger, svc.coreProcessHook)
	svc.i.Router = &ServiceRouter{
		c: svc,
	}
	svc.services = make(map[string]Service)
	return svc
}

// Register service for routing.
//
// Available since v0.5.0
func (s *ServiceController) Register(service Service) bool {
	if service.Router() != s.Router() {
		s.i.Logger.Warn("Service %s's router is invalid", service.ServiceID())
		return false
	}
	s.services[service.ServiceID()] = service
	return true
}

// Unregister a service by serviceID.
//
// Available since v0.5.0
func (s *ServiceController) Unregister(serviceID string) {
	delete(s.services, serviceID)
}

// Run the controller.
//
// Available since v0.5.0
func (s *ServiceController) Run(background bool) {
	s.SetWorker(1)

	if background {
		s.i.Background = true
		<-s.i.ExitChan
		s.i.Background = false
	}
}

// coreProcessHook is responsible for processing messages in the controller.
//
// Available since v0.5.0
func (s *ServiceController) coreProcessHook(workerID uint64, msg *ServiceMessage) *HookState {
	if msg.Extra == nil && msg.Command == "exit" {
		return &HookState{Handled: false}
	}
	serviceID := msg.Extra.(*ControllerExtra).ServiceID
	s.services[serviceID].Exec(msg.Command, msg.Params)
	return &HookState{Handled: true}
}

// ServiceRouter is responsible for routing messages between services.
//
// Available since v0.5.0
type ServiceRouter struct {
	c *ServiceController
}

// Forward the message to the specified serviceID.
//
// Available since v0.5.0
func (s *ServiceRouter) Forward(serviceID, command string, params ExecParams) {
	msg := &ServiceMessage{
		Command: command,
		Params:  params,
	}
	if serviceID != "" {
		msg.Extra = &ControllerExtra{
			ServiceID: serviceID,
		}
	}
	s.c.i.MainChan <- msg
}

// ControllerExtra contains additional information for request to the controller.
//
// Available since v0.5.0
type ControllerExtra struct {
	ServiceID string
}
