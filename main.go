package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/coreos/go-semver/semver"
	flag "github.com/spf13/pflag"
)

const template = `// This is generated code. Do not edit.
package %s

var version = "%s"
`

var (
	f            = flag.NewFlagSet("", flag.ExitOnError)
	fVersionFile = f.StringP("file", "f", "./VERSION", "path to a version FILE")
	fDryRun      = f.BoolP("dry-run", "d", false, "print new version without updating version FILE")

	fGoVersionFile       = f.StringP("go-version-file", "g", "./version.go", "path to the go version file")
	fPackageName         = f.StringP("go-package-file", "p", "main", "package name of the generated version file")
	fEnableGoVersionFile = f.BoolP("update-go", "u", false, "enable updating of go version file")

	fMeta             = f.StringP("meta", "m", "", "metadata for version")
	fPreReleaseString = f.StringP("pre-release", "r", "", "pre-release string")
)

func main() {
	f.Usage = usages
	err := f.Parse(os.Args[1:])
	args := f.Args()
	if len(args) == 0 {
		usages()
		os.Exit(1)
	}
	op := f.Args()[0]
	v, err := getVersionFromFile()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	v.Metadata = *fMeta
	v.PreRelease = semver.PreRelease(*fPreReleaseString)

	switch op {
	case "major":
		v.BumpMajor()
	case "minor":
		v.BumpMinor()
	case "patch":
		v.BumpPatch()
	case "update":
		updateGoVersionFile(v.String())
		return
	}

	if *fDryRun {
		os.Stdout.WriteString(v.String())
		return
	}
	if *fEnableGoVersionFile {
		updateGoVersionFile(v.String())
	}
	ioutil.WriteFile(*fVersionFile, []byte(v.String()), 0644)
}

func updateGoVersionFile(versionString string) {
	ioutil.WriteFile(*fGoVersionFile, []byte(fmt.Sprintf(template, *fPackageName, versionString)), 0644)
}

func usages() {
	os.Stdout.WriteString(fmt.Sprintf("Usage: %s [major|minor|patch]\nIncrease semantic version in the FILE\n\n%s", os.Args[0], f.FlagUsages()))
}

func getVersionFromFile() (*semver.Version, error) {
	data, err := ioutil.ReadFile(*fVersionFile)
	if err != nil {
		return semver.NewVersion("0.0.0")
	}
	return semver.NewVersion(string(data))
}
