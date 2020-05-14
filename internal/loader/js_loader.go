package loader

import (
	"github.com/evanw/esbuild/internal/ast"
	"github.com/evanw/esbuild/internal/logging"
	"github.com/evanw/esbuild/internal/parser"
)

type JSLoader struct{}

func (self JSLoader) Extension() string {
	return ".js"
}

func (self JSLoader) ShouldLoad(filename string) bool {
	return FileExtension(filename) == ".js"
}

func (self JSLoader) Parse(log logging.Log, source logging.Source, options parser.ParseOptions) (ast.AST, bool) {
	return parser.Parse(log, source, options)
}
