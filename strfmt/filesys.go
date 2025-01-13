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
	"os"
	"path/filepath"
	"strings"
)

var PATH_SEPARATOR = string(os.PathSeparator)

// A FileName is struct contains name and extension of a file or folder.
// Extension should have dot at the beginning.
type FileName struct {
	Prefix    string
	Name      string
	Suffix    string
	Extension string
}

func NewFileName(name, extension string) *FileName {
	return &FileName{
		Name:      name,
		Extension: extension,
	}
}

func NewFileNameFromStr(path string) *FileName {
	base := filepath.Base(path)
	if base == "." || base == PATH_SEPARATOR {
		return &FileName{
			Name:      "",
			Extension: "",
		}
	}

	ext := filepath.Ext(base)
	return &FileName{
		Name:      strings.TrimSuffix(base, ext),
		Extension: ext,
	}
}

// Make a deep copy of this Path.
func (s *FileName) Clone() *FileName {
	return &FileName{
		Prefix:    s.Prefix,
		Name:      s.Name,
		Suffix:    s.Suffix,
		Extension: s.Extension,
	}
}

// Returns full name represented by this FileName.
func (s *FileName) FullName() string {
	return s.Prefix + s.Name + s.Suffix + s.Extension
}

func AreEqualFileNames(x, y *FileName) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return x.Name == y.Name &&
		x.Extension == y.Extension &&
		x.Prefix == y.Prefix &&
		x.Suffix == y.Suffix
}
