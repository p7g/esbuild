package loader

import (
	"github.com/evanw/esbuild/internal/ast"
	"github.com/evanw/esbuild/internal/logging"
	"github.com/evanw/esbuild/internal/parser"
)

type TSXLoader struct{}

func (self TSXLoader) Extension() string {
	return ".tsx"
}

func (self TSXLoader) ShouldLoad(filename string) bool {
	return FileExtension(filename) == ".tsx"
}

func (self TSXLoader) Parse(log logging.Log, source logging.Source, options parser.ParseOptions) (ast.AST, bool) {
	options.TS.Parse = true
	options.JSX.Parse = true
	return parser.Parse(log, source, options)
}
