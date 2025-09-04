// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package diag

import "time"

// Progress is a helper type to track progress.
//
// Available since v0.7.0
type Progress struct {
	curVal *Counter
	ttlVal *Counter

	started time.Time
	updated time.Time
}

// Return new Progress tracker with total items to complete.
//
// Available since v0.7.0
func NewProgress(total float64) *Progress {
	if total <= 0 {
		panic("total must be positive")
	}
	return &Progress{
		curVal:  NewCounter(0),
		ttlVal:  NewCounter(total),
		started: time.Now(),
		updated: time.Now(),
	}
}

// Add v number of items as to be completed.
//
// Available since v0.7.0
func (p *Progress) Add(v float64) {
	if v <= 0 {
		panic("v must be positive")
	}
	p.ttlVal.Add(v)
}

// Mark v number of items as completed.
//
// Available since v0.7.0
func (p *Progress) Complete(v float64) {
	if v <= 0 {
		panic("v must be positive")
	}
	p.curVal.Add(v)
	p.updated = time.Now()
}

// Estimate when the Progress will be completed.
//
// Available since v0.7.0
func (p *Progress) EstimatedTime() time.Time {
	if p.curVal.Value() == 0 {
		return time.Now().AddDate(999, 0, 0)
	}
	passed := time.Since(p.started)
	timeToComplete := passed / time.Duration(p.curVal.Value()) * time.Duration(p.ttlVal.Value())
	return p.updated.Add(timeToComplete)
}

func (p *Progress) Percent() float64 {
	return (p.curVal.Value() / p.ttlVal.Value()) * 100
}

// Estimate how long the Progress will be completed.
//
// Available since v0.7.0
func (p *Progress) RemainTime() time.Duration {
	if p.curVal.Value() == 0 {
		return 0
	}
	passed := time.Since(p.started)
	remain := p.ttlVal.Value() - p.curVal.Value()
	remainTime := passed / time.Duration(p.curVal.Value()) * time.Duration(remain)
	return remainTime
}

// Return current value, total value and last updated time.
//
// Available since v0.7.0
func (p *Progress) Value() (float64, float64, time.Time) {
	return p.curVal.Value(), p.ttlVal.Value(), p.updated
}
