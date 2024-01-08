package generator

import (
	"path/filepath"
	"template/file"
)

// Run generates code based on the parsed YAML data.
func Run(data *Data) {
	modules := map[string]struct {
		Trigger    func(*Data)
		ModulePath string
	}{
		"domain": {
			Trigger:    generateDomainModule,
			ModulePath: filepath.Join("domain", data.FileName),
		},
		"repository": {
			Trigger:    generateInterface,
			ModulePath: filepath.Join("domain", "repository", data.FileName),
		},
		"infrastructure": {
			Trigger:    generateInfraModule,
			ModulePath: filepath.Join("infrastructure", "repository", "schema", data.FileName),
		},
	}

	// Regenerate modules based on file existence
	regenerate := func(moduleType string) {
		module := modules[moduleType]
		if file.IsFileExists(module.ModulePath) {
			if askForConfirmation(module.ModulePath + " file already exists. Do you want to regenerate it?") {
				if err := file.RemoveFile(module.ModulePath); err != nil {
					panic(err)
				}
				module.Trigger(data)
			}
		} else {
			module.Trigger(data)
		}
	}
	for _, moduleType := range []string{"domain", "repository", "infrastructure"} {
		regenerate(moduleType)
	}
}
