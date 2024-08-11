package blueprint

import (
	"html/template"
	"os"
)

// BluePrint represents any file that needs to be written to the file system.
type BluePrint struct {
	// Name of the file
	Name string
	// WritePath to where the file needs to be written
	WritePath string
	// Values needed for the data to be templated out
	Values any
	// Data the template that will be written to the filesystem
	Data string
}

// NewBluePrint generates a new blueprint to be used for write to the filesystem
func NewBluePrint(name, writePath, data string, values any) *BluePrint {
	return &BluePrint{
		Name:      name,
		WritePath: writePath,
		Values:    values,
		Data:      data,
	}
}

// Exists allows you to check if the file within the blueprint exists
func (b *BluePrint) Exists() (os.FileInfo, error) {
	return os.Stat(b.WritePath)
}

// Write takes the BluePrint data then templates it out to the filesystem
func (b *BluePrint) Write() error {
	tmpl := template.Must(template.New(b.Name).Parse(b.Data))
	file, err := os.OpenFile(b.WritePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o755)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := tmpl.Execute(file, b.Values); err != nil {
		return err
	}
	return nil
}
