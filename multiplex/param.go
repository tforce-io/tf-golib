// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

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

// Return parameter value if any, or fallback to def.
//
// Available since v0.5.2
func (m *ServiceMessage) GetParam(key string, def interface{}) interface{} {
	if m.Params == nil {
		return def
	}
	return m.Params.Get(key, def)
}

// Set parameter with specified key.
//
// Available since v0.5.2
func (m *ServiceMessage) SetParam(key string, val interface{}) {
	if m.Params == nil {
		m.Params = make(ExecParams)
	}
	m.Params.Set(key, val)
}

// Delete parameter with specified key.
//
// Available since v0.5.2
func (m *ServiceMessage) DeleteParam(key string) {
	if m.Params != nil {
		m.Params.Delete(key)
	}
}

// Return Singal of the param.
//
// Available since v0.5.2
func (m *ServiceMessage) ReturnSignal() *sync.WaitGroup {
	if m.Params != nil {
		return m.Params.ReturnSignal()
	}
	return nil
}

// Return Result of the param.
//
// Available since v0.5.2
func (m *ServiceMessage) ReturnResult() interface{} {
	if m.Params != nil {
		return m.Params.ReturnResult()
	}
	return nil
}

// Indicate that the request expect returning result.
// This is for sender side.
//
// Available since v0.5.2
func (m *ServiceMessage) ExpectReturn() {
	if m.Params == nil {
		m.Params = make(ExecParams)
	}
	m.Params.ExpectReturn()
}

// Indicate that the request expect returning result using a custom signal.
// This is for sender side.
//
// Available since v0.5.2
func (m *ServiceMessage) ExpectReturnCustomSignal(signal *sync.WaitGroup) {
	if m.Params == nil {
		m.Params = make(ExecParams)
	}
	m.Params.ExpectReturnCustomSignal(signal)
}

// Set the returning result then signal listener that the request has been completed.
// Nothing will be done if the sender doesn't expect returns.
// This is for recipient side.
//
// Available since v0.5.2
func (m *ServiceMessage) Return(result interface{}) {
	if m.Params != nil {
		m.Params.Return(result)
	}
}

// Listen to the signal.
// The routine won't be blocked and receive nil if it doesn't expect returns.
// This is for sender side.
//
// Available since v0.5.2
func (m *ServiceMessage) Wait() {
	if m.Params != nil {
		m.Params.Wait()
	}
}

// Listen to the signal and return received result.
// The routine won't be blocked and receive nil if it doesn't expect returns.
// This is for sender side.
//
// Available since v0.5.0
func (m *ServiceMessage) WaitForReturn() interface{} {
	if m.Params != nil {
		return m.Params.WaitForReturn()
	}
	return nil
}

// Collection of parameters as key-value mapping.
//
// Available since v0.5.0
type ExecParams map[string]interface{}

// Return parameter value if any, or fallback to def.
//
// Available since v0.5.2
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

// Return Singal of the param.
//
// Available since v0.5.2
func (p ExecParams) ReturnSignal() *sync.WaitGroup {
	if p["return"] != nil {
		ret := p.Get("return", nil).(*ReturnParams)
		return ret.signal
	}
	return nil
}

// Return Result of the param.
//
// Available since v0.5.2
func (p ExecParams) ReturnResult() interface{} {
	if p["return"] != nil {
		ret := p.Get("return", nil).(*ReturnParams)
		return ret.result
	}
	return nil
}

// Indicate that the request expect returning result.
// This is for sender side.
//
// Available since v0.5.2
func (p ExecParams) ExpectReturn() {
	signal := new(sync.WaitGroup)
	signal.Add(1)
	p.ExpectReturnCustomSignal(signal)
}

// Indicate that the request expect returning result using a custom signal.
// This is for sender side.
//
// Available since v0.5.2
func (p ExecParams) ExpectReturnCustomSignal(signal *sync.WaitGroup) {
	p["return"] = &ReturnParams{
		signal: signal,
	}
}

// Set the returning result then signal listener that the request has been completed.
// Nothing will be done if the sender doesn't expect returns.
// This is for recipient side.
//
// Available since v0.5.2
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
// Available since v0.5.2
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
// Available since v0.5.2
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
