// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package stdx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBigIntFromString_Binary(t *testing.T) {
	tests := []struct {
		str      string
		expected uint64
	}{
		{"0b0100100001101001", 18537},
		{"0B0100100001101001", 18537},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			bigint, ok := BigIntFromString(tt.str)
			assert.True(t, ok, "Failed to parse octal string")
			assert.Equal(t, tt.expected, bigint.Uint64(), "Wrong value from octal")
		})
	}
}

func TestBigIntFromString_Decimal(t *testing.T) {
	tests := []struct {
		str      string
		expected uint64
	}{
		{"255", 255},
		{"7959", 7959},
		{"0255", 255},
		{"07959", 7959},
		{"007959", 7959},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			bigint, ok := BigIntFromString(tt.str)
			assert.True(t, ok, "Failed to parse decimal string")
			assert.Equal(t, tt.expected, bigint.Uint64(), "Wrong value from decimal")
		})
	}
}

func TestBigIntFromString_Hexadecimal(t *testing.T) {
	tests := []struct {
		str      string
		expected uint64
	}{
		{"0x255", 597},
		{"0X255", 597},
		{"0x1f17", 7959},
		{"0X1F17", 7959},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			bigint, ok := BigIntFromString(tt.str)
			assert.True(t, ok, "Failed to parse hexadecimal string")
			assert.Equal(t, tt.expected, bigint.Uint64(), "Wrong value from hexadecimal")
		})
	}
}

func TestBigIntFromString_Octal(t *testing.T) {
	tests := []struct {
		str      string
		expected uint64
	}{
		{"0o255", 173},
		{"0O255", 173},
		{"0o1717", 975},
		{"0O1717", 975},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			bigint, ok := BigIntFromString(tt.str)
			assert.True(t, ok, "Failed to parse octal string")
			assert.Equal(t, tt.expected, bigint.Uint64(), "Wrong value from octal")
		})
	}
}

func TestBigIntFromString_Invalid(t *testing.T) {
	tests := []struct {
		name string
		str  string
	}{
		{"Invalid decimal character", "1234a567890"},
		{"Invalid decimal character", "12345678z90"},
		{"Invalid hexadecimal character", "0x1234x567890"},
		{"Invalid hexadecimal character", "0X12345678z90"},
		{"Invalid binary character", "0b1f17"},
		{"Invalid binary character", "0b1717"},
		{"Invalid octal character", "0o1f17"},
		{"Invalid octal character", "0oO1717"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := BigIntFromString(tt.str)
			assert.False(t, ok, "Parsed number should be invalid")
		})
	}
}

func TestBigIntEqual(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected bool
	}{
		{"Equal positive numbers", "100", "100", true},
		{"Equal negative numbers", "-100", "-100", true},
		{"Equal zero", "0", "0", true},
		{"Different positive numbers", "100", "200", false},
		{"Different negative numbers", "-100", "-200", false},
		{"Positive and negative", "100", "-100", false},
		{"Zero and positive", "0", "1", false},
		{"Zero and negative", "0", "-1", false},
		{"Large equal numbers", "999999999999999999999999999999", "999999999999999999999999999999", true},
		{"Large different numbers", "999999999999999999999999999999", "999999999999999999999999999998", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := BigIntFromString(tt.a)
			b, _ := BigIntFromString(tt.b)
			result := BigIntEqual(a, b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBigIntGreatorThan(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected bool
	}{
		{"Greater positive numbers", "200", "100", true},
		{"Greater negative numbers", "-100", "-200", true},
		{"Positive greater than negative", "1", "-1", true},
		{"Positive greater than zero", "1", "0", true},
		{"Zero greater than negative", "0", "-1", true},
		{"Not greater - equal", "100", "100", false},
		{"Not greater - less than", "100", "200", false},
		{"Negative not greater than positive", "-1", "1", false},
		{"Negative not greater than zero", "-1", "0", false},
		{"Large number greater", "1000000000000000000000000000000", "999999999999999999999999999999", true},
		{"Large number not greater", "999999999999999999999999999999", "1000000000000000000000000000000", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := BigIntFromString(tt.a)
			b, _ := BigIntFromString(tt.b)
			result := BigIntGreatorThan(a, b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBigIntGreatorThanEqual(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected bool
	}{
		{"Greater positive numbers", "200", "100", true},
		{"Greater negative numbers", "-100", "-200", true},
		{"Equal positive numbers", "100", "100", true},
		{"Equal negative numbers", "-100", "-100", true},
		{"Equal zero", "0", "0", true},
		{"Positive greater than negative", "1", "-1", true},
		{"Positive greater than zero", "1", "0", true},
		{"Zero greater than negative", "0", "-1", true},
		{"Not greater or equal - less than", "100", "200", false},
		{"Negative not greater or equal to positive", "-1", "1", false},
		{"Negative not greater or equal to zero", "-1", "0", false},
		{"Large number greater or equal (greater)", "1000000000000000000000000000000", "999999999999999999999999999999", true},
		{"Large number greater or equal (equal)", "999999999999999999999999999999", "999999999999999999999999999999", true},
		{"Large number not greater or equal", "999999999999999999999999999999", "1000000000000000000000000000000", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := BigIntFromString(tt.a)
			b, _ := BigIntFromString(tt.b)
			result := BigIntGreatorThanEqual(a, b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBigIntLessThan(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected bool
	}{
		{"Less than positive numbers", "100", "200", true},
		{"Less than negative numbers", "-200", "-100", true},
		{"Negative less than positive", "-1", "1", true},
		{"Negative less than zero", "-1", "0", true},
		{"Zero less than positive", "0", "1", true},
		{"Not less than - equal", "100", "100", false},
		{"Not less than - greater", "200", "100", false},
		{"Positive not less than negative", "1", "-1", false},
		{"Zero not less than negative", "0", "-1", false},
		{"Large number less than", "999999999999999999999999999999", "1000000000000000000000000000000", true},
		{"Large number not less than", "1000000000000000000000000000000", "999999999999999999999999999999", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := BigIntFromString(tt.a)
			b, _ := BigIntFromString(tt.b)
			result := BigIntLessThan(a, b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBigIntLessThanEqual(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected bool
	}{
		{"Less than positive numbers", "100", "200", true},
		{"Less than negative numbers", "-200", "-100", true},
		{"Equal positive numbers", "100", "100", true},
		{"Equal negative numbers", "-100", "-100", true},
		{"Equal zero", "0", "0", true},
		{"Negative less than positive", "-1", "1", true},
		{"Negative less than zero", "-1", "0", true},
		{"Zero less than positive", "0", "1", true},
		{"Not less than or equal - greater", "200", "100", false},
		{"Positive not less than or equal to negative", "1", "-1", false},
		{"Zero not less than or equal to negative", "0", "-1", false},
		{"Large number less than or equal (less)", "999999999999999999999999999999", "1000000000000000000000000000000", true},
		{"Large number less than or equal (equal)", "999999999999999999999999999999", "999999999999999999999999999999", true},
		{"Large number not less than or equal", "1000000000000000000000000000000", "999999999999999999999999999999", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := BigIntFromString(tt.a)
			b, _ := BigIntFromString(tt.b)
			result := BigIntLessThanEqual(a, b)
			assert.Equal(t, tt.expected, result)
		})
	}
}
