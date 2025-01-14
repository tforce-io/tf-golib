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

package strfmt

import (
	"testing"

	"github.com/tforce-io/tf-golib/stdx"
	"github.com/tforce-io/tf-golib/stdx/opx"
)

func TestNewFileName(t *testing.T) {
	tests := []struct {
		group    string
		name     string
		ext      string
		expected *FileName
	}{
		{"empty", "", "", &FileName{}},
		{"name_only", "main", "", &FileName{Name: "main"}},
		{"name_and_ext", "main", ".go", &FileName{Name: "main", Extension: ".go"}},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			result := NewFileName(tt.name, tt.ext)
			if !AreEqualFileNames(result, tt.expected) {
				t.FailNow()
			}
		})
	}
}

func TestNewFileNameFromStr(t *testing.T) {
	type Test struct {
		group    string
		path     string
		expected *FileName
	}

	tests := opx.Ternary(stdx.IsWindows(),
		[]Test{
			{"empty", "", &FileName{}},
			{"root", "/", &FileName{}},
			{"root", `\`, &FileName{}},
			{"root", "//", &FileName{}},
			{"root", `\\`, &FileName{}},
			{"drive", `c:`, &FileName{}},
			{"drive", `c:\`, &FileName{}},
			{"drive", `c:::`, &FileName{Name: "::"}},
			{"drive", `c:abc`, &FileName{Name: "abc"}},
			{"trailing_slash", "/usr/local/tforce-io/tf-golib/", &FileName{Name: "tf-golib"}},
			{"trailing_slash", "///usr/local/tforce-io/tf-golib/", &FileName{Name: "tf-golib"}},
			{"name_only", "main", &FileName{Name: "main"}},
			{"name_and_ext", "main.go", &FileName{Name: "main", Extension: ".go"}},
			{"relative_path_and_filename", "tf-golib/main", &FileName{Name: "main"}},
			{"relative_path_and_filename", "tf-golib/main.go", &FileName{Name: "main", Extension: ".go"}},
			{"absolute_path_and_filename", "/usr/local/tforce-io/tf-golib/main", &FileName{Name: "main"}},
			{"absolute_path_and_filename", "/usr/local/tforce-io/tf-golib/main.go", &FileName{Name: "main", Extension: ".go"}},
		},
		[]Test{
			{"empty", "", &FileName{}},
			{"root", "/", &FileName{}},
			{"root", `\`, &FileName{}},
			{"root", "//", &FileName{}},
			{"root", `\\`, &FileName{}},
			{"trailing_slash", "/usr/local/tforce-io/tf-golib/", &FileName{Name: "tf-golib"}},
			{"trailing_slash", "///usr/local/tforce-io/tf-golib/", &FileName{Name: "tf-golib"}},
			{"name_only", "main", &FileName{Name: "main"}},
			{"name_and_ext", "main.go", &FileName{Name: "main", Extension: ".go"}},
			{"relative_path_and_filename", "tf-golib/main", &FileName{Name: "main"}},
			{"relative_path_and_filename", "tf-golib/main.go", &FileName{Name: "main", Extension: ".go"}},
			{"absolute_path_and_filename", "/usr/local/tforce-io/tf-golib/main", &FileName{Name: "main"}},
			{"absolute_path_and_filename", "/usr/local/tforce-io/tf-golib/main.go", &FileName{Name: "main", Extension: ".go"}},
		})
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			result := NewFileNameFromStr(tt.path)
			if !AreEqualFileNames(result, tt.expected) {
				t.Errorf("expected %v current %v", tt.expected, result)
			}
		})
	}
}

func TestFileName_FullName(t *testing.T) {
	tests := []struct {
		group    string
		name     string
		ext      string
		prefix   string
		suffix   string
		expected string
	}{
		{"name_only", "main", "", "", "", "main"},
		{"name_with_ext", "main", ".go", "", "", "main.go"},
		{"name_with_prefix", "main", "", "copy_of_", "", "copy_of_main"},
		{"name_with_prefix", "main", ".go", "copy_of_", "", "copy_of_main.go"},
		{"name_with_suffix", "main", "", "", "_original", "main_original"},
		{"name_with_suffix", "main", ".go", "", "_original", "main_original.go"},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			fileName := NewFileName(tt.name, tt.ext)
			fileName.Prefix = tt.prefix
			fileName.Suffix = tt.suffix
			fullName := fileName.FullName()
			if fullName != tt.expected {
				t.Errorf("expected %v current %v", tt.expected, fullName)
			}
		})
	}
}

func TestAreEqualFileName(t *testing.T) {
	tests := []struct {
		group    string
		x        *FileName
		y        *FileName
		expected bool
	}{
		// all_nils
		{"all_nils", nil, nil, true},
		// one_nil
		{"one_nil", &FileName{}, nil, false},
		{"one_nil", nil, &FileName{}, false},
		// deep_compare
		{"deep_compare", &FileName{}, &FileName{}, true},
		{"deep_compare", &FileName{Name: "main"}, &FileName{Name: "main"}, true},
		{"deep_compare", &FileName{Name: "main", Extension: ".go"}, &FileName{Name: "main", Extension: ".go"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			result := AreEqualFileNames(tt.x, tt.y)
			if result != tt.expected {
				t.FailNow()
			}
		})
	}
}
