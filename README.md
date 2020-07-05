# Notes

A static site generator written in Go. Templates are simple Go HTML templates.  YAML front-matter in markdown files is passed to templates as variables. That's it!

### Installation

Get the binary for your operating system and add it to your `PATH`.

windows // TODO: GOOS=windows GOARCH=amd64 go build -o notes.exe main.go
linux // TODO: GOOS=linux GOARCH=amd64
mac // TODO: GOOS=darwin GOARCH=amd64

### Usage

1. create the following directory structure:

```
/src
    note1.md            # notes can contain YAML front-matter
/template
    css/                # template/css is minified and copied to dist/css
    note.tmpl           # template/*.tmpl are simple Go HTML files that can consume note front-matter
```

2. run `notes build` to build to `dist/`

3. run `notes deploy` to deploy to github pages

### Template Binding

`/src/**/index.md` is bound to `index.tmpl`
if a `template/tags.tmpl` exists, it's passed all file front-matter by tag for any markdown file that contains tags. Tag files are output to `/dist/tags/<tag-name>.html`
If a `template/collection.tmpl` exists, it's passed all notes in title order for a given collection
All other files in `/src` are bound to `note.tmpl`

# TODO

- add github pipeline (fmt, vet, test, coverage, build, create release - https://github.com/actions/create-release)
- index.md -> index.tmpl -> index.html for TOC
- Add homepage link
- separate binary for "notes"
- args parser for "notes build"
- deploy to s3/github pages via "notes deploy"
- enable tagging -> tags.tmpl gets passed file front-matter by tag
- local server with file watcher via "notes serve"
- publish go pkg
- css minifier
- create button for any template - shows an editor and uploads file to s3 directly
- draft mode
- search: tag json file
- activity page (by date, author, tag)
- comments
- follow tags/authors
