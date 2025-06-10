// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package multiarch

import (
	"runtime"
)

var linuxFamily = []string{"linux", "freebsd", "netbsd", "openbsd", "aix", "illumos", "plan9", "solaris"}
var wasmFamily = []string{"wasip1"}

// Check whether current platform is Android.
// Available since v0.3.0
func IsAndroid() bool {
	return runtime.GOOS == "android"
}

// Check whether current platform is iOS.
// Available since v0.3.0
func IsIos() bool {
	return runtime.GOOS == "ios"
}

// Check whether current platform is Linux.
// "linux", "freebsd", "netbsd", "openbsd", "aix", "illumos", "plan9", "solaris" are considered Linux.
// Available since v0.3.0
func IsLinux() bool {
	for _, os := range linuxFamily {
		if os == runtime.GOOS {
			return true
		}
	}
	return false
}

// Check whether current platform is MacOSX.
// Available since v0.3.0
func IsMacintosh() bool {
	return runtime.GOOS == "darwin"
}

// Check whether current platform is Web Assembly.
// Available since v0.3.0
func IsWebAssembly() bool {
	for _, os := range wasmFamily {
		if os == runtime.GOOS {
			return true
		}
	}
	return false
}

// Check whether current platform is Windows.
// Available since v0.3.0
func IsWindows() bool {
	return runtime.GOOS == "windows"
}
