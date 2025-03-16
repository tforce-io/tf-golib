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

	"github.com/tforce-io/tf-golib/multiarch"
	"github.com/tforce-io/tf-golib/stdx/stringxt"
)

var PATH_SEPARATOR = string(os.PathSeparator)

// A FileName is struct contains name and extension of a file or folder.
// Extension should have dot at the beginning.
// Available since v0.2.0
type FileName struct {
	Prefix    string
	Name      string
	Suffix    string
	Extension string
}

// Create a new FileName from scratch
// Available since v0.2.0
func NewFileName(name, extension string) *FileName {
	return &FileName{
		Name:      name,
		Extension: extension,
	}
}

// Parse path string to create a new FileName
// Available since v0.2.0
func NewFileNameFromStr(path string) *FileName {
	nPath := NormalizePath(path)
	base := filepath.Base(nPath)
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
// Available since v0.2.0
func (s *FileName) Clone() *FileName {
	return &FileName{
		Prefix:    s.Prefix,
		Name:      s.Name,
		Suffix:    s.Suffix,
		Extension: s.Extension,
	}
}

// Return full name represented by this FileName.
// Available since v0.2.0
func (s *FileName) FullName() string {
	return s.Prefix + s.Name + s.Suffix + s.Extension
}

// Check whether two FileNames are equal.
// Available since v0.2.0
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

// A Path is a struct contains all smallest components of a path.
// Available since v0.2.0
type Path struct {
	Parents []string
	Name    *FileName
}

// Create a new Path from scratch
// Available since v0.2.0
func NewPath(dirs []string, name *FileName) *Path {
	return &Path{
		Parents: dirs,
		Name:    name,
	}
}

// Parse path string to create a new Path
// Available since v0.2.0
func NewPathFromStr(path string) *Path {
	nPath := NormalizePath(path)
	dir, file := filepath.Split(nPath)
	dir = strings.TrimSuffix(dir, PATH_SEPARATOR)
	dirs := ternarySlice(stringxt.IsEmpty(dir), []string{}, strings.Split(dir, PATH_SEPARATOR))
	name := NewFileNameFromStr(file)
	return &Path{
		Parents: dirs,
		Name:    name,
	}
}

// Make a deep copy of this Path.
// Available since v0.2.0
func (s *Path) Clone() *Path {
	directories := make([]string, len(s.Parents))
	_ = copy(s.Parents, directories)
	name := s.Name.Clone()
	return &Path{
		Parents: directories,
		Name:    name,
	}
}

// Check a Path is whether asbsolute path. Using the same rule as
// filepath.IsAbs
// Available since v0.2.0
func (s *Path) IsAbsolute() bool {
	if isEmptySlice(s.Parents) {
		return false
	}
	fullPath := s.FullPath()
	return filepath.IsAbs(fullPath)
}

// Return full path represented by this Path.
// Available since v0.2.0
func (s *Path) FullPath() string {
	if isEmptySlice(s.Parents) {
		return s.Name.FullName()
	}
	return s.ParentPath() + PATH_SEPARATOR + s.Name.FullName()
}

// Return parent path represented by this Path.
// Available since v0.2.0
func (s *Path) ParentPath() string {
	return ternary(isEmptySlice(s.Parents), "", strings.Join(s.Parents, PATH_SEPARATOR))
}

// Check whether two Paths are equal.
// Available since v0.2.0
func AreEqualPaths(x, y *Path) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return areEqualSlices(x.Parents, y.Parents) &&
		AreEqualFileNames(x.Name, y.Name)
}

// Clean path and make path consistent accross platforms. This function will perform:
// - Replace all slashes to backslashes if run on Windows.
// - Replace all backslashes to slashes if run on all UNIX-like OSes.
// - Clean the path.
// Available since v0.2.0
func NormalizePath(path string) string {
	nPath := ternary(multiarch.IsWindows(),
		strings.ReplaceAll(path, "/", "\\"),
		strings.ReplaceAll(path, "\\", "/"),
	)
	return filepath.Clean(nPath)
}

// forked from opx package
func areEqualSlices(x, y []string) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// forked from opx package
func isEmptySlice(slice []string) bool {
	if slice == nil {
		return true
	}
	return len(slice) == 0
}

// forked from opx package
func ternary(cond bool, x, y string) string {
	if cond {
		return x
	}
	return y
}

// forked from opx package
func ternarySlice(cond bool, x, y []string) []string {
	if cond {
		return x
	}
	return y
}
