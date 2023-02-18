package evaluator

import (
	"kinolang/ast"
	"kinolang/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}