// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package diag

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewProgress(t *testing.T) {
	// Test with positive total
	p := NewProgress(100)
	cur, total, _ := p.Value()
	assert.Equal(t, 0.0, cur)
	assert.Equal(t, 100.0, total)

	// Test with float total
	p = NewProgress(50.5)
	cur, total, _ = p.Value()
	assert.Equal(t, 0.0, cur)
	assert.InDelta(t, 50.5, total, 0.001)

	// Test panic with zero total
	assert.Panics(t, func() {
		NewProgress(0)
	})

	// Test panic with negative total
	assert.Panics(t, func() {
		NewProgress(-10)
	})
}

func TestProgress_Add(t *testing.T) {
	p := NewProgress(100)

	// Test adding positive amount
	p.Add(25)
	cur, total, _ := p.Value()
	assert.Equal(t, 0.0, cur)
	assert.Equal(t, 125.0, total)

	// Test adding more items
	p.Add(30)
	cur, total, _ = p.Value()
	assert.Equal(t, 0.0, cur)
	assert.Equal(t, 155.0, total)

	// Test adding with float
	p.Add(2.5)
	cur, total, _ = p.Value()
	assert.Equal(t, 0.0, cur)
	assert.InDelta(t, 157.5, total, 0.001)

	// Test that percentage changes when total changes
	p2 := NewProgress(100)
	p2.Complete(50) // 50% complete
	assert.Equal(t, 50.0, p2.Percent())

	p2.Add(100)                         // Total becomes 200
	assert.Equal(t, 25.0, p2.Percent()) // 50/200 = 25%

	// Test panic with zero
	assert.Panics(t, func() {
		p.Add(0)
	})

	// Test panic with negative
	assert.Panics(t, func() {
		p.Add(-5)
	})
}

func TestProgress_Complete(t *testing.T) {
	p := NewProgress(100)

	// Test completing positive amount
	p.Complete(25)
	cur, _, _ := p.Value()
	assert.Equal(t, 25.0, cur)

	// Test completing more items
	p.Complete(30)
	cur, _, _ = p.Value()
	assert.Equal(t, 55.0, cur)

	// Test completing with float
	p.Complete(2.5)
	cur, _, _ = p.Value()
	assert.InDelta(t, 57.5, cur, 0.001)

	// Test panic with zero
	assert.Panics(t, func() {
		p.Complete(0)
	})

	// Test panic with negative
	assert.Panics(t, func() {
		p.Complete(-5)
	})
}

func TestProgress_Percent(t *testing.T) {
	p := NewProgress(100)

	// Test 0% completion
	assert.Equal(t, 0.0, p.Percent(), 1e-10)

	// Test 25% completion
	p.Complete(25)
	assert.InDelta(t, 25.0, p.Percent(), 1e-10)

	// Test 50% completion
	p.Complete(25)
	assert.InDelta(t, 50.0, p.Percent(), 1e-10)

	// Test 100% completion
	p.Complete(50)
	assert.InDelta(t, 100.0, p.Percent(), 1e-10)

	// Test over 100% completion
	p.Complete(10)
	assert.InDelta(t, 110.0, p.Percent(), 1e-10)

	// Test with float values
	p2 := NewProgress(33.3)
	p2.Complete(11.1)
	assert.InDelta(t, 33.33, p2.Percent(), 0.1)
}

func TestProgress_EstimatedTime(t *testing.T) {
	p := NewProgress(100)

	// Test with no progress - should return far future
	estimated := p.EstimatedTime()
	assert.True(t, estimated.After(time.Now().AddDate(998, 0, 0)))

	// Test with some progress
	p.Complete(25)
	time.Sleep(10 * time.Millisecond) // Small delay to simulate work
	estimated = p.EstimatedTime()
	assert.True(t, estimated.After(time.Now()))
}

func TestProgress_RemainTime(t *testing.T) {
	p := NewProgress(100)

	// Test with no progress
	assert.Equal(t, time.Duration(0), p.RemainTime())

	// Test with some progress
	p.Complete(25)
	time.Sleep(10 * time.Millisecond) // Small delay to simulate work
	remainTime := p.RemainTime()
	assert.True(t, remainTime > 0)

	// Test with 100% completion
	p.Complete(75)
	remainTime = p.RemainTime()
	assert.Equal(t, time.Duration(0), remainTime)
}

