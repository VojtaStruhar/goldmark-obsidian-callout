package parser

import (
	"fmt"
	"regexp"

	fast "github.com/mangoumbrella/goldmark-figure/ast"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var imageRegexp = regexp.MustCompile(`^\[!(\w+)\]([+-])?(.*)\n$`)

type figureParagraphTransformer struct {
}

var defaultFigureParagraphTransformer = &figureParagraphTransformer{}

// NewFigureParagraphTransformer returns a new ParagraphTransformer
// that can transform paragraphs into figures.
func NewFigureParagraphTransformer() parser.ParagraphTransformer {
	return defaultFigureParagraphTransformer
}

func (b *figureParagraphTransformer) Transform(node *gast.Paragraph, reader text.Reader, pc parser.Context) {
	lines := node.Lines()
	if lines.Len() < 1 {
		return
	}
	var first_seg = lines.At(0)
	var first_line_str = first_seg.Value(reader.Source())
	// Here we simply match by regex.
	// But this simple regex ignores image descriptions that contain other links.
	// E.g. ![foo ![bar](/url)](/url2).
	// See CommonMark spec: https://spec.commonmark.org/0.30/#images.
	if !imageRegexp.Match(first_line_str) {
		return
	}
	fmt.Println("We got a match!", string(first_line_str))
	figure := fast.NewFigure()
	node.Parent().ReplaceChild(node.Parent(), node, figure)

	figureImage := fast.NewFigureImage()
	figureImage.Lines().Append(lines.At(0))
	figure.AppendChild(figure, figureImage)

	if lines.Len() >= 2 {
		figureCaption := fast.NewFigureCaption()
		for i := 1; i < lines.Len(); i++ {
			seg := lines.At(i)
			if i == lines.Len()-1 {
				// trim last newline(\n)
				seg.Stop = seg.Stop - 1
			}
			figureCaption.Lines().Append(seg)
		}
		figure.AppendChild(figure, figureCaption)
	}
}

type figureASTTransformer struct {
}

var defaultFigureASTTransformer = &figureASTTransformer{}

// NewFigureASTTransformer returns a parser.ASTTransformer for tables.
func NewFigureASTTransformer() parser.ASTTransformer {
	return defaultFigureASTTransformer
}

func (a *figureASTTransformer) Transform(node *gast.Document, reader text.Reader, pc parser.Context) {
}
