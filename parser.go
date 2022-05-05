package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	gp "go/parser"
	"go/token"
	"go/types"
	"log"
)

type Parser struct {
}

type Interface struct {
}

type Methods struct {
}

var defsOrUses = map[*ast.Ident]types.Object{}

// methods[interface][method][arg,ret][name]type
var methods = map[string]string{}

func NewParser() *Parser {
	return &Parser{}
}

func (parser *Parser) Parse(text string) string {
	fset := token.NewFileSet()

	f, err := gp.ParseFile(fset, "", text, 0)
	if err != nil {
		log.Fatal(err)
	}

	ast.Print(fset, f)

	// defsOrUses := map[*ast.Ident]types.Object{}
	info := &types.Info{
		Defs: defsOrUses,
		Uses: defsOrUses,
	}

	config := &types.Config{
		Importer: importer.Default(),
	}

	_, err = config.Check("main", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatal("Error:", err)
	}

	ast.Inspect(f, Check)

	return "Gopher"
}

func Check(node ast.Node) bool {
	ident, ok := node.(*ast.Ident)
	if !ok {
		return true
	}

	obj := defsOrUses[ident]
	if obj == nil {
		return true
	}

	switch obj.Type().(type) {
	// case *types.Named:
	// 	if !types.IsInterface(obj.Type()) {
	// 		return true
	// 	}
	// 	interfaceInfo := obj.Type().Underlying().(*types.Interface)

	// 	for i := 0; i < interfaceInfo.NumMethods(); i++ {
	// 		method := typ.Method(i)
	// 	}
	case *types.Signature:
		signatureInfo := obj.Type().(*types.Signature)
		results := signatureInfo.Results()
		for i := 0; i < results.Len(); i++ {
			fmt.Println(results.At(i).Type().String())
		}

		params := signatureInfo.Params()
		for i := 0; i < params.Len(); i++ {
			fmt.Println(params.At(i).Type().String())
		}

		recv := signatureInfo.Recv().Type().(*types.Named)
		fmt.Println(recv.Obj().Id())

		fmt.Println()
	}

	return true
}
