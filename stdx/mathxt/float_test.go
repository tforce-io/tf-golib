package mathxt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
