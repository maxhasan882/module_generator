package generator

import "template/file"

// Run generates code based on the parsed YAML data.
func Run(data *Data) {
	var err error
	if err = file.RemoveDirectory("domain"); err != nil {
		panic(err)
	}
	if err = file.RemoveDirectory("domain/repository"); err != nil {
		panic(err)
	}
	if err = file.RemoveDirectory("infrastructure/repository/"); err != nil {
		panic(err)
	}

	generateModels(data)
	generateInterface(data)
	generateSchemaModels(data)
}

// Other generator functions...
