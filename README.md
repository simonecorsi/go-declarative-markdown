# Declarative Markdown

## About

This package helps in generating markdown in a declarative way without having to handle strings yourself

## Requirements

- git
- go 1.16

## Installation

```
go get github.com/simonecorsi/go-declarative-markdown
```

## Usage

```
package main

import mkd "github.com/simonecorsi/go-declarative-markdown"

function main() {
  md := mkd.CreateMarkdown("My h1 Header")
  md.Header("About")
  md.Paragraph("Lorem Ipsum")
  md.Render() // save to file or whatever
}

```

## API

### CreateMarkdown

`CreateMarkdown(title string) *Markdown`

### GetLine

### Render

### Header

### AddLine

### Paragraph

### Quote

### Code

### Image

### List

### Task

### Table

<!-- CONTRIBUTING -->

## Contributing

Project is pretty simple and straight forward for what is my needs, but if you have any idea you're welcome.

This projects uses [commitizen](https://github.com/commitizen/cz-cli) so be sure to use standard commit format or PR won't be accepted

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'feat(scope): some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->

## Contact

Simone Corsi - [@im_simonecorsi](https://twitter.com/im_simonecorsi)
