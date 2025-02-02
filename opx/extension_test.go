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
