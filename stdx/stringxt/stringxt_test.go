// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package stringxt

import "testing"

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		group    string
		str      string
		expected bool
	}{
		// empty
		{"empty", "", true},
		// ascii
		{"ascii", "a", false},
		{"ascii", "hello", false},
		{"ascii", "hello world", false},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			result := IsEmpty(tt.str)
			if result != tt.expected {
				t.Errorf("expected %v actual %v intermediate %v", tt.expected, result, tt.str)
			}
		})
	}
}

func TestIsEmptyOrWhitespace(t *testing.T) {
	tests := []struct {
		group    string
		str      string
		expected bool
	}{
		// empty
		{"empty", "", true},
		// all whitespace
		{"all whitespace", " ", true},
		{"all whitespace", "  ", true},
		{"all whitespace", "       ", true},
		// all tabs
		{"all tabs", "\t", true},
		{"all tabs", "\t\t", true},
		{"all tabs", "\t\t\t\t\t\t\t", true},
		// mixed whitespaces and tabs
		{"mixed whitespaces and tabs", " \t", true},
		{"mixed whitespaces and tabs", " \t ", true},
		{"mixed whitespaces and tabs", " \t \t", true},
		{"mixed whitespaces and tabs", " \t \t\t \t", true},
		{"mixed whitespaces and tabs", "\t ", true},
		{"mixed whitespaces and tabs", "\t \t", true},
		{"mixed whitespaces and tabs", "\t \t ", true},
		{"mixed whitespaces and tabs", "\t \t  \t ", true},
		// ascii
		{"ascii", "a", false},
		{"ascii", "hello", false},
		{"ascii", "hello world", false},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			result := IsEmptyOrWhitespace(tt.str)
			if result != tt.expected {
				t.Errorf("expected %v actual %v intermediate %v", tt.expected, result, tt.str)
			}
		})
	}
}
