package parser

import (
	"regexp"

	calloutAst "github.com/VojtaStruhar/goldmark-callout/ast"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var calloutRegex = regexp.MustCompile(`^\[!(\w+)\]([+-])?(.*)\n$`)

type calloutType int

const (
	Info calloutType = iota
	Important
)

var calloutTypeMapping = map[string]calloutType{
	"!info":      Info,
	"!important": Important,
	"!tip":       Important,
}

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
	var firstLineBytes = firstSegment.Value(reader.Source())

	if !calloutRegex.Match(firstLineBytes) {
		return
	}

	callout := calloutAst.NewCallout()

	closingBracketIndex, err := indexOf(firstLineBytes, byte(']'))
	if err != nil {
		return
	}
	openingBracketIndex, err := indexOf(firstLineBytes, byte('['))
	if err != nil {
		return
	}

	cName := string(firstLineBytes[openingBracketIndex+1 : closingBracketIndex])
	cType := calloutTypeMapping[cName]
	callout.SetAttribute([]byte("type"), cType)

	node.Parent().ReplaceChild(node.Parent(), node, callout)

	calloutTitle := calloutAst.NewCalloutTitle()
	titleText := lines.At(0)
	// TODO: handle "+- " after the [!callout_type]
	shift := closingBracketIndex + 1

	titleText.Start += shift
	calloutTitle.Lines().Append(titleText)

	callout.AppendChild(callout, calloutTitle)

	if lines.Len() >= 2 {
		figureCaption := calloutAst.NewCalloutTitle()
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

	current := node.FirstChild()
	// TODO: Extract the walking-replacing into a function to allow nested callouts (not used often..)
	// check if current is of type gast.BlockQuote
	for current != nil {
		if current.Kind() == gast.KindBlockquote {
			// check if the blockquote has a child of type callout
			// if yes, then remove the blockquote and replace it with a callout
			if current.FirstChild().Kind() == calloutAst.KindCallout {
				// replace the blockquote with the callout
				node.ReplaceChild(node, current, current.FirstChild())
			}
		}
		current = current.NextSibling()
	}
}
