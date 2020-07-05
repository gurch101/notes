package commands

import "os"

// CleanCommand removes the provided directory and all subdirectories
type CleanCommand struct {
	outputDirectory string
}

// NewCleanCommand creates a CleanCommand
func NewCleanCommand(directory string) *CleanCommand {
	return &CleanCommand{directory}
}

// Execute removes the directory and all subdirectories
func (clean *CleanCommand) Execute() error {
	return os.RemoveAll(clean.outputDirectory)
}
