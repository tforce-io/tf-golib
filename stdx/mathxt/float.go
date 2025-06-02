package mathxt

import "math"

/* --- FLOAT32 --- */

// Return the absolute value of x for float32.
//
// Available since v0.6.0
func AbsFloat32(x float32) float32 {
	return float32(math.Abs(float64(x)))
}

// Return the smaller of x or y for float32.
//
// Available since v0.6.0
func MinFloat32(x float32, y ...float32) float32 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Return the larger of x or y for float32.
//
// Available since v0.6.0
func MaxFloat32(x float32, y ...float32) float32 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}

/* --- FLOAT64 --- */

// Return the absolute value of x for float64.
//
// Available since v0.6.0
func AbsFloat64(x float64) float64 {
	return math.Abs(x)
}

// Return the smaller of x or y for float64.
//
// Available since v0.6.0
func MinFloat64(x float64, y ...float64) float64 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Return the larger of x or y for float64.
//
// Available since v0.6.0
func MaxFloat64(x float64, y ...float64) float64 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}
