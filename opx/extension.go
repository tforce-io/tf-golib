// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package opx

// Return the first non-nil value in a array. Only support pointers of the same type.
// Available since v0.3.0
func Coalesce[T any](values ...*T) *T {
	for _, value := range values {
		if value != nil {
			return value
		}
	}
	return nil
}

// Method version of ternary assignmet. If cond is true, returns x, otherwise returns y.
// Available since v0.3.0
func Ternary[T any](cond bool, x, y T) T {
	if cond {
		return x
	}
	return y
}
