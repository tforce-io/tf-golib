// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package stringxt

// Check whether a string has zero length.
// Available since v0.3.0
func IsEmpty(str string) bool {
	return str == ""
}

// Check whether a string has zero length or contains only whitespace rune(s).
// Available since v0.3.0
func IsEmptyOrWhitespace(str string) bool {
	if len(str) > 0 {
		for _, c := range str {
			if c != 0x0020 && c != 0x0009 {
				return false
			}
		}
	}
	return true
}
