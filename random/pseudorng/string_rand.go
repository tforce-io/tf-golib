// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package pseudorng

import "strings"

// Base64 returns a pseudo-random base64 string with charset defined in RFC 4648 for UTF-8
// from the default [Source].
// The length (n) denotes number of bytes the base64 string represented, not the
// actual number of characters of the generated base64 string.
// Available since v0.3.0
func Base64(n uint) string {
	charset := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'+', '/',
	}
	ret := String(n+ceilUint(n, 3), charset)
	switch n % 3 {
	case 1:
		return ret + "=="
	case 2:
		return ret + "="
	default:
		return ret
	}
}

// Hex returns a pseudo-random hex string in lower case
// from the default [Source].
// The length (n) denotes number of bytes the hex string represented, not the
// actual number of characters of the generated hex string.
// Available since v0.3.0
func Hex(n uint) string {
	charset := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f',
	}
	return String(n*2, charset)
}

// String returns a pseudo-random string with specific length (n) and based on
// pre-defined list of characters from the default [Source].
// The length of charset should be less than 2³¹ for reliability.
// The probability for appearances of each character cannot be customized.
// Available since v0.3.0
func String(n uint, charset []rune) string {
	var sb strings.Builder
	charcount := len(charset)
	for i := uint(0); i < n; i++ {
		sb.WriteRune(charset[Intn(charcount)])
	}
	return sb.String()
}

// ceilUint returns ceiling result of x / y as same math.Ceil.
// See https://stackoverflow.com/a/2745086 for more information.
// Available since v0.3.0
func ceilUint(x, y uint) uint {
	if x == 0 {
		return 0
	}
	return 1 + ((x - 1) / y)
}
