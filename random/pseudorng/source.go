// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package pseudorng

import "math/rand"

// Seed uses the provided seed value to initialize the default Source to a
// deterministic state. Seed values that have the same remainder when
// divided by 2³¹-1 generate the same pseudo-random sequence.
// Seed, unlike the [Rand.Seed] method, is safe for concurrent use.
//
// If Seed is not called, the generator is seeded randomly at program startup.
// Available since v0.3.0
func Seed(seed int64) {
	rand.Seed(seed)
}
