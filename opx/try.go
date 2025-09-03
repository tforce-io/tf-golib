// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package opx

import "errors"

// Ensure that the ok is true, otherwise panic.
// Available since v0.7.0
func Try(ok bool) {
	if !ok {
		panic(errors.New("tried and ok is false"))
	}
}

// Ensure that the ok is true then return all preceding variables, otherwise panic.
// Used to eliminate ok check for functions.
// Available since v0.7.0
func Try1[T1 any](t1 T1, ok bool) T1 {
	if !ok {
		panic(errors.New("tried and ok is false"))
	}
	return t1
}

// Ensure that the ok is true then return all preceding variables, otherwise panic.
// Used to eliminate ok check for functions.
// Available since v0.7.0
func Try2[T1 any, T2 any](t1 T1, t2 T2, ok bool) (T1, T2) {
	if !ok {
		panic(errors.New("tried and ok is false"))
	}
	return t1, t2
}

// Ensure that the ok is true then return all preceding variables, otherwise panic.
// Used to eliminate ok check for functions.
// Available since v0.7.0
func Try3[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3, ok bool) (T1, T2, T3) {
	if !ok {
		panic(errors.New("tried and ok is false"))
	}
	return t1, t2, t3
}

// Ensure that the ok is true then return all preceding variables, otherwise panic.
// Used to eliminate ok check for functions.
// Available since v0.7.0
func Try4[T1 any, T2 any, T3 any, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, ok bool) (T1, T2, T3, T4) {
	if !ok {
		panic(errors.New("tried and ok is false"))
	}
	return t1, t2, t3, t4
}

// Ensure that the ok is true then return all preceding variables, otherwise panic.
// Used to eliminate ok check for functions.
// Available since v0.7.0
func Try5[T1 any, T2 any, T3 any, T4 any, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, ok bool) (T1, T2, T3, T4, T5) {
	if !ok {
		panic(errors.New("tried and ok is false"))
	}
	return t1, t2, t3, t4, t5
}
