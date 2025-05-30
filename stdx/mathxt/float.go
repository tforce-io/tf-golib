package mathxt

/* --- FLOAT32 --- */

// Returns the smaller of x or y for float32.
//
// Available since vTBD
func MinFloat32(x float32, y ...float32) float32 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Returns the larger of x or y for float32.
//
// Available since vTBD
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

// Returns the smaller of x or y for float64.
//
// Available since vTBD
func MinFloat64(x float64, y ...float64) float64 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Returns the larger of x or y for float64.
//
// Available since vTBD
func MaxFloat64(x float64, y ...float64) float64 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}
