// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package securerng

import "testing"

func TestBase64(t *testing.T) {
	tests := []struct {
		group    string
		len      uint
		expected uint
	}{
		{"zero_byte", 0, 0},
		{"one_byte", 1, 4},
		{"two_bytes", 2, 4},
		{"three_bytes", 3, 4},
		{"four_bytes", 4, 8},
		{"five_bytes", 5, 8},
		{"six_bytes", 6, 8},
		{"seventy_bytes", 70, 96},
		{"seventy_one_bytes", 71, 96},
		{"seventy_two_bytes", 72, 96},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			result := Base64(tt.len)
			if len(result) != int(tt.expected) {
				t.Errorf("expected %v actual %v underlined %v", tt.expected, len(result), result)
			}
		})
	}
}

func TestHex(t *testing.T) {
	tests := []struct {
		group    string
		len      uint
		expected uint
	}{
		{"zero_byte", 0, 0},
		{"one_byte", 1, 2},
		{"two_bytes", 2, 4},
		{"three_bytes", 3, 6},
		{"four_bytes", 4, 8},
		{"five_bytes", 5, 10},
		{"six_bytes", 6, 12},
		{"seventy_bytes", 70, 140},
		{"seventy_one_bytes", 71, 142},
		{"one_hundred_bytes", 100, 200},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			result := Hex(tt.len)
			if len(result) != int(tt.expected) {
				t.Errorf("expected %v actual %v underlined %v", tt.expected, len(result), result)
			}
		})
	}
}
