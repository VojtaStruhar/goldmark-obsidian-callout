package parser

import (
	"regexp"

	callout_ast "github.com/VojtaStruhar/goldmark-callout/ast"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var calloutRegex = regexp.MustCompile(`^\[!(\w+)\]([+-])?(.*)\n$`)

type calloutParagraphTransformer struct {
}

var defaultCalloutTransformer = &calloutParagraphTransformer{}

// NewCalloutParagraphTransformer returns a new ParagraphTransformer
// that can transform paragraphs into figures.
func NewCalloutParagraphTransformer() parser.ParagraphTransformer {
	return defaultCalloutTransformer
}

func (b *calloutParagraphTransformer) Transform(node *gast.Paragraph, reader text.Reader, pc parser.Context) {
	lines := node.Lines()
	if lines.Len() < 1 {
		return
	}
	var firstSegment = lines.At(0)
	var firstLineString = firstSegment.Value(reader.Source())

	if !calloutRegex.Match(firstLineString) {
		return
	}

	callout := callout_ast.NewCallout()
	node.Parent().ReplaceChild(node.Parent(), node, callout)

	figureImage := callout_ast.NewCalloutTitle()
	figureImage.Lines().Append(lines.At(0))
	callout.AppendChild(callout, figureImage)

	if lines.Len() >= 2 {
		figureCaption := callout_ast.NewCalloutTitle()
		for i := 1; i < lines.Len(); i++ {
			seg := lines.At(i)
			if i == lines.Len()-1 {
				// trim last newline(\n)
				seg.Stop = seg.Stop - 1
			}
			figureCaption.Lines().Append(seg)
		}
		callout.AppendChild(callout, figureCaption)
	}
}

type calloutAstTransformer struct {
}

var defaultCalloutAstTransformer = &calloutAstTransformer{}

// NewCalloutAstTransformer returns a parser.ASTTransformer for tables.
func NewCalloutAstTransformer() parser.ASTTransformer {
	return defaultCalloutAstTransformer
}

func (a *calloutAstTransformer) Transform(node *gast.Document, reader text.Reader, pc parser.Context) {
}
