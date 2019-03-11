package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"

	"github.com/selesy/go-genutil/genutil"
	"github.com/sirupsen/logrus"
)

const (
	ErrorNoTargetMessage = "No targets were found in the specified packages"
)

var log = logrus.New()

var enumFilter = func(node ast.Node) bool {
	//We want declarations that have a var token and a value that's an
	//anonymous array of struct.  This method just needs to take care of
	//identifying:
	//
	//	var xxx = []struct{<at least one field>}{{<at least one instance>}}
	//
	// Linting this declaration will occur later
	valSpec, ok := varDecl(node)
	if !ok {
		return false
	}

	compLit, ok := varValue(valSpec)
	if !ok || len(compLit.Elts) < 1 {
		return false
	}

	anonArr, ok := compLit.Type.(*ast.ArrayType)
	if !ok {
		return false
	}

	anonStruct, ok := anonArr.Elt.(*ast.StructType)
	if !ok || anonStruct.Fields.NumFields() < 1 {
		return false
	}

	return true
}

func main() {
	log.Level = logrus.TraceLevel

	log.Trace("-> main()")
	targets, err := genutil.FilterAstNodesFromArgs(enumFilter)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if len(targets) < 1 {
		log.Info(ErrorNoTargetMessage)
		os.Exit(0)
	}

	log.Info("Count of variable declarations with []struct type: ", len(targets))
	log.Trace("main() ->")
}

//
//AST parsing
//

func varDecl(n ast.Node) (*ast.ValueSpec, bool) {
	g, ok := n.(*ast.GenDecl)
	if !ok || g.Tok != token.VAR || len(g.Specs) < 1 {
		return nil, false
	}

	if len(g.Specs) > 1 {
		log.Error("blah")
		return nil, false
	}

	s, ok := g.Specs[0].(*ast.ValueSpec)
	if !ok || len(s.Values) < 1 {
		return nil, false
	}

	if len(s.Values) > 1 {
		log.Error("Double blah")
		return nil, false
	}

	return s, true
}

//

func varName(valSpec *ast.ValueSpec) (string, bool) {
	if len(valSpec.Names) < 1 {
		return "", false
	}
	return valSpec.Names[0].Name, true
}

func varValue(valSpec *ast.ValueSpec) (*ast.CompositeLit, bool) {
	if len(valSpec.Values) < 1 {
		return nil, false
	}

	compLit, ok := valSpec.Values[0].(*ast.CompositeLit)
	if !ok {
		return nil, false
	}

	return compLit, true
}

//Helpers related to the anonymous struct defining the enumeration

func fields(s *ast.StructType) []*ast.Field {
	for _, f := range s.Fields.List {
		fmt.Println("     ArrayStruct field:", f)
		fmt.Println("          Tags:", f.Tag)
	}
	return s.Fields.List
}

func tags(f *ast.Field) {

}

// for _, field := range arrayStructType.Fields.List {
// 	fmt.Println("     ArrayStruct field:", field)
// 	fmt.Println("          Tags:", field.Tag)
// }

// for _, compLitElt := range compLit.Elts {
// 	fmt.Println("     CompositeLit element: ", compLitElt)
// 	fmt.Println("     CompositeLit element type: ", reflect.TypeOf(compLitElt))
// 	innerCompLit, ok := compLitElt.(*ast.CompositeLit)
// 	if !ok {
// 		return false
// 	}
// 	for _, innerElt := range innerCompLit.Elts {
// 		fmt.Println("     Inner element:", innerElt)
// 	}
// }
// fmt.Println(arrType)

// fmt.Println("     CompositeLit: ", compLit)
// fmt.Println("     CompositeLit type: ", compLit.Type)
// fmt.Println("     CompositeLit type: ", reflect.TypeOf(compLit.Type))
// fmt.Println("     ArrayType:", arrType)
// fmt.Println("     ArrayType element:", arrType.Elt)
//fmt.Println("     ", arrayStructType)
