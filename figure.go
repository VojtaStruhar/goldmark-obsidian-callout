package callout

import (
	"github.com/mangoumbrella/goldmark-figure/ast"
	"github.com/yuin/goldmark"
	aparser "github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"

	fparser "github.com/mangoumbrella/goldmark-figure/parser"
)

type extension struct {
}

// Figure is an extension to render <figure> elements.
var Figure = &extension{}

func (f *extension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		aparser.WithParagraphTransformers(
			util.Prioritized(fparser.NewFigureParagraphTransformer(), 120),
		),
		aparser.WithASTTransformers(
			util.Prioritized(fparser.NewFigureASTTransformer(), 0),
		),
	)
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(ast.NewFigureHTMLRenderer(), 0),
	))
}
