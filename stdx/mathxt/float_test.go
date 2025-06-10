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

func TestAbsFloat32(t *testing.T) {
	tests := []struct {
		name     string
		x        float32
		expected float32
	}{
		{"positive_value", 5.5, 5.5},
		{"negative_value", -5.5, 5.5},
		{"zero_value", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AbsFloat32(tt.x)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinFloat32(t *testing.T) {
	tests := []struct {
		name     string
		x        float32
		y        []float32
		expected float32
	}{
		{"single_value", 5.5, nil, 5.5},
		{"multiple_values", 10.5, []float32{5.5, 20.5, 3.5}, 3.5},
		{"all_equal", 10.5, []float32{10.5, 10.5, 10.5}, 10.5},
		{"negative_values", -5.5, []float32{-10.5, -3.5}, -10.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinFloat32(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxFloat32(t *testing.T) {
	tests := []struct {
		name     string
		x        float32
		y        []float32
		expected float32
	}{
		{"single_value", 10.5, nil, 10.5},
		{"multiple_values", 10.5, []float32{5.5, 15.5, 25.5}, 25.5},
		{"all_equal", 10.5, []float32{10.5, 10.5, 10.5}, 10.5},
		{"negative_values", -5.5, []float32{-10.5, -3.5}, -3.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxFloat32(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAbsFloat64(t *testing.T) {
	tests := []struct {
		name     string
		x        float64
		expected float64
	}{
		{"positive_value", 5.5, 5.5},
		{"negative_value", -5.5, 5.5},
		{"zero_value", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AbsFloat64(tt.x)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinFloat64(t *testing.T) {
	tests := []struct {
		name     string
		x        float64
		y        []float64
		expected float64
	}{
		{"single_value", 5.5, nil, 5.5},
		{"multiple_values", 10.5, []float64{5.5, 20.5, 3.5}, 3.5},
		{"all_equal", 10.5, []float64{10.5, 10.5, 10.5}, 10.5},
		{"negative_values", -5.5, []float64{-10.5, -3.5}, -10.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinFloat64(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxFloat64(t *testing.T) {
	tests := []struct {
		name     string
		x        float64
		y        []float64
		expected float64
	}{
		{"single_value", 10.5, nil, 10.5},
		{"multiple_values", 10.5, []float64{5.5, 15.5, 25.5}, 25.5},
		{"all_equal", 10.5, []float64{10.5, 10.5, 10.5}, 10.5},
		{"negative_values", -5.5, []float64{-10.5, -3.5}, -3.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxFloat64(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}
