# goldmark-obsidian-callout

Extension for the [goldmark](https://github.com/yuin/goldmark) markdown parser. Parses [obsidian-flavored callouts](https://help.obsidian.md/Editing+and+formatting/Callouts). 

## Examples

```markdown
> [!note]+ Sidenote
> This is something really interesting
```

Gets transformed into:

<details class="callout" data-callout="note" open>
  <summary>Sidenote</summary>
	<p>This is something <strong>really</strong> interesting</p>
</details>

```html
<details class="callout" data-callout="note" open>
  <summary>Sidenote</summary>
	<p>This is something <strong>really</strong> interesting</p>
</details>
```

The CSS class and `data-callout` are there to make the callouts compatible with Obsidian stylesheets and themes.

## Features

- [x] Recognizes all Obsidian callout types
- [x] Custom titles
- [x] Multiple paragraphs inside the callout
- [x] Non-collapsible callouts by default *(requires Javascript - preview in browser)*
- [ ] Nested callouts
- [ ] Default callout title accoring to the callout type

Contributions are very much welcome!

## Credit

Project bootstrapped by the [mangoubrella/goldmark-figure](https://github.com/mangoumbrella/goldmark-figure) project. Thanks!

## License

MIT
