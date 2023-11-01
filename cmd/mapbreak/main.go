package main

import (
	"github.com/convto/mapbreak"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(mapbreak.Analyzer) }