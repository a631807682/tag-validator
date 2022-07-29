package parser

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

type StructVisitor struct {
	sts []*ast.StructType
}

func (v *StructVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.StructType:
		v.sts = append(v.sts, n)
	}

	return v
}

func ParseFile(fileSet *token.FileSet, filepath string) (*ast.File, []*ast.StructType, error) {
	parserMode := parser.ParseComments | parser.SkipObjectResolution
	astFile, err := parser.ParseFile(fileSet, filepath, nil, parserMode)
	if err != nil {
		return nil, nil, nil
	}

	v := &StructVisitor{}
	ast.Walk(v, astFile)
	return astFile, v.sts, nil
}

type TagFormatter interface {
	Match(*ast.StructType) (fields []*ast.Field)
	Format(*ast.Field)
}

func FormatTags(filename string, tagfmt TagFormatter) (res []byte, err error) {
	fset := token.NewFileSet()
	astFile, sts, err := ParseFile(fset, filename)
	if err != nil {
		return
	}

	for _, st := range sts {
		fields := tagfmt.Match(st)
		for _, field := range fields {
			tagfmt.Format(field)
		}
	}

	var buf bytes.Buffer
	err = format.Node(&buf, fset, astFile)
	if err != nil {
		return
	}

	res = buf.Bytes()
	return
}
