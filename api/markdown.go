package api

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// HTML is a transformed Markdown document
type HTML struct {
	Meta  map[string]interface{}
	CSS   bytes.Buffer
	Value bytes.Buffer
}

// A MarkdownParser transforms markdown documents into HTML.MarkdownParser
// Supports YAML front-matter and syntax highlighting.
type MarkdownParser struct {
	parser goldmark.Markdown
	cssBuf *bytes.Buffer
}

// NewMarkdownParser creates and initializes a new MarkdownParser
func NewMarkdownParser() *MarkdownParser {
	var cssBuf bytes.Buffer

	fmt.Printf(cssBuf.String())
	return &MarkdownParser{
		parser: goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(
					highlighting.WithStyle("emacs"),
					highlighting.WithFormatOptions(
						html.WithLineNumbers(true),
						html.WithClasses(true),
						html.LineNumbersInTable(true),
						html.TabWidth(2),
					),
					highlighting.WithCSSWriter(&cssBuf),
				),
				meta.Meta,
			),
			goldmark.WithParserOptions(parser.WithAttribute()),
		),
		cssBuf: &cssBuf,
	}
}

// Transform returns a HTML document with CSS (if document has code blocks) and metadata (if document has YAML front-matter)
func (p *MarkdownParser) Transform(markdown []byte) (*HTML, error) {
	var buf bytes.Buffer

	context := parser.NewContext()
	parser.WithAttribute()
	if err := p.parser.Convert(markdown, &buf, parser.WithContext(context)); err != nil {
		return nil, fmt.Errorf("unable to parse markdown")
	}

	metaData := meta.Get(context)

	if metaData == nil {
		metaData = make(map[string]interface{})
	}

	return &HTML{Meta: metaData, CSS: *p.cssBuf, Value: buf}, nil
}
