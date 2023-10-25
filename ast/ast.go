package ast

import (
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// KindCallout is a NodeKind of the Callout node.
var KindCallout = gast.NewNodeKind("Callout")

// A Callout struct represents a table of Markdown(GFM) text.
type Callout struct {
	gast.BaseBlock
}

// Kind implements Node.Kind.
func (n *Callout) Kind() gast.NodeKind {
	return KindCallout
}

// Dump implements Node.Dump
func (n *Callout) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, func(level int) {
	})
}

// NewCallout returns a new Table node.
func NewCallout() *Callout {
	return &Callout{}
}

// KindCalloutTitle is a NodeKind of the CalloutTitle node.
var KindCalloutTitle = gast.NewNodeKind("CalloutTitle")

// A CalloutTitle struct represents a table of Markdown(GFM) text.
type CalloutTitle struct {
	gast.BaseBlock
}

// Kind implements Node.Kind.
func (n *CalloutTitle) Kind() gast.NodeKind {
	return KindCalloutTitle
}

// Dump implements Node.Dump
func (n *CalloutTitle) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, func(level int) {})
}

// NewCalloutTitle returns a new Table node.
func NewCalloutTitle() *CalloutTitle {
	return &CalloutTitle{}
}

// CalloutHtmlRenderer is a renderer.NodeRenderer implementation that
// renders Callout nodes.
type CalloutHtmlRenderer struct {
}

// NewCalloutHtmlRenderer returns a new CalloutHtmlRenderer.
func NewCalloutHtmlRenderer() renderer.NodeRenderer {
	return &CalloutHtmlRenderer{}
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *CalloutHtmlRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindCallout, r.renderCallout)
	reg.Register(KindCalloutTitle, r.renderCalloutTitle)
}

func (r *CalloutHtmlRenderer) renderCallout(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<details class=\"callout\" data-callout=\"info\">\n")
	} else {
		_, _ = w.WriteString("</details>\n")
	}
	return gast.WalkContinue, nil
}

func (r *CalloutHtmlRenderer) renderCalloutTitle(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<summary>\n<p>\n")
	} else {
		_, _ = w.WriteString("</p>\n</summary>\n")
	}
	return gast.WalkContinue, nil
}
