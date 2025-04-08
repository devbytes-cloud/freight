package blueprint

import "github.com/devbytes-cloud/freight/internal/blueprint/templates"

// NewConfig generates a blueprint for the config file
func NewConfig() *BluePrint {
	return NewBluePrint("railcar.json", "railcar.json", templates.Config, nil)
}
