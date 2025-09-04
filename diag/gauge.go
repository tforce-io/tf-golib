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

// Gauge is a concurrency-safe struct that represents a single numerical value that can go up
// and go down. This implementation is optimized for a fast Set method.
//
// A Gauge is typically used for measured values like temperatures or current memory usage,
// but also "counts" that can go up and down, like the number of running goroutines.
//
// For Inc-heavy use case. Please consider using Counter instead.
//
// Available since v0.7.0
type Gauge struct {
	// valBits contains the bits of the represented float64 value. It has
	// to go first in the struct to guarantee alignment for atomic
	// operations.
	// http://golang.org/pkg/sync/atomic/#pkg-note-BUG
	valBits uint64
}

// Return new Counter with the given initial value.
//
// Available since v0.7.0
func NewGauge(initVal float64) *Gauge {
	g := &Gauge{}
	g.Set(initVal)
	return g
}

// Add the given value to the gauge.
//
// Available since v0.7.0
func (g *Gauge) Add(v float64) {
	for {
		oldBits := atomic.LoadUint64(&g.valBits)
		newBits := math.Float64bits(math.Float64frombits(oldBits) + v)
		if atomic.CompareAndSwapUint64(&g.valBits, oldBits, newBits) {
			return
		}
	}
}

// Subtract the given value to the gauge.
//
// Available since v0.7.0
func (g *Gauge) Sub(v float64) {
	g.Add(v * -1)
}

// Increment the gauge by 1.
//
// Available since v0.7.0
func (g *Gauge) Inc() {
	g.Add(1)
}

// Decrement the gauge by 1.
//
// Available since v0.7.0
func (g *Gauge) Dec() {
	g.Add(-1)
}

// Set an arbitary value to the gauge.
//
// Available since v0.7.0
func (g *Gauge) Set(v float64) {
	atomic.StoreUint64(&g.valBits, math.Float64bits(v))
}

// Return the current value stored in the gauge.
//
// Available since v0.7.0
func (g *Gauge) Value() float64 {
	return math.Float64frombits(atomic.LoadUint64(&g.valBits))
}
