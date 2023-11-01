package mapbreak

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "mapbreak",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "mapbreak detects if there is map reassignment in the range access"

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.RangeStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		rstmt, ok := n.(*ast.RangeStmt)
		if !ok {
			return
		}

		rangeTarget := getObject(pass, rstmt.X)
		if rangeTarget == nil {
			return
		}
		_, ok = deref(rangeTarget.Type()).Underlying().(*types.Map)
		if !ok {
			return
		}

		for _, stmt := range rstmt.Body.List {
			assign, ok := stmt.(*ast.AssignStmt)
			if !ok {
				continue
			}
			idx, ok := assign.Lhs[0].(*ast.IndexExpr)
			if !ok {
				continue
			}
			assignee := getObject(pass, idx.X)
			if assignee == rangeTarget {
				pass.Reportf(stmt.Pos(), "detected range access to map and reassigning it")
			}
		}
	})

	return nil, nil
}

func deref(typ types.Type) types.Type {
	if ptr, ok := typ.Underlying().(*types.Pointer); ok {
		return deref(ptr.Elem())
	}
	return typ
}

func getObject(pass *analysis.Pass, x ast.Expr) types.Object {
	switch x.(type) {
	case *ast.Ident:
		return pass.TypesInfo.ObjectOf(x.(*ast.Ident))
	case *ast.StarExpr:
		return getObject(pass, x.(*ast.StarExpr).X)
	case *ast.ParenExpr:
		return getObject(pass, x.(*ast.ParenExpr).X)
	case *ast.SelectorExpr:
		sel := x.(*ast.SelectorExpr)
		id, ok := sel.X.(*ast.Ident)
		if !ok {
			return nil
		}
		pkg := id.Name
		call := sel.Sel.Name
		imports := pass.Pkg.Imports()
		for i := range imports {
			if pkg == imports[i].Name() {
				return imports[i].Scope().Lookup(call)
			}
		}
	}
	return nil
}
