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

package opx

import (
	"math/big"
	"testing"
)

func TestAreEqualSlice(t *testing.T) {
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
		{"one nil", nil, []int{}, false},
		// different length
		{"different length", []int{1}, []int{1, 2}, false},
		{"different length", []int{1}, []int{1, 2, 3}, false},
		{"different length", []int{1, 2}, []int{1}, false},
		{"different length", []int{1, 2}, []int{1, 2, 3}, false},
		{"different length", []int{1, 2}, []int{1, 2, 3, 4}, false},
		{"different length", []int{1, 2, 3, 4}, []int{1, 2, 3}, false},
		// deep compare
		{"deep compare", []int{}, []int{}, true},
		{"deep compare", []int{1}, []int{1}, true},
		{"deep compare", []int{1, 3, 5}, []int{1, 3, 5}, true},
		{"deep compare", []int{5, 3, 1}, []int{1, 3, 5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			result := AreEqualSlice(tt.x, tt.y)
			if result != tt.expected {
				t.FailNow()
			}
		})
	}
}

func TestCoalesce(t *testing.T) {
	tests := []struct {
		group    string
		items    []*big.Int
		expected *big.Int
	}{
		// 1 element
		{"1 element", []*big.Int{nil}, nil},
		{"1 element", []*big.Int{big.NewInt(100)}, big.NewInt(100)},
		// 2 elements
		{"2 elements", []*big.Int{nil, nil}, nil},
		{"2 elements", []*big.Int{big.NewInt(100), big.NewInt(101)}, big.NewInt(100)},
		{"2 elements", []*big.Int{big.NewInt(100), nil}, big.NewInt(100)},
		{"2 elements", []*big.Int{nil, big.NewInt(101)}, big.NewInt(101)},
		// 3 elements
		{"3 elements", []*big.Int{nil, nil, nil}, nil},
		{"3 elements", []*big.Int{big.NewInt(100), big.NewInt(101), big.NewInt(102)}, big.NewInt(100)},
		{"3 elements", []*big.Int{big.NewInt(100), nil, big.NewInt(102)}, big.NewInt(100)},
		{"3 elements", []*big.Int{big.NewInt(100), nil, nil}, big.NewInt(100)},
		{"3 elements", []*big.Int{nil, big.NewInt(101), big.NewInt(102)}, big.NewInt(101)},
		{"3 elements", []*big.Int{nil, big.NewInt(101), nil}, big.NewInt(101)},
		{"3 elements", []*big.Int{nil, nil, big.NewInt(102)}, big.NewInt(102)},
		// 5 elements
		{"5 elements", []*big.Int{nil, nil, nil, nil, nil}, nil},
		{"5 elements", []*big.Int{big.NewInt(100), big.NewInt(101), big.NewInt(102), big.NewInt(103), big.NewInt(104)}, big.NewInt(100)},
		{"5 elements", []*big.Int{big.NewInt(100), nil, nil, big.NewInt(103), nil}, big.NewInt(100)},
		{"5 elements", []*big.Int{nil, big.NewInt(101), big.NewInt(102), big.NewInt(103), big.NewInt(104)}, big.NewInt(101)},
		{"5 elements", []*big.Int{nil, big.NewInt(101), nil, big.NewInt(103), nil}, big.NewInt(101)},
		{"5 elements", []*big.Int{nil, nil, big.NewInt(102), big.NewInt(103), big.NewInt(104)}, big.NewInt(102)},
		{"5 elements", []*big.Int{nil, nil, big.NewInt(102), big.NewInt(103), nil}, big.NewInt(102)},
		{"5 elements", []*big.Int{nil, nil, nil, big.NewInt(103), big.NewInt(104)}, big.NewInt(103)},
		{"5 elements", []*big.Int{nil, nil, nil, big.NewInt(103), nil}, big.NewInt(103)},
		{"5 elements", []*big.Int{nil, nil, nil, nil, big.NewInt(104)}, big.NewInt(104)},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			result := Coalesce(tt.items...)
			if result.Cmp(tt.expected) != 0 {
				t.FailNow()
			}
		})
	}
}

func TestIsEmptySlice(t *testing.T) {
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
			result := IsEmptySlice(tt.slice)
			if result != tt.expected {
				t.FailNow()
			}
		})
	}
}

func TestIsEmptyString(t *testing.T) {
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
			result := IsEmptyString(tt.str)
			if result != tt.expected {
				t.FailNow()
			}
		})
	}
}

func TestIsEmptyOrWhitespaceString(t *testing.T) {
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
			result := IsEmptyOrWhitespaceString(tt.str)
			if result != tt.expected {
				t.FailNow()
			}
		})
	}
}

func TestTernary(t *testing.T) {
	tests := []struct {
		group    string
		cond     bool
		x        int
		y        int
		expected int
	}{
		{"true", true, 100, 200, 100},
		{"false", false, 100, 200, 200},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			result := Ternary(tt.cond, tt.x, tt.y)
			if result != tt.expected {
				t.FailNow()
			}
		})
	}
}
