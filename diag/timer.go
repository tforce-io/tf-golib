// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package diag

import "time"

// Timer is a helper type to time functions.
//
// Available since v0.7.0
type Timer struct {
	created time.Time
}

// Return a new Timer starting from now.
//
// Available since v0.7.0
func NewTimer() *Timer {
	return &Timer{
		created: time.Now(),
	}
}

// Return the duration since the Timer was created.
//
// Available since v0.7.0
func (t *Timer) Duration() time.Duration {
	return time.Since(t.created)
}
