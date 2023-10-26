package callout_test

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"

	callout "github.com/VojtaStruhar/goldmark-obsidian-callout"
)

var markdown = goldmark.New(
	goldmark.WithExtensions(
		callout.ObsidianCallout,
	),
)

func TestBlockquote(t *testing.T) {
	var count = 0
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
	var count = 0
	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Callout With custom title",
		Markdown: `
> [!info] Custom callout title
`,
		Expected: `<details class="callout" data-callout="info" open onclick="return false">
<summary>
 Custom callout title
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
		Expected: `<details class="callout" data-callout="tip" open onclick="return false">
<summary>

</summary>
</details>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Unknown callout type -> note",
		Markdown: `
> [!unknown]
`,
		Expected: `<details class="callout" data-callout="note" open onclick="return false">
<summary>

</summary>
</details>
`,
	}, t)
}
func TestCalloutsWithContent(t *testing.T) {
	var count = 0
	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Callout with paragraph content",
		Markdown: `
> [!info]
> Some content here
`,
		Expected: `<details class="callout" data-callout="info" open onclick="return false">
<summary>

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
		Expected: `<details class="callout" data-callout="info" open onclick="return false">
<summary>

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
		Expected: `<details class="callout" data-callout="info" open onclick="return false">
<summary>

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
		Expected: `<details class="callout" data-callout="example" open onclick="return false">
<summary>
 Some title
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
		Expected: `<details class="callout" data-callout="example" open onclick="return false">
<summary>
 Some title
</summary>
<p>More paragraphs</p>
<p>In a single callout</p>
</details>
`,
	}, t)
}

func TestExpandableCallouts(t *testing.T) {
	var count = 0
	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Default forced-open callout",
		Markdown: `
> [!example] Some title
> Some content
`,
		Expected: `<details class="callout" data-callout="example" open onclick="return false">
<summary>
 Some title
</summary>
<p>Some content</p>
</details>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Callout closed by default",
		Markdown: `
> [!example]- Closed by default
> Some content
`,
		Expected: `<details class="callout" data-callout="example">
<summary>
 Closed by default
</summary>
<p>Some content</p>
</details>
`,
	}, t)

	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Callout open by default",
		Markdown: `
> [!example]+ Open by default
> Some content
`,
		Expected: `<details class="callout" data-callout="example" open>
<summary>
 Open by default
</summary>
<p>Some content</p>
</details>
`,
	}, t)
}
