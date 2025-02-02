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

func TestAreEqualSlices(t *testing.T) {
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
			result := AreEqualSlices(tt.x, tt.y)
			if result != tt.expected {
				t.Errorf("expected %v actual left %v actual right %v ", tt.expected, tt.x, tt.y)
			}
			result = AreEqualSlices(tt.y, tt.x)
			if result != tt.expected {
				t.Errorf("expected %v actual left %v actual right %v ", tt.expected, tt.y, tt.x)
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
				t.Errorf("expected %v actual %v intermediate %v", tt.expected, result, tt.items)
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
				t.Errorf("expected %v actual %v intermediate %v", tt.expected, result, tt.slice)
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
				t.Errorf("expected %v actual left %v actual right %v ", tt.expected, tt.x, tt.y)
			}
		})
	}
}
