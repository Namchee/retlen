package main

import (
	"github.com/Namchee/argslen/pkg/argslen"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(argslen.Analyzer)
}
