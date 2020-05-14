package loader

import (
	"github.com/evanw/esbuild/internal/ast"
	"github.com/evanw/esbuild/internal/logging"
	"github.com/evanw/esbuild/internal/parser"
)

type JSONLoader struct{}

func (self JSONLoader) Extension() string {
	return ".json"
}

func (self JSONLoader) ShouldLoad(filename string) bool {
	return FileExtension(filename) == ".json"
}

func (self JSONLoader) Parse(log logging.Log, source logging.Source, options parser.ParseOptions) (ast.AST, bool) {
	expr, ok := parser.ParseJSON(log, source, parser.ParseJSONOptions{})
	ast := parser.ModuleExportsAST(log, source, options, expr)
	return ast, ok
}
