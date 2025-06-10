// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

/*
Package pseudorng derives from math/rand package and provides pseudo-random
number generators suitable for tasks such as simulation, but it should not
be used for security-sensitive work.
The Top-level functions, such as [Int] and [Float64], are safe for concurrent
use by multiple goroutines.
Available since v0.3.0
*/
package pseudorng
