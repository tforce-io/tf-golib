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

// ProcessState indicates status of service's Process routine.
//
// Available since vTBD
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
// Available since vTBD
type HookState struct {
	Handled bool
}
