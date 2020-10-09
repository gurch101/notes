package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyPageIterator(t *testing.T) {
	iterator := NewPageIterator([]string{})

	assert.Equal(t, false, iterator.HasNext(), "no pages exist")
}

func TestPageIterator(t *testing.T) {
	iterator := NewPageIterator([]string{"my/test/path.txt", "my/path2.txt"})

	assert.Equal(t, true, iterator.HasNext(), "page exists")

	p := iterator.Next()

	assert.Equal(t, "path.txt", p.FileName, "filename parsed")
	assert.Equal(t, "path", p.FileNameWithoutExtension, "filename without extension parsed")
	assert.Equal(t, 1, p.NumberOfSubdirectories, "number of subdirectories parsed")
	assert.Equal(t, "test", p.SubDirectories, "subdirectories parsed")

	p = iterator.Next()
	assert.Equal(t, "path2.txt", p.FileName, "filename parsed")
	assert.Equal(t, "path2", p.FileNameWithoutExtension, "filename without extension parsed")
	assert.Equal(t, 0, p.NumberOfSubdirectories, "number of subdirectories parsed")
	assert.Equal(t, "", p.SubDirectories, "subdirectories parsed")
}
