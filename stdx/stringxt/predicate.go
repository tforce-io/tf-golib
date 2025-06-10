// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

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
