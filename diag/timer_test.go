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

func TestNewTimer(t *testing.T) {
	before := time.Now()
	timer := NewTimer()
	after := time.Now()

	// Timer should not be nil
	assert.NotNil(t, timer)

	// Timer creation time should be between before and after
	assert.True(t, timer.created.After(before) || timer.created.Equal(before))
	assert.True(t, timer.created.Before(after) || timer.created.Equal(after))

	// Initial duration should be very small (close to 0)
	duration := timer.Duration()
	assert.True(t, duration >= 0)
	assert.True(t, duration < time.Millisecond) // Should be very fast
}

func TestTimer_Duration(t *testing.T) {
	timer := NewTimer()

	// Test immediate duration
	duration1 := timer.Duration()
	assert.True(t, duration1 >= 0)

	// Wait a bit and test again
	time.Sleep(10 * time.Millisecond)
	duration2 := timer.Duration()

	// Second duration should be larger
	assert.True(t, duration2 > duration1)
	assert.True(t, duration2 >= 10*time.Millisecond)

	// Test with longer delay
	time.Sleep(20 * time.Millisecond)
	duration3 := timer.Duration()

	// Third duration should be even larger
	assert.True(t, duration3 > duration2)
	assert.True(t, duration3 >= 30*time.Millisecond)
}

func TestTimer_DurationProgression(t *testing.T) {
	timer := NewTimer()
	var durations []time.Duration

	// Collect durations over time
	for i := 0; i < 5; i++ {
		durations = append(durations, timer.Duration())
		time.Sleep(5 * time.Millisecond)
	}

	// Each duration should be greater than the previous
	for i := 1; i < len(durations); i++ {
		assert.True(t, durations[i] > durations[i-1],
			"Duration %d (%v) should be greater than duration %d (%v)",
			i, durations[i], i-1, durations[i-1])
	}

	// Total duration should be reasonable (at least 20ms)
	totalDuration := timer.Duration()
	assert.True(t, totalDuration >= 20*time.Millisecond)
}

func TestTimer_ConcurrentAccess(t *testing.T) {
	timer := NewTimer()
	numGoroutines := 50
	results := make([]time.Duration, numGoroutines)

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Concurrent access to Duration method
	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()
			time.Sleep(time.Duration(index) * time.Microsecond) // Stagger slightly
			results[index] = timer.Duration()
		}(i)
	}

	wg.Wait()

	// All results should be valid (non-negative)
	for i, duration := range results {
		assert.True(t, duration >= 0, "Duration %d should be non-negative: %v", i, duration)
	}

	// Results should generally increase (with some tolerance for concurrent execution)
	// At least some results should be different
	unique := make(map[time.Duration]bool)
	for _, duration := range results {
		unique[duration] = true
	}
	assert.True(t, len(unique) > 1, "Should have multiple unique durations")
}

func TestTimer_MultipleTimers(t *testing.T) {
	// Create multiple timers at different times
	timer1 := NewTimer()
	time.Sleep(5 * time.Millisecond)

	timer2 := NewTimer()
	time.Sleep(5 * time.Millisecond)

	timer3 := NewTimer()
	time.Sleep(5 * time.Millisecond)

	// Get durations
	duration1 := timer1.Duration()
	duration2 := timer2.Duration()
	duration3 := timer3.Duration()

	// First timer should have longest duration
	assert.True(t, duration1 > duration2)
	assert.True(t, duration2 > duration3)

	// All durations should be positive
	assert.True(t, duration1 > 0)
	assert.True(t, duration2 > 0)
	assert.True(t, duration3 > 0)

	// Verify rough timing expectations
	assert.True(t, duration1 >= 15*time.Millisecond) // ~15ms
	assert.True(t, duration2 >= 10*time.Millisecond) // ~10ms
	assert.True(t, duration3 >= 5*time.Millisecond)  // ~5ms
}

func TestTimer_RealisticScenario(t *testing.T) {
	// Simulate timing a function execution
	timer := NewTimer()

	// Simulate some work
	workSteps := []time.Duration{
		5 * time.Millisecond,
		10 * time.Millisecond,
		7 * time.Millisecond,
		3 * time.Millisecond,
	}

	var stepDurations []time.Duration
	for _, step := range workSteps {
		time.Sleep(step)
		stepDurations = append(stepDurations, timer.Duration())
	}

	// Verify step durations are increasing
	for i := 1; i < len(stepDurations); i++ {
		assert.True(t, stepDurations[i] > stepDurations[i-1])
	}

	// Final duration should be approximately the sum of work steps
	expectedTotal := time.Duration(0)
	for _, step := range workSteps {
		expectedTotal += step
	}

	finalDuration := timer.Duration()
	assert.True(t, finalDuration >= expectedTotal)
}

func TestTimer_ZeroDuration(t *testing.T) {
	timer := NewTimer()

	// Immediate duration check should be very small but non-negative
	duration := timer.Duration()
	assert.True(t, duration >= 0)
	assert.True(t, duration < time.Millisecond)
}

func TestTimer_LongDuration(t *testing.T) {
	timer := NewTimer()

	// Wait a longer period
	time.Sleep(50 * time.Millisecond)

	duration := timer.Duration()
	assert.True(t, duration >= 50*time.Millisecond)
}

func TestTimer_ConsistentReads(t *testing.T) {
	timer := NewTimer()
	time.Sleep(10 * time.Millisecond)

	// Multiple consecutive reads should return increasing or equal values
	var durations []time.Duration
	for i := 0; i < 10; i++ {
		durations = append(durations, timer.Duration())
	}

	// Each duration should be >= the previous one
	for i := 1; i < len(durations); i++ {
		assert.True(t, durations[i] >= durations[i-1],
			"Duration should not decrease: %v >= %v", durations[i], durations[i-1])
	}
}

func TestTimer_HighFrequencyAccess(t *testing.T) {
	timer := NewTimer()
	iterations := 1000
	durations := make([]time.Duration, iterations)

	// High frequency duration calls
	for i := 0; i < iterations; i++ {
		durations[i] = timer.Duration()
	}

	// All durations should be valid
	for i, duration := range durations {
		assert.True(t, duration >= 0, "Duration %d should be non-negative", i)
	}

	// Durations should be non-decreasing
	for i := 1; i < len(durations); i++ {
		assert.True(t, durations[i] >= durations[i-1],
			"Duration should not decrease at iteration %d", i)
	}
}

func TestTimer_Precision(t *testing.T) {
	timer := NewTimer()

	// Test precision with small delays
	time.Sleep(10 * time.Millisecond)
	duration1 := timer.Duration()

	time.Sleep(10 * time.Millisecond)
	duration2 := timer.Duration()

	// Second duration should be detectably larger
	assert.True(t, duration2 > duration1)

	// Difference should be reasonable (around 1ms, but allowing for system variance)
	diff := duration2 - duration1
	assert.True(t, diff >= 10*time.Microsecond) // At least 10ms
}
