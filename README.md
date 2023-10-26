# goldmark-callout

## Documentation

The AST is imperfect. It still has the parenting blockquote, like this:

- Blockquote
  - Callout (`<details>`)
    - CalloutTitle (`<summary>`)
    - ...Paragraph

Only in the AST transformer, blockquotes whose first child is a Callout are *dismissed* in a way - no 
`<blockquote>` element is rendered.

## Research

Keep details open:

```html
<details open onclick="return false">
  <summary>Details</summary>
  Something small enough to escape casual notice.
</details>

```

## License

MIT
