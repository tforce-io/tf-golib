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

package pseudorng

import "math/rand"

// Int returns a non-negative pseudo-random int from the default [Source].
// Available since v0.3.0
func Int() int {
	return rand.Int()
}

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n)
// from the default [Source].
// It panics if n <= 0.
// Available since v0.3.0
func Intn(n int) int {
	return rand.Intn(n)
}

// Intr returns, as an int, a non-negative pseudo-random number in the half-open interval [i,n)
// from the default [Source].
// It panics if n <= 0 or n <= i.
// Available since v0.3.0
func Intr(i, n int) int {
	if n <= 0 || n <= i {
		panic("invalid argument to Intr")
	}
	return i + rand.Intn(n-i)
}

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32
// from the default [Source].
// Available since v0.3.0
func Int31() int32 {
	return rand.Int31()
}

// Int31n returns, as an int32, a non-negative pseudo-random number in the half-open interval [0,n)
// from the default [Source].
// It panics if n <= 0.
// Available since v0.3.0
func Int31n(n int32) int32 {
	return rand.Int31n(n)
}

// Int31r returns, as an int32, a non-negative pseudo-random number in the half-open interval [i,n)
// from the default [Source].
// It panics if n <= 0 or n <= i.
// Available since v0.3.0
func Int31r(i, n int32) int32 {
	if n <= 0 || n <= i {
		panic("invalid argument to Int31r")
	}
	return i + rand.Int31n(n-i)
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64
// from the default [Source].
// Available since v0.3.0
func Int63() int64 {
	return rand.Int63()
}

// Int63n returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n)
// from the default [Source].
// It panics if n <= 0.
// Available since v0.3.0
func Int63n(n int64) int64 {
	return rand.Int63n(n)
}

// Int63r returns, as an int64, a non-negative pseudo-random number in the half-open interval [i,n)
// from the default [Source].
// It panics if n <= 0 or n <= i.
// Available since v0.3.0
func Int63r(i, n int64) int64 {
	if n <= 0 || n <= i {
		panic("invalid argument to Int63r")
	}
	return i + rand.Int63n(n-i)
}
