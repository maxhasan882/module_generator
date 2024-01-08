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

	generateDomainModule(data)
	generateInterface(data)
	generateInfraModule(data)
}

// Other generator functions...
