// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package stdx

import "encoding/hex"

// A Bytes is wrapped byte slice to provide fluent api for type casting.
// Available since v0.1.0
type Bytes []byte

// Cast Bytes as byte slice.
// Available since v0.1.0
func (s Bytes) ByteArr() []byte {
	return s
}

// Cast Bytes to a Hex.
// Available since v0.1.0
func (s Bytes) Hex() *Hex {
	return NewHex(s, false)
}

// Cast Bytes to a hex string without "0x" prefix.
// Available since v0.1.0
func (s Bytes) HexStr() string {
	return hex.EncodeToString(s)
}

// Return underlying value of Bytes as byte slice.
// Available since v0.1.0
func (s Bytes) Value() []byte {
	return s
}

// A Hex is another presetnation of hex string to provide fluent api for type casting.
// Available since v0.1.0
type Hex struct {
	value     Bytes
	hasPrefix bool
}

// Create new instance of Hex using Bytes.
// Must indicate whether that Hex has prefix of "0x".
// Available since v0.1.0
func NewHex(value Bytes, hasPrefix bool) *Hex {
	return &Hex{value, hasPrefix}
}

// Cast Hex as Bytes.
// Available since v0.1.0
func (s *Hex) Bytes() Bytes {
	return s.value
}

// Cast Hex as byte slice.
// Available since v0.1.0
func (s *Hex) ByteArr() []byte {
	return s.value
}

// Return underlying value of Hex as string.
// Available since v0.1.0
func (s *Hex) Value() string {
	if s.hasPrefix {
		return "0x" + s.value.HexStr()
	}
	return s.value.HexStr()
}
