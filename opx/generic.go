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

// Compare two slices x and y of comparable types that is deep equals.
// Available since v0.3.0
func AreEqualSlices[S ~[]T, T comparable](x, y S) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// Compare two slices x and y of any type that is deep equals.
// Available since v0.3.0
func AreEqualSlicesFunc[S ~[]T, T any](x, y S, equalFunc func(x, y T) bool) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if !equalFunc(x[i], y[i]) {
			return false
		}
	}
	return true
}

// Check whether a value v is a member of slice s.
// Available since v0.3.0
func Contains[S ~[]T, T comparable](s S, v T) bool {
	for i := range s {
		if v == s[i] {
			return true
		}
	}
	return false
}

// Check whether a value v is a member of slice s.
// Available since v0.3.0
func ContainsFunc[S ~[]T, T comparable](s S, v T, equalFunc func(x, y T) bool) bool {
	for i := range s {
		if equalFunc(v, s[i]) {
			return true
		}
	}
	return false
}

// Check whether is slice is Nil nor zero in length.
// Available since v0.3.0
func IsEmptySlice[S ~[]T, T any](slice S) bool {
	if slice == nil {
		return true
	}
	return len(slice) == 0
}
