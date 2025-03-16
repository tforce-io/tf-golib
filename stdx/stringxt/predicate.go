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

import (
	"regexp"
	"strings"
)

// A Predicate holds all the requirements to test a string for matching.
// Available since v0.4.0
type Predicate struct {
	Prefix string
	Suffix string
	Regexp string
}

// Test whether string str satisfied all requirements in the Predicate.
// In the case Regexp is invalid, the result is always false.
// Available since v0.4.0
func (p *Predicate) Match(str string) (bool, error) {
	result := true
	var err error
	if p.Prefix != "" {
		if !strings.HasPrefix(str, p.Prefix) {
			result = false
		}
	}
	if p.Suffix != "" {
		if !strings.HasSuffix(str, p.Suffix) {
			result = false
		}
	}
	if p.Regexp != "" {
		r, err := regexp.MatchString(p.Regexp, str)
		if !r || err != nil {
			result = false
		}
	}
	return result, err
}
