package main

import "fmt"

// GitCommit the current git commit
// will be injected during build
var GitCommit string

// Version
var version string

// HumanVersion version with commit
var HumanVersion = fmt.Sprintf("%s-(%s)", version, GitCommit)

func main() {

	fmt.Printf("Hello world %s\n", HumanVersion)

}
