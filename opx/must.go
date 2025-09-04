// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package opx

// Ensure that the err is nil, otherwise panic.
// Available since v0.7.0
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Ensure that the err is nil then return all preceding variables, otherwise panic.
// Used to eliminate intermediate error for functions.
// Available since v0.7.0
func Must1[T1 any](t1 T1, err error) T1 {
	if err != nil {
		panic(err)
	}
	return t1
}

// Ensure that the err is nil then return all preceding variables, otherwise panic.
// Used to eliminate intermediate error for functions.
// Available since v0.7.0
func Must2[T1 any, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return t1, t2
}

// Ensure that the err is nil then return all preceding variables, otherwise panic.
// Used to eliminate intermediate error for functions.
// Available since v0.7.0
func Must3[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3
}

// Ensure that the err is nil then return all preceding variables, otherwise panic.
// Used to eliminate intermediate error for functions.
// Available since v0.7.0
func Must4[T1 any, T2 any, T3 any, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) (T1, T2, T3, T4) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3, t4
}

// Ensure that the err is nil then return all preceding variables, otherwise panic.
// Used to eliminate intermediate error for functions.
// Available since v0.7.0
func Must5[T1 any, T2 any, T3 any, T4 any, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) (T1, T2, T3, T4, T5) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3, t4, t5
}
