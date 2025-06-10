// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

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

// Uint32 returns, as an uint32, a uniform integer
// from the default [Reader].
// Available since v0.3.0
func Uint32() uint32 {
	v, err := rand.Int(rand.Reader, new(big.Int).SetUint64(uint64(uint32Max)))
	if err != nil {
		panic(err)
	}
	return uint32(v.Uint64())
}

// Uint32n returns, as an uint32, a uniform number in the half-open interval [0,n)
// from the default [Reader].
// It panics if n <= 0.
// Available since v0.3.0
func Uint32n(n uint32) uint32 {
	if n <= 0 {
		panic("invalid argument to Uint32n")
	}
	v, err := rand.Int(rand.Reader, new(big.Int).SetUint64(uint64(n)))
	if err != nil {
		panic(err)
	}
	return uint32(v.Uint64())
}

// Uint32r returns, as an uint32, a uniform number in the half-open interval [i,n)
// from the default [Reader].
// It panics if n <= 0 or n <= i.
// Available since v0.3.0
func Uint32r(i, n uint32) uint32 {
	if n <= 0 || n <= i {
		panic("invalid argument to Uint32r")
	}
	return i + Uint32n(n-i)
}

// Uint64 returns, as an uint64, a uniform integer
// from the default [Reader].
// Available since v0.3.0
func Uint64() uint64 {
	v, err := rand.Int(rand.Reader, new(big.Int).SetUint64(uint64Max))
	if err != nil {
		panic(err)
	}
	return v.Uint64()
}

// Uint64n returns, as an uint64, a uniform number in the half-open interval [0,n)
// from the default [Reader].
// It panics if n <= 0.
// Available since v0.3.0
func Uint64n(n uint64) uint64 {
	if n <= 0 {
		panic("invalid argument to Uint64n")
	}
	v, err := rand.Int(rand.Reader, new(big.Int).SetUint64(n))
	if err != nil {
		panic(err)
	}
	return v.Uint64()
}

// Uint64r returns, as an uint64, a uniform number in the half-open interval [i,n)
// from the default [Reader].
// It panics if n <= 0 or n <= i.
// Available since v0.3.0
func Uint64r(i, n uint64) uint64 {
	if n <= 0 || n <= i {
		panic("invalid argument to Uint64r")
	}
	return i + Uint64n(n-i)
}
