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

import "sync"

// Struct Uint64ThreadSafe is a thread-safe counter with underlying value of uint64.
//
// Available since vTBD
type Uint64ThreadSafe struct {
	value uint64
	lock  sync.Mutex
}

// Return current value exclusively.
//
// Available since vTBD
func (c *Uint64ThreadSafe) Value() uint64 {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.value
}

// Return current value.
// Recommended to use with Lock() and Unlock() to prevent unexpected result.
// Maybe inconsistent if there are many routines updating its value at the same time.
//
// Available since vTBD
func (c *Uint64ThreadSafe) ValueNoLock() uint64 {
	return c.value
}

// Set value eclusively.
//
// Available since vTBD
func (c *Uint64ThreadSafe) Set(value uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value = value
}

// Set value.
// Recommended to use with Lock() and Unlock() to prevent unexpected result.
//
// Available since vTBD
func (c *Uint64ThreadSafe) SetNoLock(value uint64) {
	c.value = value
}

// Add current value with n exclusively.
//
// Available since vTBD
func (c *Uint64ThreadSafe) Add(n uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value += n
}

// AAdd current value with n.
// Recommended to use with Lock() and Unlock() to prevent unexpected result.
//
// Available since vTBD
func (c *Uint64ThreadSafe) AddNoLock(n uint64) {
	c.value += n
}

// Subtract current value with n exclusively.
//
// Available since vTBD
func (c *Uint64ThreadSafe) Sub(n uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value -= n
}

// Subtract current value with n.
// Recommended to use with Lock() and Unlock() to prevent unexpected result.
//
// Available since vTBD
func (c *Uint64ThreadSafe) SubNoLock(n uint64) {
	c.value -= n
}

// Multiply current value with n exclusively.
//
// Available since vTBD
func (c *Uint64ThreadSafe) Mul(n uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value *= n
}

// Multiply current value with n.
// Recommended to use with Lock() and Unlock() to prevent unexpected result.
//
// Available since vTBD
func (c *Uint64ThreadSafe) MulNoLock(value uint64) {
	c.value *= value
}

// Divide current value with n exclusively.
//
// Available since vTBD
func (c *Uint64ThreadSafe) Div(n uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value /= n
}

// Divide current value with n.
// Recommended to use with Lock() and Unlock() to prevent unexpected result.
//
// Available since vTBD
func (c *Uint64ThreadSafe) DivNoLock(value uint64) {
	c.value /= value
}

// Acquire the lock manually.
//
// Available since vTBD
func (c *Uint64ThreadSafe) Lock() {
	c.lock.Lock()
}

// Release the lock manually.
//
// Available since vTBD
func (c *Uint64ThreadSafe) Unlock() {
	c.lock.Unlock()
}
