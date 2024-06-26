package blueprint

import "github.com/devbytes-cloud/hookinator/internal/blueprint/templates"

// NewConfig generates a blueprint for the config file
func NewConfig() *BluePrint {
	return NewBluePrint("config.json", "config.json", templates.Config, nil)
}
