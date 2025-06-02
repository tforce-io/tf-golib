package mathxt

/* --- INT --- */

// Return the absolute value of x.
//
// Available since vTBD
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Returns the smaller of x or y for int.
//
// Available since vTBD
func MinInt(x int, y ...int) int {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Returns the larger of x or y for int.
//
// Available since vTBD
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

// Return the absolute value of x.
//
// Available since vTBD
func AbsInt8(x int8) int8 {
	if x < 0 {
		return -x
	}
	return x
}

// Returns the smaller of x or y for int8.
//
// Available since vTBD
func MinInt8(x int8, y ...int8) int8 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Returns the larger of x or y for int8.
//
// Available since vTBD
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

// Return the absolute value of x.
//
// Available since vTBD
func AbsInt16(x int16) int16 {
	if x < 0 {
		return -x
	}
	return x
}

// Returns the smaller of x or y for int16.
//
// Available since vTBD
func MinInt16(x int16, y ...int16) int16 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Returns the larger of x or y for int16.
//
// Available since vTBD
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

// Return the absolute value of x.
//
// Available since vTBD
func AbsInt32(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// Returns the smaller of x or y for int32.
//
// Available since vTBD
func MinInt32(x int32, y ...int32) int32 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Returns the larger of x or y for int32.
//
// Available since vTBD
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

// Return the absolute value of x.
//
// Available since vTBD
func AbsInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// Returns the smaller of x or y for int64.
//
// Available since vTBD
func MinInt64(x int64, y ...int64) int64 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Returns the larger of x or y for int64.
//
// Available since vTBD
func MaxInt64(x int64, y ...int64) int64 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}
