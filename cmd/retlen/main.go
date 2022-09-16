package main

import (
	"github.com/Namchee/funclint/pkg/retlen"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(retlen.Analyzer)
}
