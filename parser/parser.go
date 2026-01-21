package parser

import (
	"regexp"

	calloutAst "github.com/VojtaStruhar/goldmark-obsidian-callout/ast"
	"github.com/VojtaStruhar/goldmark-obsidian-callout/helper"
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
	var firstLineBytes = firstSegment.Value(reader.Source())

	// Greedy!
	// If previous sibling is Callout, append this paragraph to the callout
	if node.PreviousSibling() != nil && node.PreviousSibling().Kind() == calloutAst.KindCallout {
		callout := node.PreviousSibling()
		node.Parent().RemoveChild(node.Parent(), node)
		callout.AppendChild(callout, node)
		return
	}

	// If this paragraph begins with the [!callout] tag
	if calloutRegex.Match(firstLineBytes) {

		callout := calloutAst.NewCallout()
		calloutTitle := calloutAst.NewCalloutTitle()

		closingBracketIndex, err := helper.IndexOf(firstLineBytes, byte(']'))
		if err != nil {
			return
		}
		openingBracketIndex, err := helper.IndexOf(firstLineBytes, byte('['))
		if err != nil {
			return
		}

		node.Parent().ReplaceChild(node.Parent(), node, callout)

		titleTextSegment := lines.At(0)
		titleTextSegment.Start += closingBracketIndex + 1

		{ // Type of the callout
			cName := string(firstLineBytes[openingBracketIndex+1 : closingBracketIndex])
			cType := helper.CalloutTypeMapping[cName]
			callout.SetAttribute([]byte("type"), cType)
		}

		{ // Determine the open-close mode of the callout
			calloutMode := helper.ForceOpen // default
			if len(firstLineBytes) > closingBracketIndex+1 {
				letterAfterClosingBracket := firstLineBytes[closingBracketIndex+1] // symbol right after the [!callout]

				if letterAfterClosingBracket == byte('+') {
					calloutMode = helper.OpenByDefault
					titleTextSegment.Start += 1
				} else if letterAfterClosingBracket == byte('-') {
					calloutMode = helper.ClosedByDefault
					titleTextSegment.Start += 1
				}
			}
			callout.SetAttribute([]byte("mode"), calloutMode)
		}

		// TODO: Rather than leaving the title text empty, supply a capitalized callout type
		calloutTitle.Lines().Append(titleTextSegment)

		callout.AppendChild(callout, calloutTitle)

		// If the callout has some content
		if lines.Len() >= 2 {
			calloutContent := gast.NewParagraph()
			for i := 1; i < lines.Len(); i++ {
				seg := lines.At(i)
				if i == lines.Len()-1 {
					// trim last newline (\n)
					seg.Stop = seg.Stop - 1
				}
				calloutContent.Lines().Append(seg)
			}
			callout.AppendChild(callout, calloutContent)
		}
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

	// Recursively walk every node and replace any blockquote whose first child
	// is a Callout with the Callout node itself. When replacing, move the
	// remaining children of the blockquote into the callout so nested callouts
	// (>> style) are preserved.
	var replace func(parent gast.Node)
	replace = func(parent gast.Node) {
		for current := parent.FirstChild(); current != nil; {
			next := current.NextSibling()

			if current.Kind() == gast.KindBlockquote {
				firstChild := current.FirstChild()
				if firstChild != nil && firstChild.Kind() == calloutAst.KindCallout {
					// Move all siblings (children of the blockquote) after the first child
					// into the callout so their content is not dropped when we replace
					// the blockquote with the callout.
					for child := firstChild.NextSibling(); child != nil; {
						childNext := child.NextSibling()
						// detach child from the blockquote
						current.RemoveChild(current, child)
						// append it to the callout
						firstChild.AppendChild(firstChild, child)
						child = childNext
					}
					// replace the blockquote with the callout (firstChild)
					parent.ReplaceChild(parent, current, firstChild)
					// continue traversal inside the moved callout
					replace(firstChild)
				} else {
					// Recurse into the blockquote to find nested callouts deeper down
					replace(current)
				}
			} else {
				// Recurse into other node types as well to find blockquotes containing callouts at any depth
				replace(current)
			}

			current = next
		}
	}

	replace(node)
}
