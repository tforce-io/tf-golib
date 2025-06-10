// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package slicext

import (
	"testing"
)

func TestAreEqual(t *testing.T) {
	tests := []struct {
		group    string
		x        []int
		y        []int
		expected bool
	}{
		// all nils
		{"all nils", nil, nil, true},
		// one nils
		{"one nil", []int{}, nil, false},
		// different length
		{"different length", []int{1}, []int{1, 2}, false},
		{"different length", []int{1}, []int{1, 2, 3}, false},
		{"different length", []int{1, 2}, []int{1, 2, 3}, false},
		{"different length", []int{1, 2}, []int{1, 2, 3, 4}, false},
		// deep compare
		{"deep compare", []int{}, []int{}, true},
		{"deep compare", []int{1}, []int{1}, true},
		{"deep compare", []int{1, 3, 5}, []int{1, 3, 5}, true},
		{"deep compare", []int{5, 3, 1}, []int{1, 3, 5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			result := AreEqual(tt.x, tt.y)
			if result != tt.expected {
				t.Errorf("expected %v actual left %v actual right %v ", tt.expected, tt.x, tt.y)
			}
			result = AreEqual(tt.y, tt.x)
			if result != tt.expected {
				t.Errorf("expected %v actual left %v actual right %v ", tt.expected, tt.y, tt.x)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		group    string
		items    []int
		value    int
		expected bool
	}{
		{"true", []int{3, 5, 7, 2, 4, 6}, 7, true},
		{"true", []int{3, 5, 7, 2, 4, 6}, 4, true},
		{"true", []int{}, 7, false},
		{"true", []int{3, 5, 2, 4, 6}, 7, false},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			result := Contains(tt.items, tt.value)
			if result != tt.expected {
				t.Errorf("expected %v actual %v intermediate %v %v", tt.expected, result, tt.items, tt.value)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		group    string
		slice    []int
		expected bool
	}{
		// nil
		{"nil", nil, true},
		// empty
		{"empty", []int{}, true},
		// one or more
		{"one or more", []int{1}, false},
		{"one or more", []int{1, 2, 3}, false},
		{"one or more", []int{1, 2, 3, 4}, false},
		{"one or more", []int{1, 2, 3, 4, 5, 6, 7}, false},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			result := IsEmpty(tt.slice)
			if result != tt.expected {
				t.Errorf("expected %v actual %v intermediate %v", tt.expected, result, tt.slice)
			}
		})
	}
}
