package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"

	highlighting "github.com/yuin/goldmark-highlighting"

	"github.com/gurch101/notes/utils"
)

func main() {
	os.RemoveAll("dist")
	matches, err := filepath.Glob("post/**/*.md")
	if err != nil {
		panic(err)
	}
	tmpl, err := ioutil.ReadFile("template/post.tmpl")
	tpl := template.Must(template.New("post").Parse(string(tmpl)))

	var cssBuf bytes.Buffer

	markdown := goldmark.New(
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
	)

	for _, match := range matches {
		var buf bytes.Buffer

		md, err := ioutil.ReadFile(match)
		if err != nil {
			panic(err)
		}

		context := parser.NewContext()

		if err := markdown.Convert(md, &buf, parser.WithContext(context)); err != nil {
			panic(err)
		}
		metaData := meta.Get(context)

		if metaData == nil {
			metaData = make(map[string]interface{})
		}
		metaData["StaticDirectory"] = strings.Repeat("../", strings.Count(match, "/")-1)
		metaData["Content"] = template.HTML(buf.String())

		fileName := match[strings.LastIndex(match, "/"):]
		dirName := fmt.Sprintf("dist/%s", match[5:strings.LastIndex(match, "/")])
		if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
			panic(err)
		}
		outFile := fmt.Sprintf("%s/%s.html", dirName, strings.TrimSuffix(fileName, ".md"))
		f, err := os.Create(outFile)
		if err != nil {
			panic(err)
		}
		if err := tpl.Execute(f, metaData); err != nil {
			panic(err)
		}
		f.Close()
	}

	utils.CopyDir("template/css", "dist/css")
	ioutil.WriteFile("dist/css/code.css", cssBuf.Bytes(), os.ModePerm)
}
