// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package bigxt

import "math/big"

// Compare that two big.Int a == b
//
// Available since v0.10.0
func AreEqualInt(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}

// Compare that two big.Int a > b
//
// Available since v0.10.0
func IsGreaterThanInt(a, b *big.Int) bool {
	return a.Cmp(b) > 0
}

// Compare that two big.Int a >= b
//
// Available since v0.10.0
func IsGreaterThanOrEqualInt(a, b *big.Int) bool {
	return a.Cmp(b) >= 0
}

// Compare that two big.Int a < b
//
// Available since v0.10.0
func IsLessThanInt(a, b *big.Int) bool {
	return a.Cmp(b) < 0
}

// Compare that two big.Int a <= b
//
// Available since v0.10.0
func IsLessThanOrEqualInt(a, b *big.Int) bool {
	return a.Cmp(b) <= 0
}
