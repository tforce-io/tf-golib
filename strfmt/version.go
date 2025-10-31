// Copyright (C) 2025 T-Force I/O
//
// TF GoLib is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package strfmt

import (
	"fmt"
	"regexp"
	"strconv"
)

// Version represents a struct that followed the rule of semantic version.
// For more information, please refer to https://semver.org/
// For strict semantic version compliant, please keep Revision as 0.
//
// Available since v0.8.0
type Version struct {
	Major    uint64
	Minor    uint64
	Patch    uint64
	Revision uint64

	PreRelease string
	BuildMeta  string
}

// Returns a new version instance.
//
// Available since v0.8.0
func NewVersion(major, minor, patch, revision uint64, preRelease, buildMeta string) Version {
	return Version{
		Major:    major,
		Minor:    minor,
		Patch:    patch,
		Revision: revision,

		PreRelease: preRelease,
		BuildMeta:  buildMeta,
	}
}

// Parses a string and returns a new version instance.
//
// Available since v0.8.0
func FromString(versionStr string) Version {
	semVerRegEx := regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:\.(?P<revision>[1-9]\d*))?(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	matches := semVerRegEx.FindStringSubmatch(versionStr)

	if len(matches) == 0 {
		return Version{}
	}

	major, _ := strconv.ParseUint(matches[semVerRegEx.SubexpIndex("major")], 10, 0)
	minor, _ := strconv.ParseUint(matches[semVerRegEx.SubexpIndex("minor")], 10, 0)
	patch, _ := strconv.ParseUint(matches[semVerRegEx.SubexpIndex("patch")], 10, 0)
	revision, _ := strconv.ParseUint(matches[semVerRegEx.SubexpIndex("revision")], 10, 0)
	preRelease := matches[semVerRegEx.SubexpIndex("prerelease")]
	buildMeta := matches[semVerRegEx.SubexpIndex("buildmetadata")]
	return Version{
		Major:    major,
		Minor:    minor,
		Patch:    patch,
		Revision: revision,

		PreRelease: preRelease,
		BuildMeta:  buildMeta,
	}
}

// Checks current Version is compliant with semantic versioning specification.
//
// Available since v0.8.0
func (v Version) IsValid() bool {
	prRegex := `^(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?$`
	bmRegex := `^(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*)?$`
	r1, _ := regexp.MatchString(prRegex, v.PreRelease)
	r2, _ := regexp.MatchString(bmRegex, v.BuildMeta)
	return r1 && r2
}

// Returns a string represents semantic version.
//
// Available since v0.8.0
func (v Version) String() string {
	version := ""
	if v.Revision == 0 {
		version = fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	} else {
		version = fmt.Sprintf("%d.%d.%d.%d", v.Major, v.Minor, v.Patch, v.Revision)
	}
	if v.PreRelease != "" {
		version += "-" + v.PreRelease
	}
	if v.BuildMeta != "" {
		version += "+" + v.BuildMeta
	}
	return version
}
