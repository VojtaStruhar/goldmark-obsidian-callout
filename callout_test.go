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

func TestMultipleCallouts(t *testing.T) {
	var count = 0
	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Multiple callouts",
		Markdown: `
> [!info] First callout
> Some content in the first callout

> [!info] Second callout
> Some content in the second callout
`,
		Expected: `<details class="callout" data-callout="info" open onclick="return false">
<summary>
 First callout
</summary>
<p>Some content in the first callout</p>
</details>
<details class="callout" data-callout="info" open onclick="return false">
<summary>
 Second callout
</summary>
<p>Some content in the second callout</p>
</details>
`,
	}, t)
}

func TestMultipleCalloutsWithBlockquote(t *testing.T) {
	var count = 0
	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Multiple callouts with a blockquote in between",
		Markdown: `
> [!info] First callout
> Some content in the first callout

> All we are is dust in the wind...dude.
> - Ted 

> [!info] Second callout
> Some content in the second callout
`,
		Expected: `<details class="callout" data-callout="info" open onclick="return false">
<summary>
 First callout
</summary>
<p>Some content in the first callout</p>
</details>
<blockquote>
<p>All we are is dust in the wind...dude.</p>
<ul>
<li>Ted</li>
</ul>
</blockquote>
<details class="callout" data-callout="info" open onclick="return false">
<summary>
 Second callout
</summary>
<p>Some content in the second callout</p>
</details>
`,
	}, t)
}

func TestNestedCallouts(t *testing.T) {
	var count = 0
	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Multiple callouts",
		Markdown: `
>[!info] This is a callout
>Text for the callout
>>[!info] This is inside the outer callout
>>More text inside the callout.
		`,
		Expected: `
<details class="callout" data-callout="info" open onclick="return false">
<summary>
 This is a callout
</summary>
<p>Text for the callout</p>
<details class="callout" data-callout="info" open onclick="return false">
<summary>
 This is inside the outer callout
</summary>
<p>More text inside the callout.</p>
</details>
</details>
		`,
	}, t)
}
func TestNestedBlockquoteInsideCallout(t *testing.T) {
	var count = 0
	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Multiple callouts",
		Markdown: `
>[!info] This is a callout
>Text for the callout
>>This is a blockquote inside the callout
>
>More text inside the callout.
		`,
		Expected: `
<details class="callout" data-callout="info" open onclick="return false">
<summary>
 This is a callout
</summary>
<p>Text for the callout</p>
<blockquote>
<p>This is a blockquote inside the callout</p>
</blockquote>
<p>More text inside the callout.</p>
</details>
		`,
	}, t)
}
func TestNestedBlockquoteInsideNestedCallout(t *testing.T) {
	var count = 0
	count++
	testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
		No:          count,
		Description: "Multiple callouts",
		Markdown: `
>[!info] This is a callout
>Text for the callout
>>[!info] This is an inner callout
>>Text inside the inner clalout
>>>This is a blockquote inside the callout
>>
>>More text inside the inner callout.
>
>More text inside the outer callout
		`,
		Expected: `
<details class="callout" data-callout="info" open onclick="return false">
<summary>
 This is a callout
</summary>
<p>Text for the callout</p>
<details class="callout" data-callout="info" open onclick="return false">
<summary>
 This is an inner callout
</summary>
<p>Text inside the inner clalout</p>
<blockquote>
<p>This is a blockquote inside the callout</p>
</blockquote>
<p>More text inside the inner callout.</p>
</details>
<p>More text inside the outer callout</p>
</details>
		`,
	}, t)
}
