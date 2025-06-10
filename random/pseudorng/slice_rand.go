// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package pseudorng

import "math/rand"

// Read generates len(p) random bytes from the default [Source] and
// writes them into p. It always returns len(p) and a nil error.
// Read, unlike the [Rand.Read] method, is safe for concurrent use.
// Available since v0.3.0
func Read(p []byte) (n int, err error) {
	return rand.Read(p)
}

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers
// in the half-open interval [0,n) from the default [Source].
// Available since v0.3.0
func Perm(n int) []int {
	return rand.Perm(n)
}

// Shuffle pseudo-randomizes the order of elements using the default [Source].
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
// Available since v0.3.0
func Shuffle(n int, swap func(i, j int)) {
	rand.Shuffle(n, swap)
}
