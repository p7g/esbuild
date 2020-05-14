package loader

import (
	"strings"

	"github.com/evanw/esbuild/internal/ast"
	"github.com/evanw/esbuild/internal/logging"
	"github.com/evanw/esbuild/internal/parser"
)

type Loader interface {
	Extension() string
	ShouldLoad(filename string) bool
	Parse(logging.Log, logging.Source, parser.ParseOptions) (ast.AST, bool)
}

func DefaultLoaders() []Loader {
	return []Loader{
		JSLoader{},
		JSXLoader{},
		TSLoader{},
		TSXLoader{},
		JSONLoader{},
	}
}

func FileExtension(path string) string {
	if lastDot := strings.LastIndexByte(path, '.'); lastDot >= 0 {
		return path[lastDot:]
	}
	return ""
}
