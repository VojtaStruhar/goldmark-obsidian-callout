package callout

import (
	"github.com/VojtaStruhar/goldmark-obsidian-callout/ast"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"

	calloutParser "github.com/VojtaStruhar/goldmark-obsidian-callout/parser"
)

type obsidianCalloutExtension struct {
}

// ObsidianCallout is an extension to render <details> and <summary> elements for Obsidian flavored markdown style
var ObsidianCallout = &obsidianCalloutExtension{}

func (f *obsidianCalloutExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithParagraphTransformers(
			util.Prioritized(calloutParser.NewCalloutParagraphTransformer(), 120),
		),
		parser.WithASTTransformers(
			util.Prioritized(calloutParser.NewCalloutAstTransformer(), 0),
		),
	)
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		// Priority is at 0, because we are rendering our own node kind. Anyone can override us.
		util.Prioritized(ast.NewCalloutHtmlRenderer(), 0),
	))
}
