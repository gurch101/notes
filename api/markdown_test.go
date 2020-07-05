package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyMarkdown(t *testing.T) {
	parser := NewMarkdownParser()

	doc, err := parser.Transform([]byte(""))
	assert.Nil(t, err)
	assert.Equal(t, "", doc.Value.String(), "empty doc should return blank string")
	assert.Equal(t, "", doc.CSS.String(), "empty doc should return no css")
	assert.Equal(t, 0, len(doc.Meta), "empty doc should have no metadata")
}

func TestMarkdownNoCode(t *testing.T) {
	parser := NewMarkdownParser()

	doc, err := parser.Transform([]byte("# Hello World"))
	assert.Nil(t, err)
	assert.Equal(t, "<h1>Hello World</h1>\n", doc.Value.String(), "markdown converted to HTML")
	assert.Equal(t, "", doc.CSS.String(), "markdown without code should return no css")
	assert.Equal(t, 0, len(doc.Meta), "markdown without metadata should return no metadata")
}

func TestMarkdownNoMetadata(t *testing.T) {
	parser := NewMarkdownParser()

	doc, err := parser.Transform([]byte("# Hello World\n```go\nfunc main(){}\n```"))
	assert.Nil(t, err)
	assert.NotEmpty(t, doc.Value.String(), "markdown converted to HTML")
	assert.NotEmpty(t, doc.CSS.String(), "markdown with code should return css")
	assert.Equal(t, 0, len(doc.Meta), "empty doc should have no metadata")
}

func TestMarkdown(t *testing.T) {
	parser := NewMarkdownParser()

	doc, err := parser.Transform([]byte("---\nFoo: Bar\n---\n\n# Hello World\n```go\nfunc main(){}\n```"))
	assert.Nil(t, err)
	assert.NotEmpty(t, doc.Value.String(), "markdown converted to HTML")
	assert.NotEmpty(t, doc.CSS.String(), "markdown with code should return css")
	assert.Equal(t, 1, len(doc.Meta), "empty doc should have no metadata")
}
