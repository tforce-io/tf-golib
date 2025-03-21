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

package stringxt

// Check whether a string has zero length.
// Available since v0.3.0
func IsEmpty(str string) bool {
	return str == ""
}

// Check whether a string has zero length or contains only whitespace rune(s).
// Available since v0.3.0
func IsEmptyOrWhitespace(str string) bool {
	if len(str) > 0 {
		for _, c := range str {
			if c != 0x0020 && c != 0x0009 {
				return false
			}
		}
	}
	return true
}
