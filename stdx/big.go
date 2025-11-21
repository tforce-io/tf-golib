// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package stdx

import (
	"math/big"
)

// A Uint256 is wrapped *big.Int to intergrate with system that use large number
// and provide fluent apis for type casting.
//
// Available since v0.9.0
type Uint256 big.Int

// Cast Uint256 to *big.Int.
// Available since v0.1.0
func (s *Uint256) BigInt() *big.Int {
	return (*big.Int)(s)
}

// Return underlying value of Bytes as *big.Int.
// Available since v0.1.0
func (s *Uint256) Value() *big.Int {
	return (*big.Int)(s)
}

// Parse a string into big.Int. Decimal doesn't have prefix. Hexadecimal, Binary,
// and Octal must have prefix 0x, 0b, 0o respectively.
//
// Available since v0.9.0
func BigIntFromString(s string) (*big.Int, bool) {
	if s == "" {
		return new(big.Int), true
	}
	var bigint *big.Int
	var ok bool
	if len(s) >= 2 && s[0] == '0' {
		switch s[1] {
		case 'x', 'X':
			bigint, ok = new(big.Int).SetString(s[2:], 16)
		case 'b', 'B':
			bigint, ok = new(big.Int).SetString(s[2:], 2)
		case 'o', 'O':
			bigint, ok = new(big.Int).SetString(s[2:], 8)
		default:
			bigint, ok = new(big.Int).SetString(s, 10)
		}
	} else {
		bigint, ok = new(big.Int).SetString(s, 10)
	}
	return bigint, ok
}
