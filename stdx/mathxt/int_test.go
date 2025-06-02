package mathxt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsInt(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected int
	}{
		{"positive_value", 5, 5},
		{"negative_value", -5, 5},
		{"zero_value", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AbsInt(tt.x)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinInt(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		y        []int
		expected int
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []int{5, 20, 3}, 3},
		{"negative_values", -10, []int{-5, -20, -3}, -20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinInt(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxInt(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		y        []int
		expected int
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []int{5, 20, 3}, 20},
		{"negative_values", -10, []int{-5, -20, -3}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxInt(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAbsInt8(t *testing.T) {
	tests := []struct {
		name     string
		x        int8
		expected int8
	}{
		{"positive_value", 5, 5},
		{"negative_value", -5, 5},
		{"zero_value", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AbsInt8(tt.x)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinInt8(t *testing.T) {
	tests := []struct {
		name     string
		x        int8
		y        []int8
		expected int8
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []int8{5, 20, 3}, 3},
		{"negative_values", -10, []int8{-5, -20, -3}, -20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinInt8(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxInt8(t *testing.T) {
	tests := []struct {
		name     string
		x        int8
		y        []int8
		expected int8
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []int8{5, 20, 3}, 20},
		{"negative_values", -10, []int8{-5, -20, -3}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxInt8(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAbsInt16(t *testing.T) {
	tests := []struct {
		name     string
		x        int16
		expected int16
	}{
		{"positive_value", 5, 5},
		{"negative_value", -5, 5},
		{"zero_value", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AbsInt16(tt.x)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinInt16(t *testing.T) {
	tests := []struct {
		name     string
		x        int16
		y        []int16
		expected int16
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []int16{5, 20, 3}, 3},
		{"negative_values", -10, []int16{-5, -20, -3}, -20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinInt16(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxInt16(t *testing.T) {
	tests := []struct {
		name     string
		x        int16
		y        []int16
		expected int16
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []int16{5, 20, 3}, 20},
		{"negative_values", -10, []int16{-5, -20, -3}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxInt16(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAbsInt32(t *testing.T) {
	tests := []struct {
		name     string
		x        int32
		expected int32
	}{
		{"positive_value", 5, 5},
		{"negative_value", -5, 5},
		{"zero_value", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AbsInt32(tt.x)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinInt32(t *testing.T) {
	tests := []struct {
		name     string
		x        int32
		y        []int32
		expected int32
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []int32{5, 20, 3}, 3},
		{"negative_values", -10, []int32{-5, -20, -3}, -20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinInt32(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxInt32(t *testing.T) {
	tests := []struct {
		name     string
		x        int32
		y        []int32
		expected int32
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []int32{5, 20, 3}, 20},
		{"negative_values", -10, []int32{-5, -20, -3}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxInt32(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAbsInt64(t *testing.T) {
	tests := []struct {
		name     string
		x        int64
		expected int64
	}{
		{"positive_value", 5, 5},
		{"negative_value", -5, 5},
		{"zero_value", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AbsInt64(tt.x)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinInt64(t *testing.T) {
	tests := []struct {
		name     string
		x        int64
		y        []int64
		expected int64
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []int64{5, 20, 3}, 3},
		{"negative_values", -10, []int64{-5, -20, -3}, -20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinInt64(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxInt64(t *testing.T) {
	tests := []struct {
		name     string
		x        int64
		y        []int64
		expected int64
	}{
		{"single_value", 5, nil, 5},
		{"multiple_values", 10, []int64{5, 20, 3}, 20},
		{"negative_values", -10, []int64{-5, -20, -3}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxInt64(tt.x, tt.y...)
			assert.Equal(t, tt.expected, result)
		})
	}
}
