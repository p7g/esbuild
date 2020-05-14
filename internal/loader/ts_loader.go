package loader

import (
	"github.com/evanw/esbuild/internal/ast"
	"github.com/evanw/esbuild/internal/logging"
	"github.com/evanw/esbuild/internal/parser"
)

type TSLoader struct{}

func (self TSLoader) Extension() string {
	return ".ts"
}

func (self TSLoader) ShouldLoad(filename string) bool {
	return FileExtension(filename) == ".ts"
}

func (self TSLoader) Parse(log logging.Log, source logging.Source, options parser.ParseOptions) (ast.AST, bool) {
	options.TS.Parse = true
	return parser.Parse(log, source, options)
}
