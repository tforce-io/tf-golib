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
)

// ServiceMessage defines a request for processing by Service.
//
// Available since v0.5.0
type ServiceMessage struct {
	Command string
	Params  ExecParams
	Extra   interface{}
}

// Collection of parameters as key-value mapping.
//
// Available since v0.5.0
type ExecParams map[string]interface{}

// Return parameter value if any, or fallback to def.
//
// Available since vTBD
func (p ExecParams) Get(key string, def interface{}) interface{} {
	if val, ok := p[key]; ok {
		return val
	}
	return def
}

// Set parameter with specified key.
//
// Available since v0.5.0
func (p ExecParams) Set(key string, val interface{}) {
	p[key] = val
}

// Delete parameter with specified key.
//
// Available since v0.5.0
func (p ExecParams) Delete(key string) {
	delete(p, key)
}

// Indicate that the request expect returning result.
// This is for sender side.
//
// Available since vTBD
func (p ExecParams) ExpectReturn() {
	signal := new(sync.WaitGroup)
	signal.Add(1)
	p.ExpectReturnCustomSignal(signal)
}

// Indicate that the request expect returning result using a custom signal.
// This is for sender side.
//
// Available since vTBD
func (p ExecParams) ExpectReturnCustomSignal(signal *sync.WaitGroup) {
	p["return"] = &ReturnParams{
		signal: signal,
	}
}

// Set the returning result then signal listener that the request has been completed.
// Nothing will be done if the sender doesn't expect returns.
// This is for recipient side.
//
// Available since vTBD
func (p ExecParams) Return(result interface{}) {
	if p["return"] != nil {
		ret := p.Get("return", nil).(*ReturnParams)
		ret.result = result
		ret.signal.Done()
	}
}

// Listen to the signal.
// The routine won't be blocked and receive nil if it doesn't expect returns.
// This is for sender side.
//
// Available since vTBD
func (p ExecParams) Wait() {
	if p["return"] != nil {
		ret := p.Get("return", nil).(*ReturnParams)
		ret.signal.Wait()
	}
}

// Listen to the signal and return received result.
// The routine won't be blocked and receive nil if it doesn't expect returns.
// This is for sender side.
//
// Available since vTBD
func (p ExecParams) WaitForReturn() interface{} {
	if p["return"] != nil {
		ret := p.Get("return", nil).(*ReturnParams)
		ret.signal.Wait()
		return ret.result
	}
	return nil
}

// ReturnParams comprises of a dyanmic type result and a signal for synchronous support.
//
// Available since v0.5.0
type ReturnParams struct {
	signal *sync.WaitGroup
	result interface{}
}

// Return Singal of the param.
//
// Available since v0.5.1
func (p *ReturnParams) Signal() *sync.WaitGroup {
	return p.signal
}

// Return Result of the param.
//
// Available since v0.5.1
func (p *ReturnParams) Result() interface{} {
	return p.result
}
