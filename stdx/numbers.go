// Copyright (C) 2025 T-Force I/O
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stdx

import (
	"math/big"
)

// A Uint256 is wrapped *big.Int to intergrate with system that use large number and provide fluent apis for type casting.
// Added in v0.1.0
type Uint256 struct {
	value *big.Int
}

// Cast Uint256 to *big.Int.
// Added in v0.1.0
func (s *Uint256) BigInt() *big.Int {
	return s.value
}

// Cast Uint256 to uint64. This may cause data loss when the value is larger than 18446744073709551615.
// Added in v0.1.0
func (s *Uint256) Uint64() uint64 {
	return s.value.Uint64()
}

// Returns underlying value of Bytes as *big.Int.
// Added in v0.1.0
func (s *Uint256) Value() *big.Int {
	return s.value
}
