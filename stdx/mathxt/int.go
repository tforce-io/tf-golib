package mathxt

/* --- INT --- */

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
