package main

import (
	"github.com/gurch101/notes/commands"
)

func main() {
	clean := commands.NewCleanCommand("dist")
	err := clean.Execute()
	if err != nil {
		panic(err)
	}
	build := commands.NewBuildCommand("pages", "templates", "dist")
	err = build.Execute()
	if err != nil {
		panic(err)
	}
}
