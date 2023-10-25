package callout_test

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"

	callout "github.com/VojtaStruhar/goldmark-callout"
)

func TestBlockquote(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			callout.ObsidianCallout,
		),
	)
	count := 0

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Default blockquote",
		Markdown: `
> This is a blockquote
`,
		Expected: `
<blockquote>
<p>This is a blockquote</p>
</blockquote>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Default blockquote",
		Markdown: `
> [!info] This is a blockquote
`,
		Expected: `
<blockquote>
<p>This is a blockquote</p>
</blockquote>
`,
	}, t)
}
