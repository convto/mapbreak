package main

import (
	"subpkg/pkg"
)

func main() {
	for k, v := range pkg.M {
		pkg.M[k] = "reassigned: " + v // want "detected range access to map and reassigning it"
	}
}
