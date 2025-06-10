// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package mathxt

/* --- INT --- */

// Return the absolute value of x for int.
//
// Available since v0.6.0
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Return the smaller of x or y for int.
//
// Available since v0.6.0
func MinInt(x int, y ...int) int {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Return the larger of x or y for int.
//
// Available since v0.6.0
func MaxInt(x int, y ...int) int {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}

/* --- INT8 --- */

// Return the absolute value of x for int8.
//
// Available since v0.6.0
func AbsInt8(x int8) int8 {
	if x < 0 {
		return -x
	}
	return x
}

// Return the smaller of x or y for int8.
//
// Available since v0.6.0
func MinInt8(x int8, y ...int8) int8 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Return the larger of x or y for int8.
//
// Available since v0.6.0
func MaxInt8(x int8, y ...int8) int8 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}

/* --- INT16 --- */

// Return the absolute value of x for int16.
//
// Available since v0.6.0
func AbsInt16(x int16) int16 {
	if x < 0 {
		return -x
	}
	return x
}

// Return the smaller of x or y for int16.
//
// Available since v0.6.0
func MinInt16(x int16, y ...int16) int16 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Return the larger of x or y for int16.
//
// Available since v0.6.0
func MaxInt16(x int16, y ...int16) int16 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}

/* --- INT32 --- */

// Return the absolute value of x for int32.
//
// Available since v0.6.0
func AbsInt32(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// Return the smaller of x or y for int32.
//
// Available since v0.6.0
func MinInt32(x int32, y ...int32) int32 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Return the larger of x or y for int32.
//
// Available since v0.6.0
func MaxInt32(x int32, y ...int32) int32 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}

/* --- INT64 --- */

// Return the absolute value of x for int64.
//
// Available since v0.6.0
func AbsInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// Return the smaller of x or y for int64.
//
// Available since v0.6.0
func MinInt64(x int64, y ...int64) int64 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Return the larger of x or y for int64.
//
// Available since v0.6.0
func MaxInt64(x int64, y ...int64) int64 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}
