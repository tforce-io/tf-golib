// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package opx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTry(t *testing.T) {
	t.Run("does not panic when ok is true", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Try(true)
		})
	})
	t.Run("panics when ok is false", func(t *testing.T) {
		assert.Panics(t, func() {
			Try(false)
		})
	})
}

func TestTry1(t *testing.T) {
	t.Run("returns value when ok is true", func(t *testing.T) {
		val := 17
		result := Try1(val, true)
		assert.Equal(t, val, result)
	})
	t.Run("panics when ok is false", func(t *testing.T) {
		assert.Panics(t, func() {
			Try1(17, false)
		})
	})
}

func TestTry2(t *testing.T) {
	t.Run("returns values when ok is true", func(t *testing.T) {
		a, b := 1, "foo"
		r1, r2 := Try2(a, b, true)
		assert.Equal(t, a, r1)
		assert.Equal(t, b, r2)
	})
	t.Run("panics when ok is false", func(t *testing.T) {
		assert.Panics(t, func() {
			Try2(1, "foo", false)
		})
	})
}

func TestTry3(t *testing.T) {
	t.Run("returns values when ok is true", func(t *testing.T) {
		a, b, c := 1, "foo", 3.14
		r1, r2, r3 := Try3(a, b, c, true)
		assert.Equal(t, a, r1)
		assert.Equal(t, b, r2)
		assert.Equal(t, c, r3)
	})
	t.Run("panics when ok is false", func(t *testing.T) {
		assert.Panics(t, func() {
			Try3(1, "foo", 3.14, false)
		})
	})
}

func TestTry4(t *testing.T) {
	t.Run("returns values when ok is true", func(t *testing.T) {
		a, b, c, d := 1, "foo", 3.14, true
		r1, r2, r3, r4 := Try4(a, b, c, d, true)
		assert.Equal(t, a, r1)
		assert.Equal(t, b, r2)
		assert.Equal(t, c, r3)
		assert.Equal(t, d, r4)
	})
	t.Run("panics when ok is false", func(t *testing.T) {
		assert.Panics(t, func() {
			Try4(1, "foo", 3.14, true, false)
		})
	})
}

func TestTry5(t *testing.T) {
	t.Run("returns values when ok is true", func(t *testing.T) {
		a, b, c, d, e := 1, "foo", 3.14, true, int64(99)
		r1, r2, r3, r4, r5 := Try5(a, b, c, d, e, true)
		assert.Equal(t, a, r1)
		assert.Equal(t, b, r2)
		assert.Equal(t, c, r3)
		assert.Equal(t, d, r4)
		assert.Equal(t, e, r5)
	})
	t.Run("panics when ok is false", func(t *testing.T) {
		assert.Panics(t, func() {
			Try5(1, "foo", 3.14, true, int64(99), false)
		})
	})
}
