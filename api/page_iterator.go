package api

import (
	"strings"
)

type PageIterator struct {
	currentIndex int
	pages        []string
}

type Page struct {
	FullPath                 string
	FileName                 string
	FileNameWithoutExtension string
	NumberOfSubdirectories   int
	SubDirectories           string
}

func NewPage(path string) *Page {
	filename := path[strings.LastIndex(path, "/")+1:]
	filenameWithoutExtension := filename[:strings.LastIndex(filename, ".")]
	numberOfSubdirectories := strings.Count(path, "/") - 1
	subDirectories := ""
	if numberOfSubdirectories > 0 {
		subDirectories = path[strings.Index(path, "/")+1 : strings.LastIndex(path, "/")]
	}
	return &Page{
		FullPath:                 path,
		FileName:                 filename,
		FileNameWithoutExtension: filenameWithoutExtension,
		NumberOfSubdirectories:   numberOfSubdirectories,
		SubDirectories:           subDirectories,
	}
}

func NewPageIterator(pages []string) *PageIterator {
	return &PageIterator{currentIndex: 0, pages: pages}
}

func (p *PageIterator) HasNext() bool {
	return p.currentIndex < len(p.pages)
}

func (p *PageIterator) Next() *Page {
	pagePath := p.pages[p.currentIndex]
	p.currentIndex++
	return NewPage(pagePath)
}
