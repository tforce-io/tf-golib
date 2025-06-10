// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package securerng

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

// Base64 returns a uniform base64 string with charset defined in RFC 4648 for UTF-8
// from the default [Reader].
// The length (n) denotes number of bytes the base64 string represented, not the
// actual number of characters of the generated base64 string.
// Available since v0.3.0
func Base64(n uint) string {
	var sb strings.Builder
	i := uint(0)
	buflen := uint(30)
	for i < n {
		var buf []byte
		if i+buflen <= n {
			buf = make([]byte, buflen)
		} else {
			buf = make([]byte, n-i)
		}
		_, err := rand.Read(buf)
		if err != nil {
			panic(err)
		}
		sb.WriteString(base64.StdEncoding.EncodeToString(buf))
		i += buflen
	}
	return sb.String()
}

// Hex returns a uniform hex string in lower case
// from the default [Reader].
// The length (n) denotes number of bytes the hex string represented, not the
// actual number of characters of the generated hex string.
// Available since v0.3.0
func Hex(n uint) string {
	var sb strings.Builder
	i := uint(0)
	buflen := uint(32)
	for i < n {
		var buf []byte
		if i+buflen <= n {
			buf = make([]byte, buflen)
		} else {
			buf = make([]byte, n-i)
		}
		_, err := rand.Read(buf)
		if err != nil {
			panic(err)
		}
		sb.WriteString(hex.EncodeToString(buf))
		i += buflen
	}
	return sb.String()
}

// String returns a uniform string with specific length (n) and based on
// pre-defined list of characters from the default [Reader].
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
