package analyzer

import (
	"flag"
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
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
	Name:     "retlen",
	Doc:      "Checks whether a function returns too many variables",
	Run:      run,
	Flags:    flags,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		fun, _ := node.(*ast.FuncDecl)

		funcName := fun.Name.Name
		if !includeTest && strings.HasPrefix(funcName, "Test") {
			return
		}

		returnCount := 0
		if fun.Type.Results != nil {
			returnCount = len(fun.Type.Results.List)
		}

		if uint(returnCount) > maxReturn {
			msg := fmt.Sprintf("Function %s returns too many variables", funcName)
			pass.Reportf(node.Pos(), msg)
		}
	})

	return nil, nil
}
