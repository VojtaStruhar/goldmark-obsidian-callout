package callout_test

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"

	callout "github.com/VojtaStruhar/goldmark-callout"
)

var markdown = goldmark.New(
	goldmark.WithExtensions(
		callout.ObsidianCallout,
	),
)
var count = 0

func TestAll(t *testing.T) {
	t.Run("Default Blockquote Behavior", TestBlockquote)
	t.Run("Empty Callout", TestEmptyCallouts)
	t.Run("Callouts with content", TestCalloutContent)
	fmt.Println("Ran", count, "tests total")
}

func TestBlockquote(t *testing.T) {

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
}

func TestEmptyCallouts(t *testing.T) {
	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Callout With custom title",
		Markdown: `
> [!info] Custom callout title
`,
		Expected: `<details class="callout" data-callout="info">
<summary>
<p>
 Custom callout title</p>
</summary>
</details>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Empty callout: important = tip alias",
		Markdown: `
> [!important]
`,
		Expected: `<details class="callout" data-callout="tip">
<summary>
<p>
</p>
</summary>
</details>
`,
	}, t)
}
func TestCalloutContent(t *testing.T) {
	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Callout with paragraph content",
		Markdown: `
> [!info]
> Some content here
`,
		Expected: `<details class="callout" data-callout="info">
<summary>
<p>
</p>
</summary>
<p>Some content here</p>
</details>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Callout with formatted paragraph content",
		Markdown: `
> [!info]
> **Bold** and *emphasis* formatting still works
`,
		Expected: `<details class="callout" data-callout="info">
<summary>
<p>
</p>
</summary>
<p><strong>Bold</strong> and <em>emphasis</em> formatting still works</p>
</details>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Callout with multiple paragraphs",
		Markdown: `
> [!info]
> More paragraphs
> 
> In a single callout
`,
		Expected: `<details class="callout" data-callout="info">
<summary>
<p>
</p>
</summary>
<p>More paragraphs</p>
<p>In a single callout</p>
</details>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Callout with multiple paragraphs and a custom title",
		Markdown: `
> [!example] Some title
> More paragraphs
> 
> In a single callout
`,
		Expected: `<details class="callout" data-callout="example">
<summary>
<p>
 Some title</p>
</summary>
<p>More paragraphs</p>
<p>In a single callout</p>
</details>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Callout with multiple paragraphs, custom title and a blank line under the title",
		Markdown: `
> [!example] Some title
>
> More paragraphs
> 
> In a single callout
`,
		Expected: `<details class="callout" data-callout="example">
<summary>
<p>
 Some title</p>
</summary>
<p>More paragraphs</p>
<p>In a single callout</p>
</details>
`,
	}, t)
}
