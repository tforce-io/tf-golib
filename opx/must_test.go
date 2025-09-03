// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package opx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	t.Run("does not panic when err is nil", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Must(nil)
		})
	})

	t.Run("panics when err is not nil", func(t *testing.T) {
		err := errors.New("test error")
		assert.Panics(t, func() {
			Must(err)
		})
	})
}

func TestMust1(t *testing.T) {
	t.Run("returns value when err is nil", func(t *testing.T) {
		val := 17
		result := Must1(val, nil)
		assert.Equal(t, val, result)
	})

	t.Run("panics when err is not nil", func(t *testing.T) {
		val := 17
		err := errors.New("test error")
		assert.Panics(t, func() {
			Must1(val, err)
		})
	})
}

func TestMust2(t *testing.T) {
	t.Run("returns values when err is nil", func(t *testing.T) {
		a, b := 1, "foo"
		r1, r2 := Must2(a, b, nil)
		assert.Equal(t, a, r1)
		assert.Equal(t, b, r2)
	})

	t.Run("panics when err is not nil", func(t *testing.T) {
		a, b := 1, "foo"
		err := errors.New("test error")
		assert.Panics(t, func() {
			Must2(a, b, err)
		})
	})
}

func TestMust3(t *testing.T) {
	t.Run("returns values when err is nil", func(t *testing.T) {
		a, b, c := 1, "foo", 3.14
		r1, r2, r3 := Must3(a, b, c, nil)
		assert.Equal(t, a, r1)
		assert.Equal(t, b, r2)
		assert.Equal(t, c, r3)
	})

	t.Run("panics when err is not nil", func(t *testing.T) {
		a, b, c := 1, "foo", 3.14
		err := errors.New("test error")
		assert.Panics(t, func() {
			Must3(a, b, c, err)
		})
	})
}

func TestMust4(t *testing.T) {
	t.Run("returns values when err is nil", func(t *testing.T) {
		a, b, c, d := 1, "foo", 3.14, true
		r1, r2, r3, r4 := Must4(a, b, c, d, nil)
		assert.Equal(t, a, r1)
		assert.Equal(t, b, r2)
		assert.Equal(t, c, r3)
		assert.Equal(t, d, r4)
	})

	t.Run("panics when err is not nil", func(t *testing.T) {
		a, b, c, d := 1, "foo", 3.14, true
		err := errors.New("test error")
		assert.Panics(t, func() {
			Must4(a, b, c, d, err)
		})
	})
}

func TestMust5(t *testing.T) {
	t.Run("returns values when err is nil", func(t *testing.T) {
		a, b, c, d, e := 1, "foo", 3.14, true, int64(99)
		r1, r2, r3, r4, r5 := Must5(a, b, c, d, e, nil)
		assert.Equal(t, a, r1)
		assert.Equal(t, b, r2)
		assert.Equal(t, c, r3)
		assert.Equal(t, d, r4)
		assert.Equal(t, e, r5)
	})

	t.Run("panics when err is not nil", func(t *testing.T) {
		a, b, c, d, e := 1, "foo", 3.14, true, int64(99)
		err := errors.New("test error")
		assert.Panics(t, func() {
			Must5(a, b, c, d, e, err)
		})
	})
}
