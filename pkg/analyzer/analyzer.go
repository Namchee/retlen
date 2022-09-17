package analyzer

import (
	"flag"
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var (
	maxReturn   uint
	includeTest bool
)

var flags flag.FlagSet

func init() {
	flags.UintVar(&maxReturn, "maxReturn", 2, "maximum number of variables to be returned")
	flags.BoolVar(&includeTest, "includeTest", false, "include test files to analysis")
}

var Analyzer = &analysis.Analyzer{
	Name:  "retlen",
	Doc:   "Checks whether a function returns too many variables",
	Run:   run,
	Flags: flags,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(node ast.Node) bool {
			fun, ok := node.(*ast.FuncDecl)
			if !ok {
				return true
			}

			funcName := fun.Name.Name
			if !includeTest && strings.HasPrefix(funcName, "test") {
				return true
			}

			returnCount := 0
			if fun.Type.Results != nil {
				returnCount = len(fun.Type.Results.List)
			}

			if uint(returnCount) > maxReturn {
				msg := fmt.Sprintf("Function %s returns too many variables", funcName)
				pass.Reportf(node.Pos(), msg)
			}

			return true
		})
	}

	return nil, nil
}
