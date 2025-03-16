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

package slicext

// Compare two slices x and y of comparable types that is deep equals.
// Available since v0.3.0
func AreEqual[S ~[]T, T comparable](x, y S) bool {
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
func AreEqualFunc[S ~[]T, T any](x, y S, equalFunc func(x, y T) bool) bool {
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

// Check whether slice s is Nil nor zero in length.
// Available since v0.3.0
func IsEmpty[S ~[]T, T any](s S) bool {
	if s == nil {
		return true
	}
	return len(s) == 0
}

// Map a slice contains items of type T to a new slice contains equivalent number items of type N.
// Nil handling and special conversion will be taken care of in mapFunc.
// Available since v0.4.0
func Map[S ~[]T, M ~[]N, T, N any](s S, mapFunc func(t T) N) M {
	m := make([]N, len(s))
	for i, t := range s {
		m[i] = mapFunc(t)
	}
	return m
}

// Map a slice contains items of type T to a Key-Value map.
// keyFunc must returns extra bool to indicate whether slice item T will be included in the resulted map.
// Available since v0.4.0
func MapKV[S ~[]T, K comparable, V, T any](s S, keyFunc func(t T) (K, bool), valFunc func(t T) V) map[K]V {
	m := make(map[K]V)
	for _, t := range s {
		if k, ok := keyFunc(t); ok {
			m[k] = valFunc(t)
		}
	}
	return m
}
