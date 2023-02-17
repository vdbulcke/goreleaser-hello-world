package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// GitCommit the current git commit
// will be injected during build
var GitCommit string

// Version
var Version string

// Date
var Date string

// BuiltBy
var BuiltBy string

// HumanVersion version with commit
var HumanVersion = fmt.Sprintf("%s-(%s)", Version, GitCommit)

var short bool

// ref: goreleaser
func buildVersion(version, commit, date, builtBy string) string {
	result := version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	result = fmt.Sprintf("%s\ngoos: %s\ngoarch: %s", result, runtime.GOOS, runtime.GOARCH)
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		result = fmt.Sprintf("%s\nmodule version: %s, checksum: %s", result, info.Main.Version, info.Main.Sum)
	}
	return result
}

func main() {

	fmt.Println(buildVersion(Version, GitCommit, Date, BuiltBy))

}
