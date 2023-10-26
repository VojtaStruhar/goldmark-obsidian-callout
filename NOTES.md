# Developer notes

The AST is imperfect. It still has the parenting blockquote, when it gets to the AST transformer:

- Blockquote
  - Callout (future `<details>`)
    - CalloutTitle (future `<summary>`)
    - ...Paragraph

I guess this is why the AST transformer is there in the first place? But still, could be cleaner. See `parser.calloutAstTransformer.Transform`