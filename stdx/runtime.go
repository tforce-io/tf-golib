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
	"runtime"
	"slices"
)

var _LINUX_FAMILIES = []string{"linux", "freebsd", "netbsd", "openbsd", "aix", "illumos", "plan9", "solaris"}
var _WEB_ASSEMBLY_FAMILIES = []string{"wasip1"}

func IsAndroid() bool {
	return runtime.GOOS == "android"
}

func IsIos() bool {
	return runtime.GOOS == "ios"
}

func IsLinux() bool {
	return slices.Contains(_LINUX_FAMILIES, runtime.GOOS)
}

func IsMacintosh() bool {
	return runtime.GOOS == "darwin"
}

func IsWebAssembly() bool {
	return slices.Contains(_WEB_ASSEMBLY_FAMILIES, runtime.GOOS)
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}
