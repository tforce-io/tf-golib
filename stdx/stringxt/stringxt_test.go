// Copyright (C) 2025 T-Force I/O
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
