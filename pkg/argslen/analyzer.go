package argslen

import (
	"errors"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "argslen",
	Doc:  "Checks whether a function has too many arguments",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, errors.New("unimplemented")
}
