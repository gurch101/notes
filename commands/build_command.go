package commands

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gurch101/notes/api"
	"github.com/gurch101/notes/utils"
)

// A BuildCommand binds all input files to their templates and renders the result to the output directory
type BuildCommand struct {
	inputDirectory    string
	templateDirectory string
	outputDirectory   string
}

// NewBuildCommand constructs a BuildCommand
func NewBuildCommand(inputDirectory string, templateDirectory string, outputDirectory string) *BuildCommand {
	return &BuildCommand{inputDirectory, templateDirectory, outputDirectory}
}

func (build *BuildCommand) copyCSS(codeCSSBuffer bytes.Buffer) {
	utils.CopyDir(
		fmt.Sprintf("%s/css", build.templateDirectory),
		fmt.Sprintf("%s/css", build.outputDirectory),
	)

	if codeCSSBuffer.Len() > 0 {
		ioutil.WriteFile(
			fmt.Sprintf("%s/css/code.css", build.outputDirectory),
			codeCSSBuffer.Bytes(),
			os.ModePerm,
		)
	}
}

func (build *BuildCommand) getPagePathsFromInputDirectory() []string {
	matches, err := filepath.Glob(fmt.Sprintf("%s/**/*.md", build.inputDirectory))
	if err != nil {
		panic(err)
	}
	return matches
}

func (build *BuildCommand) createOutputDirectoriesIfNotExists(path string) (string, error) {
	dirName := fmt.Sprintf("%s/%s", build.outputDirectory, path)
	if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
		return "", fmt.Errorf("could not create directory %s", dirName)
	}
	return dirName, nil
}

// Execute binds each input file in inputDirectory to the appropriate template in templateDirectory and writes the result to the outputDirectory
func (build *BuildCommand) Execute() error {
	var cssBuf bytes.Buffer

	iterator := api.NewPageIterator(build.getPagePathsFromInputDirectory())
	markdown := api.NewMarkdownParser()
	// TODO: template loader
	tmpl, err := ioutil.ReadFile(fmt.Sprintf("%s/page.tmpl", build.templateDirectory))
	if err != nil {
		return fmt.Errorf("Could not read template")
	}
	tpl := template.Must(template.New("page").Parse(string(tmpl)))

	for iterator.HasNext() {
		page := iterator.Next()
		md, err := ioutil.ReadFile(page.FullPath)
		if err != nil {
			return fmt.Errorf("unable to read file %s", page.FullPath)
		}
		html, err := markdown.Transform(md)
		if err != nil {
			return err
		}
		if html.CSS.Len() > 0 {
			cssBuf = html.CSS
		}
		html.Meta["StaticDirectory"] = strings.Repeat("../", page.NumberOfSubdirectories)
		html.Meta["Content"] = template.HTML(html.Value.String())
		outputDir, err := build.createOutputDirectoriesIfNotExists(page.SubDirectories)
		if err != nil {
			return err
		}
		outputFile := fmt.Sprintf("%s/%s.html", outputDir, page.FileNameWithoutExtension)
		f, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("Could not create file %s", outputFile)
		}
		if err := tpl.Execute(f, html.Meta); err != nil {
			return fmt.Errorf("Could not render template for page %s", page.FullPath)
		}
		f.Close()
	}

	build.copyCSS(cssBuf)
	return nil
}