func TestProgress_Value(t *testing.T) {
	start := time.Now()
	p := NewProgress(100)

	// Test initial values
	cur, total, updated := p.Value()
	assert.Equal(t, 0.0, cur)
	assert.Equal(t, 100.0, total)
	assert.True(t, updated.After(start) || updated.Equal(start))

	// Test after completion
	beforeComplete := time.Now()
	p.Complete(50)
	cur, total, updated = p.Value()
	assert.Equal(t, 50.0, cur)
	assert.Equal(t, 100.0, total)
	assert.True(t, updated.After(beforeComplete) || updated.Equal(beforeComplete))
}

func TestProgress_ConcurrentOperations(t *testing.T) {
	p := NewProgress(1000)
	numGoroutines := 50
	completionPerGoroutine := 2.0

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Concurrent Complete operations
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			p.Complete(completionPerGoroutine)
		}()
	}

	wg.Wait()

	// Check final value
	cur, total, _ := p.Value()
	expected := float64(numGoroutines) * completionPerGoroutine
	assert.Equal(t, expected, cur)
	assert.Equal(t, 1000.0, total)
}

func TestProgress_ConcurrentReadWrite(t *testing.T) {
	p := NewProgress(1000)
	done := make(chan bool)

	// Reader goroutines
	for i := 0; i < 10; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					p.Value()
					p.Percent()
					p.EstimatedTime()
					p.RemainTime()
				}
			}
		}()
	}

	// Writer goroutines
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				p.Complete(1)
				time.Sleep(time.Microsecond)
			}
		}()
	}

	wg.Wait()
	close(done)

	// Verify final state
	cur, _, _ := p.Value()
	assert.Equal(t, 100.0, cur)
}

func TestProgress_RealisticScenario(t *testing.T) {
	// Simulate file download progress
	totalBytes := 1024.0
	p := NewProgress(totalBytes)

	// Initial state
	assert.Equal(t, 0.0, p.Percent())

	// Download in chunks
	chunkSizes := []float64{128, 256, 128, 256, 128, 128}
	expectedProgress := 0.0

	for _, chunk := range chunkSizes {
		p.Complete(chunk)
		expectedProgress += chunk

		cur, total, _ := p.Value()
		assert.Equal(t, expectedProgress, cur)
		assert.Equal(t, totalBytes, total)

		expectedPercent := (expectedProgress / totalBytes) * 100
		assert.InDelta(t, expectedPercent, p.Percent(), 0.001)

		if expectedProgress > 0 {
			assert.True(t, p.RemainTime() >= 0)
		}
	}

	// Final check
	assert.Equal(t, 100.0, p.Percent())
}

func TestProgress_EdgeCases(t *testing.T) {
	// Test very large total
	p := NewProgress(1e15)
	p.Complete(1e14)
	assert.InDelta(t, 10.0, p.Percent(), 0.001)

	// Test very small total
	p2 := NewProgress(1e-5)
	p2.Complete(5e-6)
	assert.InDelta(t, 50.0, p2.Percent(), 0.001)

	// Test completion beyond total
	p3 := NewProgress(100)
	p3.Complete(150)
	assert.Equal(t, 150.0, p3.Percent())

	cur, total, _ := p3.Value()
	assert.Equal(t, 150.0, cur)
	assert.Equal(t, 100.0, total)
}

func TestProgress_TimeCalculations(t *testing.T) {
	p := NewProgress(100)

	// Test that estimated time changes as progress is made
	p.Complete(10)
	time.Sleep(25 * time.Millisecond)

	firstEstimate := p.EstimatedTime()
	firstRemain := p.RemainTime()

	p.Complete(10)
	time.Sleep(25 * time.Millisecond)

	secondEstimate := p.EstimatedTime()
	secondRemain := p.RemainTime()

	// More progress should result in less remaining time
	assert.True(t, secondRemain < firstRemain)

	// Estimates should be in the future
	assert.True(t, firstEstimate.After(time.Now()))
	assert.True(t, secondEstimate.After(time.Now()))
}

func TestProgress_ZeroRemainTime(t *testing.T) {
	p := NewProgress(50)

	// Complete all work
	p.Complete(50)

	// Remaining time should be 0
	assert.Equal(t, time.Duration(0), p.RemainTime())

	// Percent should be 100%
	assert.Equal(t, 100.0, p.Percent())
}
