package ast

import (
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// KindFigure is a NodeKind of the Figure node.
var KindFigure = gast.NewNodeKind("Figure")

// A Figure struct represents a table of Markdown(GFM) text.
type Figure struct {
	gast.BaseBlock
}

// Kind implements Node.Kind.
func (n *Figure) Kind() gast.NodeKind {
	return KindFigure
}

// Dump implements Node.Dump
func (n *Figure) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, func(level int) {
	})
}

// NewFigure returns a new Table node.
func NewFigure() *Figure {
	return &Figure{}
}

// KindFigureImage is a NodeKind of the FigureImage node.
var KindFigureImage = gast.NewNodeKind("FigureImage")

// A FigureImage struct represents a table of Markdown(GFM) text.
type FigureImage struct {
	gast.BaseBlock
}

// Kind implements Node.Kind.
func (n *FigureImage) Kind() gast.NodeKind {
	return KindFigureImage
}

// Dump implements Node.Dump
func (n *FigureImage) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, func(level int) {
	})
}

// NewFigureImage returns a new Table node.
func NewFigureImage() *FigureImage {
	return &FigureImage{}
}

// KindFigureCaption is a NodeKind of the FigureCaption node.
var KindFigureCaption = gast.NewNodeKind("FigureCaption")

// A FigureCaption struct represents a table of Markdown(GFM) text.
type FigureCaption struct {
	gast.BaseBlock
}

// Kind implements Node.Kind.
func (n *FigureCaption) Kind() gast.NodeKind {
	return KindFigureCaption
}

// Dump implements Node.Dump
func (n *FigureCaption) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, func(level int) {
	})
}

// NewFigureCaption returns a new Table node.
func NewFigureCaption() *FigureCaption {
	return &FigureCaption{}
}

// FigureHTMLRenderer is a renderer.NodeRenderer implementation that
// renders Figure nodes.
type FigureHTMLRenderer struct {
}

// NewFigureHTMLRenderer returns a new FigureHTMLRenderer.
func NewFigureHTMLRenderer() renderer.NodeRenderer {
	return &FigureHTMLRenderer{}
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *FigureHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindFigure, r.renderFigure)
	reg.Register(KindFigureCaption, r.renderFigureCaption)
}

func (r *FigureHTMLRenderer) renderFigure(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<figure>\n")
	} else {
		_, _ = w.WriteString("</figure>\n")
	}
	return gast.WalkContinue, nil
}

func (r *FigureHTMLRenderer) renderFigureCaption(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<figcaption><p>")
	} else {
		_, _ = w.WriteString("</p></figcaption>\n")
	}
	return gast.WalkContinue, nil
}
