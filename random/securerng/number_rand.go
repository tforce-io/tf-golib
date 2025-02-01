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

package securerng

import (
	"crypto/rand"
	"math/big"
)

// See https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go#comment75694539_6878590
// and https://go.dev/doc/effective_go#printing for reference.
var (
	uintMax   = ^uint(0)
	uint32Max = ^uint32(0)
	uint64Max = ^uint64(0)
	intMax    = int(uintMax >> 1)
	int32Max  = int32(uint32Max >> 1)
	int64Max  = int64(uint64Max >> 1)
)

// Int returns, as an int, a non-negative uniform integer
// from the default [Reader].
// Available since v0.3.0
func Int() int {
	v, err := rand.Int(rand.Reader, big.NewInt(int64(intMax)))
	if err != nil {
		panic(err)
	}
	return int(v.Int64())
}

// Intn returns, as an int, a non-negative uniform number in the half-open interval [0,n)
// from the default [Reader].
// It panics if n <= 0.
// Available since v0.3.0
func Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	v, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		panic(err)
	}
	return int(v.Int64())
}

// Intr returns, as an int, a non-negative uniform number in the half-open interval [i,n)
// from the default [Reader].
// It panics if n <= 0 or n <= i.
// Available since v0.3.0
func Intr(i, n int) int {
	if n <= 0 || n <= i {
		panic("invalid argument to Intr")
	}
	return i + Intn(n-i)
}

// Int31 returns, as an int32, a non-negative uniform integer
// from the default [Reader].
// Available since v0.3.0
func Int31() int32 {
	v, err := rand.Int(rand.Reader, big.NewInt(int64(int32Max)))
	if err != nil {
		panic(err)
	}
	return int32(v.Int64())
}

// Int31n returns, as an int32, a non-negative uniform number in the half-open interval [0,n)
// from the default [Reader].
// It panics if n <= 0.
// Available since v0.3.0
func Int31n(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}
	v, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		panic(err)
	}
	return int32(v.Int64())
}

// Int31r returns, as an int32, a non-negative uniform number in the half-open interval [i,n)
// from the default [Reader].
// It panics if n <= 0 or n <= i.
// Available since v0.3.0
func Int31r(i, n int32) int32 {
	if n <= 0 || n <= i {
		panic("invalid argument to Int31r")
	}
	return i + Int31n(n-i)
}

// Int63 returns, as an int64, a non-negative uniform integer
// from the default [Reader].
// Available since v0.3.0
func Int63() int64 {
	v, err := rand.Int(rand.Reader, big.NewInt(int64Max))
	if err != nil {
		panic(err)
	}
	return v.Int64()
}

// Int63n returns, as an int64, a non-negative uniform number in the half-open interval [0,n)
// from the default [Reader].
// It panics if n <= 0.
// Available since v0.3.0
func Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	v, err := rand.Int(rand.Reader, big.NewInt(n))
	if err != nil {
		panic(err)
	}
	return v.Int64()
}

// Int63r returns, as an int64, a non-negative uniform number in the half-open interval [i,n)
// from the default [Reader].
// It panics if n <= 0 or n <= i.
// Available since v0.3.0
func Int63r(i, n int64) int64 {
	if n <= 0 || n <= i {
		panic("invalid argument to Int63r")
	}
	return i + Int63n(n-i)
}
