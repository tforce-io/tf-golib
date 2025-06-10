// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package multiplex

// ProcessState indicates status of service's Process routine.
//
// Available since v0.5.0
type ProcessState int8

const (
	// Process routine started for the first time.
	InitState ProcessState = iota

	// Process routine wil be terminated.
	ExitState

	// The request is processed successfully.
	SuccessState

	// The request is processed failed.
	ErrorState

	// The request is processed failed and need to be retried.
	RetryState
)

// ProcessState indicates status of service's Hook processing.
//
// Available since v0.5.0
type HookState struct {
	Handled bool
}
