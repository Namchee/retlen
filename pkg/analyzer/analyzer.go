package analyzer

import (
	"errors"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "retlen",
	Doc:  "Checks whether a function returns too many variables",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, errors.New("unimplemented")
}
