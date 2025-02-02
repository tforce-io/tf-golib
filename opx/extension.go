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

package opx

// Returns the first non-nil value in a array. Only support pointers of the same type.
// Available since v0.3.0
func Coalesce[T any](values ...*T) *T {
	for _, value := range values {
		if value != nil {
			return value
		}
	}
	return nil
}

// Method version of ternary assignmet. If cond is true, returns x, otherwise returns y.
// Available since v0.3.0
func Ternary[T any](cond bool, x, y T) T {
	if cond {
		return x
	}
	return y
}
