package generator

import (
	"path/filepath"
	"template/file"
)

// Run generates code based on the parsed YAML data.
func Run(data *Data) {
	if file.IsFileExists(filepath.Join("domain", data.FileName)) {
		if askForConfirmation(filepath.Join("domain", data.FileName) + " file already exist\nWant to regenerate it?") {
			if err := file.RemoveFile(filepath.Join("domain", data.FileName)); err != nil {
				panic(err)
			}
			generateDomainModule(data)
		}
	} else {
		generateDomainModule(data)
	}

	if file.IsFileExists(filepath.Join("domain", "repository", data.FileName)) {
		if askForConfirmation(filepath.Join("domain", "repository", data.FileName) + " file already exist\nWant to regenerate it?") {
			if err := file.RemoveFile(filepath.Join("domain", "repository", data.FileName)); err != nil {
				panic(err)
			}
			generateInterface(data)
		}
	} else {
		generateInterface(data)
	}

	if file.IsFileExists(filepath.Join("infrastructure", "repository", "schema", data.FileName)) {
		if askForConfirmation(filepath.Join("infrastructure", "repository", "schema", data.FileName) + " file already exist\nWant to regenerate it?") {
			if err := file.RemoveFile(filepath.Join("infrastructure", "repository", "schema", data.FileName)); err != nil {
				panic(err)
			}
			generateInfraModule(data)
		}
	} else {
		generateInfraModule(data)
	}
}
