# Semantic Version Bumper

A simple tool that reads the semantic version from a version file (default is `VERSION`) and bumps it up.

# Installation

* `go get github.com/daniellowtw//bumpsemver`
* `go install github.com/daniellowtw/bumpsemver`

# Usage

```
Usage: ./bumpsemver [major|minor|patch]
Increase semantic version in the FILE

  -d, --dry-run                  print new version without updating version FILE
  -f, --file string              path to a version FILE (default "./VERSION")
  -p, --go-package-file string   package name of the generated version file (default "main")
  -g, --go-version-file string   path to the go version file (default "./version.go")
  -m, --meta string              metadata for version
  -r, --pre-release string       pre-release string
  -u, --update-go                enable updating of go version file

```

# Examples

```
./bumpsemver [major|minor|patch] -f version.txt # updates the version.txt file
./bumpsemver major # creates or updates the VERSION file
./bumpsemver major -u # creates or update the VERSION file and create a version.go file
```

## Use it with go

main.go

```
package main

import "fmt"

func main() {
  fmt.Println(version)

}
```

1. `bumpsemver major -u`
2. `go build -o foo`
3. `./foo #should print 1.0.0`
4. `bumpsemver major -u`
5. `go build -o foo`
6. `./foo #should print 2.0.0`

## Use it with script

Makefile:
```
all:
  go build -ldflags "-X main.version=$(shell cat ./VERSION)"
```


