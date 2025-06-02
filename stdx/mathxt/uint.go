package mathxt

/* --- UINT --- */

// Return the smaller of x or y for uint.
//
// Available since v0.6.0
func MinUint(x uint, y ...uint) uint {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Return the larger of x or y for uint.
//
// Available since v0.6.0
func MaxUint(x uint, y ...uint) uint {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}

/* --- UINT8 --- */

// Return the smaller of x or y for uint8.
//
// Available since v0.6.0
func MinUint8(x uint8, y ...uint8) uint8 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Return the larger of x or y for uint8.
//
// Available since v0.6.0
func MaxUint8(x uint8, y ...uint8) uint8 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}

/* --- UINT16 --- */

// Return the smaller of x or y for uint16.
//
// Available since v0.6.0
func MinUint16(x uint16, y ...uint16) uint16 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Return the larger of x or y for uint16.
//
// Available since v0.6.0
func MaxUint16(x uint16, y ...uint16) uint16 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}

/* --- UINT32 --- */

// Return the smaller of x or y for uint32.
//
// Available since v0.6.0
func MinUint32(x uint32, y ...uint32) uint32 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Return the larger of x or y for uint32.
//
// Available since v0.6.0
func MaxUint32(x uint32, y ...uint32) uint32 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}

/* --- UINT64 --- */

// Return the smaller of x or y for uint64.
//
// Available since v0.6.0
func MinUint64(x uint64, y ...uint64) uint64 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

// Return the larger of x or y for uint64.
//
// Available since v0.6.0
func MaxUint64(x uint64, y ...uint64) uint64 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}
