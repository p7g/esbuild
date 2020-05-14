package loader

import (
	"github.com/evanw/esbuild/internal/ast"
	"github.com/evanw/esbuild/internal/logging"
	"github.com/evanw/esbuild/internal/parser"
)

type JSXLoader struct{}

func (self JSXLoader) Extension() string {
	return ".jsx"
}

func (self JSXLoader) ShouldLoad(filename string) bool {
	return FileExtension(filename) == ".jsx"
}

func (self JSXLoader) Parse(log logging.Log, source logging.Source, options parser.ParseOptions) (ast.AST, bool) {
	options.JSX.Parse = true
	return parser.Parse(log, source, options)
}
