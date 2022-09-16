package main

import (
	"github.com/Namchee/funclint/pkg/argslen"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(argslen.Analyzer)
}
