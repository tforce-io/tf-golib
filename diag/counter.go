// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package diag

import (
	"math"
	"sync/atomic"
)

// Counter is a concurrency-safe struct that represents a single numerical value that can go up
// and go down. This implementation is optimized for a fast Inc/Dec/Add/Sub methods.
//
// A Counter is typically used to count requests served, tasks completed, errors occurred, etc.
//
// For Set-heavy use case. Please consider using Gauge instead.
//
// Available since v0.7.0
type Counter struct {
	// valBits contains the bits of the represented float64 value, while
	// valInt stores values that are exact integers. Both have to go first
	// in the struct to guarantee alignment for atomic operations.
	// http://golang.org/pkg/sync/atomic/#pkg-note-BUG
	valBits uint64
	valInt  int64
}

// Return new Counter with the given initial value.
//
// Available since v0.7.0
func NewCounter(initVal float64) *Counter {
	c := &Counter{}
	c.Add(initVal)
	return c
}

// Add the given value to the counter.
//
// Available since v0.7.0
func (c *Counter) Add(v float64) {
	ival := int64(v)
	if float64(ival) == v {
		atomic.AddInt64(&c.valInt, ival)
		return
	}

	for {
		oldBits := atomic.LoadUint64(&c.valBits)
		newBits := math.Float64bits(math.Float64frombits(oldBits) + v)
		if atomic.CompareAndSwapUint64(&c.valBits, oldBits, newBits) {
			return
		}
	}
}

// Subtract the given value to the counter.
//
// Available since v0.7.0
func (c *Counter) Sub(v float64) {
	c.Add(v * -1)
}

// Increment the counter by 1.
//
// Available since v0.7.0
func (c *Counter) Inc() {
	c.Add(1)
}

// Decrement the counter by 1.
//
// Available since v0.7.0
func (c *Counter) Dec() {
	c.Add(-1)
}

// Set an arbitary value to the counter.
//
// Available since v0.7.0
func (c *Counter) Set(v float64) {
	for {
		oldBits := atomic.LoadUint64(&c.valBits)
		ival := atomic.LoadInt64(&c.valInt)
		newBits := math.Float64bits(v - float64(ival))
		if atomic.CompareAndSwapUint64(&c.valBits, oldBits, newBits) {
			return
		}
	}
}

// Return the current value stored in the counter.
//
// Available since v0.7.0
func (c *Counter) Value() float64 {
	fval := math.Float64frombits(atomic.LoadUint64(&c.valBits))
	ival := atomic.LoadInt64(&c.valInt)
	return fval + float64(ival)
}
