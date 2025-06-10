// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package multiplex

import "sync"

// Struct Int64ThreadSafe is a thread-safe counter with underlying value of int64.
//
// Available since v0.5.0
type Int64ThreadSafe struct {
	value int64
	lock  sync.Mutex
}

// Return current value exclusively.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) Value() int64 {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.value
}

// Return current value.
// Recommended to use with Lock() and Unlock() to prevent unexpected result.
// Maybe inconsistent if there are many routines updating its value at the same time.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) ValueNoLock() int64 {
	return c.value
}

// Set value eclusively.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) Set(value int64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value = value
}

// Set value.
// Recommended to use with Lock() and Unlock() to prevent unexpected result.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) SetNoLock(value int64) {
	c.value = value
}

// Add current value with n exclusively.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) Add(n int64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value += n
}

// AAdd current value with n.
// Recommended to use with Lock() and Unlock() to prevent unexpected result.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) AddNoLock(n int64) {
	c.value += n
}

// Subtract current value with n exclusively.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) Sub(n int64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value -= n
}

// Subtract current value with n.
// Recommended to use with Lock() and Unlock() to prevent unexpected result.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) SubNoLock(n int64) {
	c.value -= n
}

// Multiply current value with n exclusively.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) Mul(n int64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value *= n
}

// Multiply current value with n.
// Recommended to use with Lock() and Unlock() to prevent unexpected result.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) MulNoLock(value int64) {
	c.value *= value
}

// Divide current value with n exclusively.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) Div(n int64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value /= n
}

// Divide current value with n.
// Recommended to use with Lock() and Unlock() to prevent unexpected result.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) DivNoLock(value int64) {
	c.value /= value
}

// Acquire the lock manually.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) Lock() {
	c.lock.Lock()
}

// Release the lock manually.
//
// Available since v0.5.0
func (c *Int64ThreadSafe) Unlock() {
	c.lock.Unlock()
}
