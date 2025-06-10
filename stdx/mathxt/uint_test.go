// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package mathxt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinUint(t *testing.T) {
	tests := []struct {
		name     string
		x        uint
		y        []uint
		expected uint
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []uint{5, 20, 3}, 3},
		{"all_equal", 10, []uint{10, 10, 10}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinUint(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxUint(t *testing.T) {
	tests := []struct {
		name     string
		x        uint
		y        []uint
		expected uint
	}{
		{"single_value", 10, nil, 10},
		{"multiple_values", 10, []uint{5, 15, 25}, 25},
		{"all_equal", 10, []uint{10, 10, 10}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxUint(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinUint8(t *testing.T) {
	tests := []struct {
		name     string
		x        uint8
		y        []uint8
		expected uint8
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []uint8{5, 20, 3}, 3},
		{"all_equal", 10, []uint8{10, 10, 10}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinUint8(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxUint8(t *testing.T) {
	tests := []struct {
		name     string
		x        uint8
		y        []uint8
		expected uint8
	}{
		{"single_value", 10, nil, 10},
		{"multiple_values", 10, []uint8{5, 15, 25}, 25},
		{"all_equal", 10, []uint8{10, 10, 10}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxUint8(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinUint16(t *testing.T) {
	tests := []struct {
		name     string
		x        uint16
		y        []uint16
		expected uint16
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []uint16{5, 20, 3}, 3},
		{"all_equal", 10, []uint16{10, 10, 10}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinUint16(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxUint16(t *testing.T) {
	tests := []struct {
		name     string
		x        uint16
		y        []uint16
		expected uint16
	}{
		{"single value", 10, nil, 10},
		{"two values", 10, []uint16{20}, 20},
		{"multiple values", 10, []uint16{5, 15, 25}, 25},
		{"all equal values", 10, []uint16{10, 10, 10}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxUint16(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinUint32(t *testing.T) {
	tests := []struct {
		name     string
		x        uint32
		y        []uint32
		expected uint32
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []uint32{5, 20, 3}, 3},
		{"all_equal", 10, []uint32{10, 10, 10}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinUint32(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxUint32(t *testing.T) {
	tests := []struct {
		name     string
		x        uint32
		y        []uint32
		expected uint32
	}{
		{"single_value", 10, nil, 10},
		{"multiple_values", 10, []uint32{5, 15, 25}, 25},
		{"all_equal", 10, []uint32{10, 10, 10}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxUint32(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinUint64(t *testing.T) {
	tests := []struct {
		name     string
		x        uint64
		y        []uint64
		expected uint64
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []uint64{5, 20, 3}, 3},
		{"all_equal", 10, []uint64{10, 10, 10}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinUint64(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxUint64(t *testing.T) {
	tests := []struct {
		name     string
		x        uint64
		y        []uint64
		expected uint64
	}{
		{"single_value", 10, nil, 10},
		{"multiple_values", 10, []uint64{5, 15, 25}, 25},
		{"all_equal", 10, []uint64{10, 10, 10}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxUint64(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}
